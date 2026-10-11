package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

// ========================================= //
// =========== TESTS PONCTUATION =========== //
// ========================================= //

// Fonction générale pour tester TraitementPonctuation //
func VerifierTraitementPonctuation(t *testing.T, entree []string, attendu []string) {
	resultat := regles.TraitementPonctuation(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Résultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestTraitementPonctuationCasNormal(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici,", "un", "exemple"}, []string{"Voici,", "un", "exemple"})
}

func TestTraitementPonctuationPonctuationApresMot(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici", ",un", "exemple"}, []string{"Voici,", "un", "exemple"})
}

func TestTraitementPonctuationPlusieursPonctuations(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici", ",,un", "exemple"}, []string{"Voici,,", "un", "exemple"})
}

func TestTraitementPonctuationPlusieursCas(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici", "?un", "exemple", "."}, []string{"Voici?", "un", "exemple."})
}

func TestTraitementPonctuationDebutMot(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"!Voici", "un", "exemple"}, []string{"!Voici", "un", "exemple"})
}

func TestTraitementPonctuationMultiplesCas(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"%µVoici,", "un", "...exemple", "?", "fin", "!?"}, []string{"%µVoici,", "un...", "exemple?", "fin!?"})
}

func TestTraitementPonctuationSujet(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"then", "BAMM", "!!"}, []string{"then", "BAMM!!"})
}

func TestTraitementPonctuationSymbolleSpecifique(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"then", "BAMM", "…"}, []string{"then", "BAMM…"})
}
