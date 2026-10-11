package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

// ===================================== //
// =========== TESTS A ET AN =========== //
// ===================================== //

// Fonction générale pour tester TraitementAEtAn //
func VerifierTraitementAEtAn(t *testing.T, entree []string, attendu []string) {
	resultat := regles.TraitementAEtAn(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu: %q", resultat, attendu)
	}
}

func TestTraitementAEtAnMajuscule(t *testing.T) {
	VerifierTraitementAEtAn(t, []string{"A", "human"}, []string{"An", "human"})
}

func TestTraitementAEtAnMinuscule(t *testing.T) {
	VerifierTraitementAEtAn(t, []string{"a", "apple"}, []string{"an", "apple"})
}

func TestTraitementAEtAnPlusieursCas(t *testing.T) {
	VerifierTraitementAEtAn(t, []string{"a", "Elephant", "an", "doctor", "a", "universe", "A", "object"}, []string{"an", "Elephant", "an", "doctor", "an", "universe", "An", "object"})
}

func TestTraitementAEtAnDernierMot(t *testing.T) {
	VerifierTraitementAEtAn(t, []string{"I", "want", "a"}, []string{"I", "want", "a"})
}
