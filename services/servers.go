package services

import (
	"net/http"

	"github.com/zefrenchwan/m3.git/storage"
)

func BuildServices(serviceContainer *ServiceServer, context ServiceContext) {
	serviceContainer.AddServiceHandler("GET", "/user", context, upsertUser)
}

func LaunchHandler(addr string, dao storage.Dao) {
	mux := http.NewServeMux()
	handler := ServiceServer{mux}
	context := ServiceContext{Dao: dao}
	BuildServices(&handler, context)
	server := &http.Server{
		Addr:    addr,
		Handler: handler.mux,
	}
	server.ListenAndServe()
}
