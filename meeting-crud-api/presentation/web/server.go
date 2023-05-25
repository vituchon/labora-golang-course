package web

// The presentation layer contains all resources concerned with creating an application interface
// Contains code designed to be used for http rest based api interface

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vituchon/labora-golang-course/meeting-crud-api/presentation/web/controllers"
)

func StartServer() {

	router := buildRouter()
	server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 300 * time.Second,
	}
	fmt.Printf("animal crud api server listening at port %v", server.Addr)
	server.ListenAndServe()
}

func buildRouter() *mux.Router {
	root := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./"))
	// TODO : word "presentation" in the path may be redudant, perpahs using just "assets" would be enought!
	root.PathPrefix("/presentation/web/assets").Handler(fileServer)
	root.NotFoundHandler = http.HandlerFunc(NoMatchingHandler)

	Get := BuildSetHandleFunc(root, "GET")
	Get("/healthcheck", controllers.Healthcheck)
	Get("/version", controllers.Version)

	api := root.PathPrefix("/api/v1").Subrouter()
	api.Use(AccessLogMiddleware) // only logs api calls
	apiGet := BuildSetHandleFunc(api, "GET")
	apiPost := BuildSetHandleFunc(api, "POST")
	apiPut := BuildSetHandleFunc(api, "PUT")
	apiDelete := BuildSetHandleFunc(api, "DELETE")

	apiGet("/animals", controllers.GetAnimals)                  //curl -X GET http://localhost:9090/api/v1/animals
	apiGet("/animals/{id:[0-9]+}", controllers.GetAnimalById)   //curl -X GET http://localhost:9090/api/v1/animals/1
	apiPost("/animals", controllers.CreateAnimal)               //curl -X POST http://localhost:9090/api/v1/animals --data-binary '{ "name" : "Igo", "kind": 2 }'
	apiPut("/animals/{id:[0-9]+}", controllers.UpdateAnimal)    //curl -X PUT http://localhost:9090/api/v1/animals/1 --data-binary '{ "Id": 1, "name" : "Koko", "kind": 0 }'
	apiDelete("/animals/{id:[0-9]+}", controllers.DeleteAnimal) //curl -x DELETE http://localhost:9090/api/v1/animals/5

	return root
}

type setHandlerFunc func(path string, f http.HandlerFunc)

// Creates a function for register a handler for a path for the given router and http methods
func BuildSetHandleFunc(router *mux.Router, methods ...string) setHandlerFunc {
	return func(path string, f http.HandlerFunc) {
		router.HandleFunc(path, f).Methods(methods...)
	}
}

func NoMatchingHandler(response http.ResponseWriter, request *http.Request) {
	errMsg := fmt.Sprintf("No maching route for " + request.URL.Path)
	fmt.Println(errMsg)
	http.Error(response, errMsg, http.StatusNotFound)

	/*if request.URL.Path == "/favicon.ico" { // avoids to trigger another request to landing or login on the "silent" http request by chrome to get an icon! I guess i could tell chrome for ubuntu that redirection for an icon can create more and bigger troubles than solutions... i mean nobody dies for an icon... for now...
		response.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(response, request, "/presentation/web/assets/images/logo.png", http.StatusSeeOther)*/
}

// Adds a logging handler for logging each request's in Apache Common Log Format (CLF).
// With this middleware we ensure that each requests will be, at least, logged once.
func AccessLogMiddleware(h http.Handler) http.Handler {
	loggingHandler := handlers.LoggingHandler(os.Stdout, h)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggingHandler.ServeHTTP(w, r)
	})
}
