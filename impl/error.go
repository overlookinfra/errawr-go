package impl

import (
	"fmt"

	errawr "github.com/reflect/errawr-go"
)

type ErrorDomain struct {
	Key   string
	Title string
}

type ErrorSection struct {
	Key   string
	Title string
}

type ErrorDescription struct {
	Friendly  string
	Technical string
}

type Error struct {
	ErrorDomain      *ErrorDomain
	ErrorSection     *ErrorSection
	ErrorCode        string
	ErrorTitle       string
	ErrorDescription *ErrorDescription
	ErrorArguments   ErrorArguments

	causes []errawr.Error
	buggy  bool
}

func (e Error) Code() string {
	return fmt.Sprintf(`%s_%s_%s`, e.ErrorDomain.Key, e.ErrorSection.Key, e.ErrorCode)
}

func (e Error) Arguments() map[string]interface{} {
	m := make(map[string]interface{})
	for k, a := range e.ErrorArguments {
		m[k] = a.Value
	}

	return m
}

func (e *Error) Bug() errawr.Error {
	e.buggy = true
	return e
}

func (e *Error) IsBug() bool {
	return e.buggy
}

func (e *Error) WithCause(cause errawr.Error) errawr.Error {
	e.causes = append(e.causes, cause)
	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf(`%s: %s`, e.Code(), e.ErrorDescription)
}
