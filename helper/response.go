package helper

import (
	"net/http"
	"strings"
)

func SuccessResponse(code int, message string, data ...any) (int, map[string]interface{}) {
	response := make(map[string]interface{})
	response["message"] = message

	switch len(data) {
	case 1:
		response["data"] = data[0]
	case 2:
		response["data"] = data[0]
		response["token"] = data[1]
	}
	return code, response
}

func ErrorResponse(err error) (int, interface{}) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError
	msg := err.Error()

	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "not found"):
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "Duplicate"):
		if strings.Contains(msg, "username") {
			resp["message"] = "username is already in use"
			code = http.StatusConflict
		} else if strings.Contains(msg, "email") {
			resp["message"] = "email is already in use"
			code = http.StatusConflict
		} else {
			resp["message"] = "internal server error"
			code = http.StatusInternalServerError
		}
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "validation"):
		resp["message"] = ValidationError(err)
		code = http.StatusBadRequest
	case strings.Contains(msg, "unmarshal"):
		if strings.Contains(msg, "fullname") {
			resp["message"] = "invalid fullname of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "username") {
			resp["message"] = "invalid username of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "gender") {
			resp["message"] = "invalid gender of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "email") {
			resp["message"] = "invalid email of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "password") {
			resp["message"] = "invalid password of type string"
			code = http.StatusBadRequest
		}
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	}
	return code, resp
}