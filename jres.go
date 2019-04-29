package jres

import (
	"encoding/json"
	"net/http"
)

type resModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Info    interface{} `json:"info"`

	Errors []string `json:"errors"`
}

func Send(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(status)

	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}

	return nil
}

func Accepted(w http.ResponseWriter, data interface{}) error {
	return Send(w, http.StatusAccepted, resModel{Data: data})
}

func BadRequest(w http.ResponseWriter, msg string) error {
	res := resModel{Message: "bad request"}
	if msg == "" {
		res.Message = msg
	}

	return Send(w, http.StatusBadRequest, res)
}

func Conflict(w http.ResponseWriter, errors []string) error {
	return Send(w, http.StatusConflict, resModel{Message: "conflict", Errors: errors})
}

func Created(w http.ResponseWriter, at string, data interface{}) error {
	w.Header().Add("Location", at)

	return Send(w, http.StatusCreated, resModel{Message: "entity created", Data: data})
}

func NoContent(w http.ResponseWriter) error {
	return Send(w, http.StatusNoContent, nil)
}

func NotFound(w http.ResponseWriter, msg string) error {
	data := resModel{Message: "record not found"}
	if msg == "" {
		data.Message = msg
	}

	return Send(w, http.StatusNotFound, data)
}

func Forbidden(w http.ResponseWriter) error {
	return Send(w, http.StatusForbidden, resModel{Message: "forbidden"})
}

func MethodNotAllwed(w http.ResponseWriter, errors []string) error {
	return Send(w, http.StatusMethodNotAllowed, resModel{Message: "method not allowed", Errors: errors})
}

func OK(w http.ResponseWriter, data interface{}) error {
	return Send(w, http.StatusOK, resModel{Data: data})
}

func Redirect(w http.ResponseWriter, to string) error {
	w.Header().Add("Location", to)

	return Send(w, http.StatusTemporaryRedirect, nil)
}

func ServerError(w http.ResponseWriter) error {
	return Send(w, http.StatusInternalServerError, resModel{Message: "internal server error"})
}

func Unauthorized(w http.ResponseWriter, data interface{}) error {
	return Send(w, http.StatusUnauthorized, resModel{Message: "unauthorized", Data: data})
}

func UnprocessableEntity(w http.ResponseWriter, data interface{}) error {
	return Send(w, http.StatusUnprocessableEntity, resModel{Message: "unprocessable entity", Data: data})
}

func ValidationError(w http.ResponseWriter, errors []string) error {
	return Send(w, http.StatusBadRequest, resModel{Message: "validation error", Errors: errors})
}
