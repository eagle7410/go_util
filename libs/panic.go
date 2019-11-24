package lib

import (
	"errors"
	"fmt"
	"net/http"
)

type PanicData struct {
	Type int
	Mess string
}

func (i *PanicData) CheckAndPanic(isPanic bool, m string, data ...interface{}) {
	if isPanic {
		i.Panic(m, data...)
	}
}

func (i *PanicData) CheckAndPanicBadReq(isPanic bool, m string, data ...interface{}) {
	if isPanic {
		i.PanicBadReq(m, data...)
	}
}

func (i *PanicData) IsTechProblem() bool {
	return i.Type == http.StatusInternalServerError
}

func (i *PanicData) CheckAndPanicTechProblem(isPanic bool, m string, data ...interface{}) {
	if isPanic {
		i.PanicTechProblem(m, data...)
	}
}

func (i *PanicData) PanicTechProblem(m string, data ...interface{}) {
	i.Type = http.StatusInternalServerError
	i.Mess = fmt.Sprintf(m, data...)
	panic(i)
}

func (i *PanicData) PanicBadReq(m string, data ...interface{}) {
	i.Type = http.StatusBadRequest
	i.Mess = fmt.Sprintf(m, data...)
	panic(i)
}

func (i *PanicData) Panic(m string, data ...interface{}) {
	i.Mess = fmt.Sprintf(m, data...)
	panic(i)
}

func (i *PanicData) SetDataRender(t int, m string, data ...interface{}) *PanicData {
	i.Type = t
	i.Mess = fmt.Sprintf(m, data...)
	return i
}

func (i *PanicData) SetData(m string, t int) *PanicData {
	i.Type = t
	i.Mess = m

	return i
}

func (i *PanicData) SetType(t int) *PanicData {
	i.Type = t
	return i
}

func (i *PanicData) CheckError(e error) {
	if e != nil {
		panic(i.SetError(e))
	}
}

func (i *PanicData) SetError(e error) *PanicData {
	i.Mess = e.Error()
	return i
}

func (i *PanicData) SetMessage(mess string) *PanicData {
	i.Mess = mess
	return i
}

func (i *PanicData) Error() string {
	return i.Mess
}

func (i *PanicData) NewError() error {
	return errors.New(i.Mess)
}
