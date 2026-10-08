package tests

import (
	"go-reloaded/regles"
	"testing"
)

// ======================================== //
// =========== TESTS TRAITEMENT =========== //
// ======================================== //

// Tests de la fonction de traitement //
func VerifierTraitement(t *testing.T, entree string, attendu string) {
	resultat := regles.Traitement(entree)
	if resultat != attendu {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestTraitementPlusieursLignes(t *testing.T) {
	VerifierTraitement(t, "10 (bin) fichiers traités\nB (hex) dossiers", "2 fichiers traités\n11 dossiers")
}

func TestTraitementLigneVideEntreDeuxLignes(t *testing.T) {
	VerifierTraitement(t, "10 (bin) fichiers traités\n\nB (hex) dossiers", "2 fichiers traités\n\n11 dossiers")
}

func TestTraitementSujet(t *testing.T) {
	VerifierTraitement(t, "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68.")
}
