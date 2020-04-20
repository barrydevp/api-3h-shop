package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

type Exec func(Context *gin.Context) (interface{}, error)
type SuccessFunc func(Context *gin.Context, data interface{})
type ErrorFunc func(Context *gin.Context, error error)

type Handle struct {
	Context *gin.Context
	error   error
	data    interface{}
	done    bool
}

func (handle *Handle) checkContext() bool {
	if handle.Context == nil {
		handle.error = errors.New("missing Context in Handle")
		handle.done = true

		return false
	}

	return true
}

func (handle *Handle) Try(exec Exec) *Handle {
	if handle.done || !handle.checkContext() {
		log.Println("handle is completed before")

		return handle
	}

	data, error := exec(handle.Context)
	handle.done = true

	if error != nil {
		handle.error = error

		return handle
	}

	handle.data = data

	return handle
}

func (handle *Handle) Then(callback SuccessFunc) *Handle {
	if !handle.done {
		log.Println("Handle need to run Try before Then")

		return handle
	}

	if handle.error != nil {
		return handle
	}

	callback(handle.Context, handle.data)

	return handle
}

func (handle *Handle) Catch(callback ErrorFunc) *Handle {
	if !handle.done {
		log.Println("Handle need to run Try before Catch")

		return handle
	}

	if handle.error == nil {
		return handle
	}

	callback(handle.Context, handle.error)

	return handle
}
