package web

// The presentation layer contains all resources concerned with creating an application interface
// Contains code designed to be used for http rest based api interface

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "net/http/pprof"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vituchon/labora-golang-course/meeting-crud-api/presentation/web/controllers"
)

func StartServer() {
	go func() { // una forma que se acopla bien a cualquier implementación de servidor web que usemos (http, mux, chi, etc.)
		const port = 9091
		var portStr string = ":" + strconv.Itoa(port)
		fmt.Printf("Server for profiling listening at port %v\n", portStr)
		err := http.ListenAndServe(portStr, nil)
		if err != nil {
			fmt.Println("Unable to start server for profiling afairrs: ", err)
		}
		// Para que funciones se llevan más tiempo (y detectar posibles cuellos de botella) se debe hacer lo que llaman un "profile":
		// -- en realidad sirve para ver cuanto tiempo de CPU asigna a cada función de nuestro programa, ojota que aparecen MUCHISIMAS funciones que son de los paquetes que importamos y no tenemos idea de que existian hasta ahora... ufff, bueno vale ignorar! se puede buscar por nombre de función y van a ver que se remarcan los cuadros de sus funciones con las métricas!!!!
		// 1) arrancar la app
		// 2) luego ejecutar `curl --output pprof.out "localhost:9091/debug/pprof/profile?seconds=10"` y se guardar en el archivo pprof.out información de perfil, realmente hay que hacerlo trabajar al servidor para que salga el reporte!! consideren usar el truco de hacer muchos hits (peticiones a un endpoint) usando los comandos que vienen con el interprete de comandos (bash), yo hice algo como esto: `for((i=1;i<=100;i+=1)); do curl "http://localhost:9090/api/v1/animals"; done``, sí se ejecuta 100 veces un mismo curl!! que fácil es bombardear la red de peticiones!!!! no lo hagan en casa!
		// 3) luego ejecutar `go tool pprof -http localhost:9092 profile.out` para ue se vea lindo a travéz de una pestaña del navegador
		// ----
		// de forma alternativa pueden entrar a http://localhost:9091/debug/pprof/ y explorar un poco... se ven muchas cosas... no llegue muy lejos tampoco...
	}()
	router := buildRouter()

	// Descomentar en caso de querer probar forma alternativa que sirve para mux!. Tomada de https://www.jajaldoang.com/post/profiling-go-app-with-pprof/ y https://groups.google.com/g/golang-nuts/c/TjDMXyBDYG4
	/*router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/heap", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	*/
	server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 300 * time.Second,
	}
	fmt.Printf("Animal crud api server listening at port %v\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Unable to start app: ", err)
	}
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
