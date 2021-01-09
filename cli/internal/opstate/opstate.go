package opstate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/opctl/opctl/cli/internal/clioutput"
	"github.com/opctl/opctl/sdks/go/model"
	"golang.org/x/term"
)

// CallGraph maintains a record of the current state of an op
type CallGraph struct {
	rootNode *callGraphNode
	errors   []error
}

type callGraphNode struct {
	call     *model.Call
	state    string
	loader   *Loading
	children []*callGraphNode
}

func newCallGraphNode(call *model.Call) *callGraphNode {
	return &callGraphNode{
		call:     call,
		children: []*callGraphNode{},
		loader:   &Loading{},
	}
}

var errNotFoundInGraph = errors.New("not found in graph")

func (n *callGraphNode) insert(call *model.Call) error {
	if n.call.ID == *call.ParentID {
		n.children = append(n.children, newCallGraphNode(call))
		return nil
	}
	for _, child := range n.children {
		err := child.insert(call)
		if err == nil {
			return nil
		}
	}
	return errNotFoundInGraph
}

func (n *callGraphNode) find(call *model.Call) (*callGraphNode, error) {
	if call.ID == n.call.ID {
		return n, nil
	}
	for _, child := range n.children {
		c, _ := child.find(call)
		if c != nil {
			return c, nil
		}
	}
	return nil, errNotFoundInGraph
}

var muted = color.New(color.Faint)
var highlighted = color.New(color.Bold)
var success = color.New(color.FgGreen)
var failed = color.New(color.FgRed)
var warning = color.New(color.FgYellow)

func (n callGraphNode) String(cliOutput clioutput.CliOutput) string {
	childLen := len(n.children)
	var desc string
	if n.call.Container != nil {
		desc = muted.Sprint(n.call.Container.ContainerID[:8]) + " "
		if n.call.Container.Name != nil {
			desc += highlighted.Sprint(*n.call.Container.Name)
		} else {
			desc += *n.call.Container.Image.Ref
			if len(n.call.Container.Cmd) > 0 {
				desc += " " + muted.Sprint(strings.ReplaceAll(strings.Join(n.call.Container.Cmd, " "), "\n", "\\n"))
			}
		}
	} else if n.call.Op != nil {
		desc = highlighted.Sprint(cliOutput.FormatOpRef(n.call.Op.OpPath))
	} else if n.call.Parallel != nil {
		desc = "parallel"
	} else if n.call.ParallelLoop != nil {
		desc = "parallel loop"
	} else if n.call.Serial != nil {
		desc = "serial"
	} else if n.call.SerialLoop != nil {
		desc = "serial loop"
	}
	str := "◉ "
	switch n.state {
	case model.OpOutcomeSucceeded:
		str += success.Sprint("☑ ")
	case model.OpOutcomeFailed:
		str += failed.Sprint("⚠ ")
	case model.OpOutcomeKilled:
		str += "️☒ "
	case "":
		// only display loading spinner on leaf nodes
		if len(n.children) == 0 {
			str += n.loader.String() + " "
		}
	default:
		str += n.state + " "
	}
	if n.call.Name != nil {
		str += highlighted.Sprint(*n.call.Name) + " "
	}
	str += desc
	for i, child := range n.children {
		childLines := strings.Split(child.String(cliOutput), "\n")
		for j, part := range childLines {
			if j == 0 {
				if i < childLen-1 {
					str += fmt.Sprintf("\n├─%s", part)
				} else {
					str += fmt.Sprintf("\n└─%s", part)
				}
			} else if i < childLen-1 {
				str += fmt.Sprintf("\n│ %s", part)
			} else {
				str += fmt.Sprintf("\n  %s", part)
			}
		}
	}
	return str
}

// String returns a visual representation of the current state of the call graph
func (g CallGraph) String(cliOutput clioutput.CliOutput) string {
	str := g.rootNode.String(cliOutput)
	for _, err := range g.errors {
		str += "\n" + warning.Sprint("⚠️  ") + err.Error()
	}
	return str
}

// HandleEvent accepts an opctl event and updates the call graph appropriately
func (g *CallGraph) HandleEvent(event *model.Event) error {
	if event.CallStarted != nil {
		if event.CallStarted.Call.ParentID == nil {
			if g.rootNode == nil {
				g.rootNode = newCallGraphNode(&event.CallStarted.Call)
				return nil
			}
			return errors.New("parent node already set")
		}
		return g.rootNode.insert(&event.CallStarted.Call)
	} else if event.CallEnded != nil {
		node, err := g.rootNode.find(&event.CallEnded.Call)
		if err != nil {
			err = fmt.Errorf("bad ended event %s, %v: %v", event.CallEnded.Call.ID, event.CallEnded.Ref, err)
			g.errors = append(g.errors, err)
			return err
		}
		node.state = event.CallEnded.Outcome
	}
	return nil
}

// OutputManager allows printing a "resettable" thing at the bottom of a stream
// of terminal output, when a tty is used
type OutputManager struct {
	lastHeight int
}

// Clear clears the last thing printed by this object
func (o *OutputManager) Clear() {
	// cursor to start of line (real big number)
	fmt.Print("\033[100000D")
	// clear line
	fmt.Print("\033[K")
	for i := 1; i < o.lastHeight; i++ {
		// move up one line
		fmt.Print("\033[1A")
		// clear line
		fmt.Print("\033[K")
	}
}

var ansiRegex = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")

func stripAnsi(str string) string {
	return ansiRegex.ReplaceAllString(str, "")
}

func countChars(str string) int {
	count := 0
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		count++
		str = str[:len(str)-size]
	}
	return count
}

// Print prints the given string, with a preceding separator and width limited
// to the size of the terminal
func (o *OutputManager) Print(str string) error {
	w, _, err := term.GetSize(0)
	if err != nil {
		return err
	}
	lines := strings.Split(str, "\n")

	ruleWidth := 0
	for _, line := range lines {
		visualLen := countChars(stripAnsi(line))
		if visualLen > ruleWidth {
			ruleWidth = visualLen
		}
	}
	if ruleWidth > w {
		ruleWidth = w
	}

	fmt.Println(strings.Repeat("┄", ruleWidth))

	for i, line := range lines {
		withoutAnsi := stripAnsi(line)
		if countChars(withoutAnsi) > w {
			// TODO: use original line with ansi codes stripped appropriately
			fmt.Print(withoutAnsi[:w-1] + "…")
		} else {
			fmt.Print(line)
		}
		if i < len(lines)-1 {
			fmt.Print("\n")
		}
	}

	o.lastHeight = len(lines) + 1
	return nil
}

// Loading has a string representation that dynamically changes over time to
// display a visual loading spinner
type Loading struct {
	state       int
	lastChanged time.Time
}

var loadingRunes = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}

func (l *Loading) String() string {
	now := time.Now()
	ms := now.UnixNano() / int64(time.Millisecond)
	r := loadingRunes[(ms/int64(100))%int64(len(loadingRunes))]
	return string(r)
}