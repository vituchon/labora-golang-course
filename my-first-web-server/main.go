package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "¡Hola, mundooooo!")
	// })

	//Tarea.
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/item", getItem).Methods("GET")

	//Más adelante.
	// router.HandleFunc("/items", createItem).Methods("POST")
	// router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	// router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":9999", router)

}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Paula"},
	{ID: "2", Name: "Lucas"},
	{ID: "3", Name: "Misa"},
	{ID: "4", Name: "Rosario"},
	{ID: "5", Name: "Epik High"},
	{ID: "6", Name: "Pepe"},
	{ID: "7", Name: "Misa2"},
	{ID: "8", Name: "Rosario siempre estuvo cerca"},
	{ID: "9", Name: "Bokita"},
	{ID: "10", Name: "Burzaco"},
}

func getItems(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	v := query.Get("id")

	for _, item := range items {
		if item.ID == v {
			json.NewEncoder(w).Encode(item.Name)
		}
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Función para crear un nuevo elemento
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
}

/*
var items []Item = make([]Item, 0)

func main() {
	router := mux.NewRouter() // aca creamos el coso para configurar el servidor, el coso se llama router

	// aca definimos el comportamiento del servidor
	router.HandleFunc("/", handleRootResource).Methods("GET")
	router.HandleFunc("/otro", handleOtherResource).Methods("GET")
	router.HandleFunc("/chatgpt", chatgptHandler).Methods("GET")
	// aca termine de definir el comportamiento
	// los endpoints DETERMINAN LA API del server!!!!!!! (QUE SEA REST O NO DEPENDE DE SI SEGUIMOS O NO UNA CONVENCION!!)

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
	//request.URL.Query().Get("itemName")
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
*/
