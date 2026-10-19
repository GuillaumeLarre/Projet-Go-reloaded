package main

import (
	"fmt"
	"go-reloaded/regles"
	"os"
)

// ============================ //
// ====== MON PROGRAMME ======= //
// ============================ //

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
	texteTraite := regles.Traitement(string(contenu))
	err = os.WriteFile(os.Args[2], []byte(texteTraite), 0644)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Printf("Le fichier %s a bien été créé.\n", os.Args[2])
}
