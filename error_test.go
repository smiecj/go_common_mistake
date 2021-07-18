package go_common_mistake

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"testing"
)

var (
	errorCodeToMsgErr = map[int]string{
		1: "connect error",
		2: "handle error",
	}
)

type MyError struct {
	Code     int
	ErrorMsg string
}

func (err MyError) Error() string {
	return fmt.Sprintf("%s", err.ErrorMsg)
}

func newError(code int) error {
	if msg, ok := errorCodeToMsgErr[code]; ok {
		err := new(MyError)
		err.Code, err.ErrorMsg = code, msg
		return errors.Wrapf(err, "server error")
	}
	return errors.New("normal error")
}

func TestWrapfAndCause(t *testing.T) {
	serverErr := newError(1)
	log.Println(serverErr.Error())

	switch errors.Cause(serverErr).(type) {
	case *MyError:
		log.Println("This is a server error")
	default:
		log.Println("This is an unknown error")
	}
}
