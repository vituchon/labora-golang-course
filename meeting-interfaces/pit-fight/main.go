package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/vituchon/labora-golang-course/meeting-interfaces/fighters"
)

func main() {
	const col = 15
	scanner := bufio.NewScanner(os.Stdin)
	var valueEnteredSignal chan int = make(chan int)
	var wasEnterPressed bool

	var contenders []fighters.Contender = buildContenders()
	var areAllAlive = areAllContendersAlive(contenders)
	for areAllAlive {

		var intensity float64 = float64(contenders[0].ThrowAttack())

		wasEnterPressed = false
		go drawBarUntil(func(t int) bool {
			return !wasEnterPressed
		}, valueEnteredSignal, col)

		scanner.Scan()
		wasEnterPressed = true
		var value float64 = float64(<-valueEnteredSignal)

		fmt.Println("Has hecho un ataque de intensidad ", intensity, " y se apaciguara en ", (value * 100 / col), "%")
		intensity = intensity * (value / col)

		fmt.Println(contenders[0].GetName(), " tira golpe con intensidad =", intensity)
		contenders[1].RecieveAttack(int(intensity))

		if contenders[1].IsAlive() {
			intensity := contenders[1].ThrowAttack()
			fmt.Println(contenders[1].GetName(), " tira golpe con intensidad =", intensity)
			contenders[0].RecieveAttack(intensity)
		}

		fmt.Printf("%sLife=%d, %sLife=%d\n", contenders[0].GetName(), contenders[0].GetLife(), contenders[1].GetName(), contenders[1].GetLife())
		areAllAlive = areAllContendersAlive(contenders)
		time.Sleep(3 * time.Second)
	}
}

func drawBarUntil(cond func(t int) bool, valueEnteredSignal chan int, col int) int {
	bar := fmt.Sprintf("[%%-%vs]", col)
	var t int = 0
	for cond(t) {
		fmt.Print("\033[H\033[2J")
		fmt.Printf(bar, strings.Repeat("=", t%col)+"🤜")

		time.Sleep(20 * time.Millisecond)

		t++
	}
	valueEnteredSignal <- t % col
	return t % col
}

var enterPressedSignal chan bool = make(chan bool)

//var valueEnteredSignal chan int = make(chan int)

func IsKeyPressed() bool {
	select {
	case <-enterPressedSignal:
		return true
	default:
		return false
	}
}

func main2() {
	var contenders []fighters.Contender = buildContenders()
	var areAllAlive = areAllContendersAlive(contenders)
	for areAllAlive {
		intensity := contenders[0].ThrowAttack()
		fmt.Println(contenders[0].GetName(), " tira golpe con intensidad =", intensity)
		contenders[1].RecieveAttack(intensity)

		if contenders[1].IsAlive() {
			intensity := contenders[1].ThrowAttack()
			fmt.Println(contenders[1].GetName(), " tira golpe con intensidad =", intensity)
			contenders[0].RecieveAttack(intensity)
		}

		fmt.Printf("%sLife=%d, %sLife=%d\n", contenders[0].GetName(), contenders[0].GetLife(), contenders[1].GetName(), contenders[1].GetLife())
		areAllAlive = areAllContendersAlive(contenders)
		time.Sleep(3 * time.Second)
	}
}

func areAllContendersAlive(contenders []fighters.Contender) bool {
	size := len(contenders)
	var i int = 0
	for i = 0; i < size && contenders[i].IsAlive(); i++ {
	}
	return i == size
}

func buildContenders() []fighters.Contender {
	var police fighters.Police = fighters.Police{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
		Armour: 5,
	}
	/*var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}*/
	var paladin fighters.Paladin = fighters.Paladin{
		BaseFighter: fighters.BaseFighter{
			Life: fighters.PaladinInitialLife,
		},
	}
	var contenders []fighters.Contender = make([]fighters.Contender, 2)

	randomValueBetweenOneAndZero := rand.Intn(2)
	contenders[randomValueBetweenOneAndZero] = &police
	contenders[(randomValueBetweenOneAndZero+1)%2] = &paladin
	return contenders
}

func main_legacy() {

	var police fighters.Police = fighters.Police{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
		Armour: 5,
	}
	var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}

	randomValueBetweenOneAndZero := rand.Intn(2)
	policeHitFirst := randomValueBetweenOneAndZero == 1

	var areBothAlive = police.IsAlive() && criminal.IsAlive()
	for areBothAlive {
		if policeHitFirst {
			intesity := police.ThrowAttack()
			fmt.Println("Policia tira golpe con intensidad =", intesity)
			criminal.RecieveAttack(intesity)

			if criminal.IsAlive() {
				intesity := criminal.ThrowAttack()
				fmt.Println("Criminal tira golpe con intensidad =", intesity)
				police.RecieveAttack(intesity)
			}
		} else {
			intesity := criminal.ThrowAttack()
			fmt.Println("Criminal tira golpe con intensidad =", intesity)
			police.RecieveAttack(intesity)

			if police.IsAlive() {
				intesity := police.ThrowAttack()
				fmt.Println("Policia tira golpe con intensidad =", intesity)
				criminal.RecieveAttack(intesity)
			}
		}
		fmt.Printf("PoliceLife=%d, CriminalLife=%d\n", police.Life, criminal.Life)
		areBothAlive = police.IsAlive() && criminal.IsAlive()
		time.Sleep(3 * time.Second)
	}

}
