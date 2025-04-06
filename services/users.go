package services

import (
	"net/http"
)

func upsertUser(context ServiceContext, w http.ResponseWriter, r *http.Request) ServiceError {
	result := ServiceError{}
	w.Write([]byte("C'est reparti comme en 46!!!"))
	return result
}
