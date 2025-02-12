package main

import (
	"fmt"
	"math/rand"
	"os"
	"src/fonctions"
	"time"
)

func main() {
	var listePropositions []string
	var proposition string
	var tentatives int
	pv := 10
	rand.Seed(time.Now().UnixNano())
	lignesHangman := fonctions.ChargerLignes("hangman.txt")
	listeMots := fonctions.ChargerLignes(os.Args[1])
	mot := listeMots[rand.Intn(len(listeMots))]
	lettresAffiches := fonctions.InitLettresAffichees(mot)
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Tentatives réalisées: ", tentatives)
		posHangman := ((pv - ((pv - 5) * 2)) * (8)) + 8
		for _, l := range lignesHangman[posHangman-8 : posHangman] {
			fmt.Println(l)
		}
		fonctions.AfficherLettres(mot, lettresAffiches)
		proposition = fonctions.Demander(listePropositions)
		tentatives += 1
		listePropositions = append(listePropositions, proposition)
		if proposition == mot {
			fmt.Print("Bravo, vous avez trouvé le mot: ", mot)
			break
		} else if len(proposition) > 1 {
			pv -= 2
		} else {
			if fonctions.LettrePresente(mot, proposition) {
				for i, l := range mot {
					if string(l) == proposition {
						if !fonctions.ElementPresente(lettresAffiches, i) {
							lettresAffiches = append(lettresAffiches, i)
						}
					}
				}
			} else {
				pv -= 1
			}
		}
		if len(lettresAffiches) == len(mot) {
			fmt.Print("Bravo, vous avez trouvé le mot: ", mot)
			break
		}
		if pv <= 0 {
			fmt.Print("Vous avez perdu ! Le mot était ", mot)
			break
		}
	}

}
