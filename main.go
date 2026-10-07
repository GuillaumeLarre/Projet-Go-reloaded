package main

import (
	"fmt"
	"go-reloaded/regles"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Mauvais nombre de paramètre renseigné. Le programme doit être exécuté avec 'go run . sample.txt result.txt'")
		return
	}
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	mots := regles.DecoupeEnMots(string(contenu))
	mots = regles.ConvertirHexaEnNombre(mots)
	mots = regles.ConvertirBinEnNombre(mots)
	texte := regles.AssembleEnTexte(mots)
	err = os.WriteFile(os.Args[2], []byte(texte), 0644)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Printf("Le fichier %s à bien été crée/modifié", os.Args[2])
}
