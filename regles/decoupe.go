package regles

import (
	"strings"
)

// ================================ //
// =========== DECOUPES =========== //
// ================================ //

// Découpe les lignes //
func DecoupeEnLignes(texte string) []string {
	lignes := strings.Split(texte, "\n")
	return lignes
}

// Rassemble les lignes //
func AssembleEnLignes(liste []string) string {
	texte := strings.Join(liste, "\n")
	return texte
}

// Découpe les mots //
func DecoupeEnMots(texte string) []string {
	liste := strings.Fields(texte)
	return liste
}

// Rassemble les mots //
func AssembleEnTexte(liste []string) string {
	phrase := strings.Join(liste, " ")
	return phrase
}
