package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	router := mux.NewRouter() // aca creamos el coso para configurar el servidor, el coso se llama router

	// aca definimos el comportamiento del servidor
	router.HandleFunc("/", handleRootResource).Methods("GET")

	router.HandleFunc("/otro", handleOtherResource).Methods("GET")

	router.HandleFunc("/chatgpt", chatgptHandler).Methods("GET")

	// aca termine de definir el comportamiento

	// levantar el servidor en un "puerto"
	var portNumber int = 9999
	fmt.Println("Listen in port ", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ADIOS")
}

func handleRootResource(response http.ResponseWriter, request *http.Request) {
	msg := "Soy el recurso raiz"
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(msg))
}

func handleOtherResource(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("respuesta fallida"))
		}
	}()
	msg := "Soy otro recurso"
	var item *Item = nil
	fmt.Println(item.ID) // KA BOOM
	response.Write([]byte(msg))
	response.WriteHeader(http.StatusOK)
}

func chatgptHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("Endpointer, request from ip %s\n", request.RemoteAddr)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	//fmt.Fprintf(response, text)
	response.WriteHeader(http.StatusOK)
	_, err := response.Write([]byte(inputStr))
	if err != nil {
		fmt.Printf("error while writting bytes to response writer: %+v", err)
	}
}
