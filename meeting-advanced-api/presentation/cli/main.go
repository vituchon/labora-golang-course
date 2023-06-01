package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/presentation/cli/controllers"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories/postgres"
)

func main() {
	displayMenu()
	command := getCommand()
	for command != QuitCommand {
		performAction(command)
		displayMenu()
		command = getCommand()
	}
}

type Command int

const (
	CreateCommand   Command = 1
	RetrieveCommand Command = 2
	UpdateCommand   Command = 3
	DeleteCommand   Command = 4
	QuitCommand     Command = 5
)

func (c Command) Validate() error {
	if allCommands.contains(c) { // otra forma es hacer `if c >= 1 && c <= 5 { return nil }`, la desventaja es que uno tiene si o si que hacer que los numeros sean consecutivos... en cambio con este enfoque los números de comando peuden ser arbitrarios.. es menos efeciente Y permite mayor amplited (que ser menos efeciente sea malo así como mayor amplitud algo bueno es en todo caso el tema a tener en cuenta a la hora de tomar una decsisión)
		return nil
	}
	errMsg := fmt.Sprintf("'%d' NO es un comando soportado", c)
	return errors.New(errMsg)
}

var labelByCommand map[Command]string = map[Command]string{
	CreateCommand:   strconv.Itoa(int(CreateCommand)) + " => Agregar/Crear",
	RetrieveCommand: strconv.Itoa(int(RetrieveCommand)) + " => Listar/Recuperar",
	UpdateCommand:   strconv.Itoa(int(UpdateCommand)) + " => Modificar/Actualizar",
	DeleteCommand:   strconv.Itoa(int(DeleteCommand)) + " => Quitar/Eliminar",
	QuitCommand:     strconv.Itoa(int(QuitCommand)) + " => Finalizar programa",
}

func (c Command) Label() string {
	return labelByCommand[c]
}

type AllCommands []Command

var allCommands AllCommands = []Command{
	CreateCommand,
	RetrieveCommand,
	UpdateCommand,
	DeleteCommand,
	QuitCommand,
}

func (list AllCommands) contains(c Command) bool {
	for _, _c := range list {
		if c == _c {
			return true
		}
	}
	return false
}

func displayMenu() {
	fmt.Println("Seleccione una opción (escriba un número y pulse ENTER)")
	for _, command := range allCommands {
		fmt.Println(command.Label())
	}
}

var animalsRepository repositories.Animals

func init() {
	// INYECCION DE DEPEDENCIA (Repositorio en memoria o en base de daatos)
	// memoria
	//animalsRepository = memory.NewAnimalsStorage()

	// base de datos (postgres)
	animalsRepository = postgres.NewAnimalsStorage()
}

func performAction(command Command) {
	type ActionByCommand map[Command]func()
	var actionByCommand ActionByCommand = map[Command]func(){
		CreateCommand: func() {
			controllers.CreateAnimal(os.Stdin, os.Stdout)
		},
		RetrieveCommand: func() {
			displayRetrieveMenu()
			number, err := controllers.GetNumberFromStdin()
			if err == nil {
				performRetrieveOperation(number)
			} else {
				fmt.Printf("Error obteniendo opción %v", err)
			}
		},
		UpdateCommand: func() {
			controllers.UpdateAnimal(os.Stdin, os.Stdout)
		},
		DeleteCommand: func() {
			controllers.DeleteAnimal(os.Stdin, os.Stdout)
		},
	}
	action, exists := actionByCommand[command]
	if exists {
		action()
	}
}

func displayRetrieveMenu() {
	fmt.Println("Seleccione un criterio (escriba un número del 1 al 2 y pulse ENTER) para recuperar")
	fmt.Println("1 => Todos")
	fmt.Println("2 => Por id")
}

func performRetrieveOperation(number int) {
	switch number {
	case 1:
		controllers.GetAnimals(os.Stdin, os.Stdout)
		break
	case 2:
		controllers.GetAnimalById(os.Stdin, os.Stdout)
		break
	default:
		fmt.Printf("%d no es un opción invalida", number)
		break
	}
}

func getCommand() Command {
	command, err := getNumberFromStdinAndParseCommand()
	for err != nil {
		fmt.Printf("Error obteniendo comando : %v\n", err)
		command, err = getNumberFromStdinAndParseCommand()
	}
	return *command
}

func getNumberFromStdinAndParseCommand() (*Command, error) {
	number, err := controllers.GetNumberFromStdin()
	if err == nil {
		command := Command(number)
		return &command, command.Validate()
	}
	return nil, err
}
