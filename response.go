package package_http

import (
	"encoding/json"
	"net/http"
)

var Success = newResult(http.StatusOK)
var NoContent = newResult(http.StatusNoContent)
var BadRequest = newResult(http.StatusBadRequest)
var Conflict = newResult(http.StatusConflict)
var Unauthorized = newResult(http.StatusUnauthorized)
var Forbidden = newResult(http.StatusForbidden)
var UnprocessableEntity = newResult(http.StatusUnprocessableEntity)
var InternalServerError = newResult(http.StatusInternalServerError)
var TooManyRequests = newResult(http.StatusTooManyRequests)
var NotAcceptable = newResult(http.StatusNotAcceptable)
var Accepted = newResult(http.StatusAccepted)

type Response interface {
	SetData(data interface{}) *result
	SetStatusCode(statusCode int) *result
	GetStatusCode() int
	Marshal() []byte
}

type Result struct {
	Status  bool
	Message string
	Data    interface{}
}

type result struct {
	statusCode int
	data       interface{} // if is error some error data
}

func (r *result) SetData(data interface{}) *result {
	r.data = data
	return r
}

func (r *result) SetStatusCode(statusCode int) *result {
	r.statusCode = statusCode
	return r
}
func (r *result) GetStatusCode() int {
	return r.statusCode
}

func (r *result) Marshal() []byte {
	marshal, err := json.Marshal(r.data)
	if err != nil {
		return nil
	}
	return marshal
}

func newResult(statusCode int) *result {
	return &result{
		statusCode: statusCode,
	}
}
