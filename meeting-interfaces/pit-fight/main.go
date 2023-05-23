package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vituchon/labora-golang-course/meeting-interfaces/fighters"
)

func main() {
	var contenders []fighters.Contender = buildContenders()
	var areBothAlive = areAllContendersAlive(contenders)
	for areBothAlive {
		intensity := contenders[0].ThrowAttack()
		fmt.Println(contenders[0].GetName(), " tira golpe con intensidad =", intensity)
		contenders[1].RecieveAttack(intensity)

		if contenders[1].IsAlive() {
			intensity := contenders[1].ThrowAttack()
			fmt.Println(contenders[1].GetName(), " tira golpe con intensidad =", intensity)
			contenders[0].RecieveAttack(intensity)
		}

		fmt.Printf("%sLife=%d, %sLife=%d\n", contenders[0].GetName(), contenders[0].GetLife(), contenders[1].GetName(), contenders[1].GetLife())
		areBothAlive = areAllContendersAlive(contenders)
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
	var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}
	var contenders []fighters.Contender = make([]fighters.Contender, 2)

	randomValueBetweenOneAndZero := rand.Intn(2)
	contenders[randomValueBetweenOneAndZero] = &police
	contenders[(randomValueBetweenOneAndZero+1)%2] = &criminal
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
