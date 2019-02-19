// Code generated by gotemplate. DO NOT EDIT.

package exception

import (
	"bytes"
	"encoding/json"
	"foundation/log"
	"reflect"
	"strconv"
)

// template type Exception(PARENT,CODE,WHAT)

var AssertExceptionName = reflect.TypeOf(AssertException{}).Name()

type AssertException struct {
	Exception
	Elog log.Messages
}

func NewAssertException(parent Exception, message log.Message) *AssertException {
	return &AssertException{parent, log.Messages{message}}
}

func (e AssertException) Code() int64 {
	return AssertExceptionCode
}

func (e AssertException) Name() string {
	return AssertExceptionName
}

func (e AssertException) What() string {
	return "Assert Exception"
}

func (e *AssertException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e AssertException) GetLog() log.Messages {
	return e.Elog
}

func (e AssertException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e AssertException) DetailMessage() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(int(e.Code())))
	buffer.WriteByte(' ')
	buffer.WriteString(e.Name())
	buffer.Write([]byte{':', ' '})
	buffer.WriteString(e.What())
	buffer.WriteByte('\n')
	for _, l := range e.Elog {
		buffer.WriteByte('[')
		buffer.WriteString(l.GetMessage())
		buffer.Write([]byte{']', ' '})
		buffer.WriteString(l.GetContext().String())
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

func (e AssertException) String() string {
	return e.DetailMessage()
}

func (e AssertException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: AssertExceptionCode,
		Name: AssertExceptionName,
		What: "Assert Exception",
	}

	return json.Marshal(except)
}

func (e AssertException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*AssertException):
		callback(&e)
		return true
	case func(AssertException):
		callback(e)
		return true
	default:
		return false
	}
}
