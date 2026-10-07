package regles

import (
	"strings"
)

func DecoupeEnMots(texte string) []string {
	liste := strings.Fields(texte)
	return liste
}

func AssembleEnTexte(liste []string) string {
	phrase := strings.Join(liste, " ")
	return phrase
}
