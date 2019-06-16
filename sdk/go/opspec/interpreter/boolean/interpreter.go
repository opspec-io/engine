package boolean

//go:generate counterfeiter -o ./fakeInterpreter.go --fake-name FakeInterpreter ./ Interpreter

import (
	"fmt"
	"github.com/opctl/opctl/sdk/go/data/coerce"
	"github.com/opctl/opctl/sdk/go/model"
	"github.com/opctl/opctl/sdk/go/opspec/interpreter/value"
)

type Interpreter interface {
	// Interpret interprets an expression to an boolean value.
	// Expression must be a type supported by coerce.ToBoolean
	Interpret(
		scope map[string]*model.Value,
		expression interface{},
		opHandle model.DataHandle,
	) (*model.Value, error)
}

// NewInterpreter returns an initialized Interpreter instance
func NewInterpreter() Interpreter {
	return _interpreter{
		coerce:           coerce.New(),
		valueInterpreter: value.NewInterpreter(),
	}
}

type _interpreter struct {
	coerce           coerce.Coerce
	valueInterpreter value.Interpreter
}

func (itp _interpreter) Interpret(
	scope map[string]*model.Value,
	expression interface{},
	opHandle model.DataHandle,
) (*model.Value, error) {
	value, err := itp.valueInterpreter.Interpret(
		expression,
		scope,
		opHandle,
	)
	if nil != err {
		return nil, fmt.Errorf("unable to interpret %+v to boolean; error was %v", expression, err)
	}

	return itp.coerce.ToBoolean(&value)
}
