package fonctions

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Charger un fichier en paramètre et renvoyer un slice des lignes.
// Quitte le programme affiche une erreur si le fichier est invalide.
func ChargerLignes(fichier string) []string {
	file, err := os.ReadFile(fichier)
	if err != nil {
		print("Erreur: le fichier '", fichier, "' est invalide.\n")
		os.Exit(3)
	}
	lignes := strings.Split(string(file), "\n")
	return lignes

}

// Ajoute à la liste des lettres à afficher les indices du début.
func InitLettresAffichees(mot string) []int {
	var nblettres int = (len(mot) / 2) - 1
	var lettresAffiches []int
	boolt := false
	for i := 0; i < nblettres; i++ {
		erand := rand.Intn(len(mot))
		// Vérifie si erand est dans le tableau lettresAffiches.
		for _, v := range lettresAffiches {
			if v == erand {
				boolt = true
				break
			}
		}
		if boolt {
			boolt = false
			continue
		}
		lettresAffiches = append(lettresAffiches, erand)
	}
	return lettresAffiches
}

// Afficher le mot en remplaçant les lettres non trouvées par "_".
func AfficherLettres(mot string, lettresAffiches []int) {
	for i, l := range mot {
		fmt.Print(" ")
		boolt := true
		for _, v := range lettresAffiches {
			if v == i {
				fmt.Print(strings.ToUpper(string(l)))
				boolt = false
			}
		}
		if boolt {
			fmt.Print("_")
		}
	}
}

// Demande la lettre ou le mot à proposer.
func Demander(listePropositions []string) string {
	proposition := ""
	for {
		dejaFaite := false
		fmt.Print("\n\nLettre ou mot à proposer : ")
		fmt.Scan(&proposition)
		proposition = strings.ToLower(proposition)
		// Vérifie si la lettre ou le mot proposé
		// est dans la liste des propositions deja faites.
		for _, e := range listePropositions {
			if string(e) == proposition {
				fmt.Print("\nVous avez deja fait cette proposition.")
				dejaFaite = true
			}
		}
		if !dejaFaite {
			return proposition
		}
	}
}

// Renvoi true si mot est présent dans lettreProposee.
func LettrePresente(mot, lettreProposee string) bool {
	for _, e1 := range mot {
		if string(e1) == lettreProposee {
			return true
		}
	}
	return false
}

// Renvoi true si un élément de liste se trouve dans element.
func ElementPresente(liste []int, element int) bool {
	for _, e1 := range liste {
		if e1 == element {
			return true
		}
	}
	return false
}
