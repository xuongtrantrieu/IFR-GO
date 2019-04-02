package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText   string      `json:"status"`
	AppCode      int         `json:"code"`
	Message      string      `json:"message,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func new(httpStatusCode int, statusText string, appCode int, message string, errorMessage string, data interface{}) Response {
	return Response{
		HTTPStatusCode: httpStatusCode,
		StatusText:     statusText,
		AppCode:        appCode,
		Message:        message,
		ErrorMessage:   errorMessage,
		Data:           data,
	}
}

func (response Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, response.HTTPStatusCode)
	return nil
}

func ResponseOK(w http.ResponseWriter, r *http.Request, data interface{}, message ...string) error {
	_message := ""
	if len(message) > 0 {
		_message = message[0]
	}

	response := new(http.StatusOK, "OK", http.StatusOK, _message, "", data)
	render.Render(w, r, &response)
	return nil
}

func ResponseERR(w http.ResponseWriter, r *http.Request, errMessage ...string) render.Renderer {
	_errMessage := ""
	if len(errMessage) > 0 {
		_errMessage = errMessage[0]
	}

	response := new(http.StatusBadRequest, "Bad request", http.StatusBadRequest, "", _errMessage, nil)
	render.Render(w, r, &response)
	return nil
}
