package regles

import (
	"strings"
)

// Traitement des lignes //
func DecoupeEnLignes(texte string) []string {
	lignes := strings.Split(texte, "\n")
	return lignes
}

func AssembleEnLignes(liste []string) string {
	texte := strings.Join(liste, "\n")
	return texte
}

// Traitement des mots //
func DecoupeEnMots(texte string) []string {
	liste := strings.Fields(texte)
	return liste
}

func AssembleEnTexte(liste []string) string {
	phrase := strings.Join(liste, " ")
	return phrase
}
