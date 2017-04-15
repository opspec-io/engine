package validate

import (
	"errors"
	"github.com/opspec-io/sdk-golang/model"
)

// validates an value against a socket parameter
func (this validate) socketParam(
	rawValue *string,
	param *model.SocketParam,
) (errs []error) {
	errs = []error{}

	// handle no value passed
	if nil == rawValue {
		errs = append(errs, errors.New("Socket required"))
	}
	return
}
