package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

// ====================================== //
// =========== TESTS DECOUPES =========== //
// ====================================== //

// Tests de la fonction DecoupeEnMots //
func VerifierDecoupe(t *testing.T, entree string, attendu []string) {
	resultat := regles.DecoupeEnMots(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestDecoupeEnMotsPlusieursEspaces(t *testing.T) {
	VerifierDecoupe(t, "Voici  un exemple", []string{"Voici", "un", "exemple"})
}

func TestDecoupeEnMotsAucunEspaceFinal(t *testing.T) {
	VerifierDecoupe(t, "Un nouveau test", []string{"Un", "nouveau", "test"})
}

func TestDecoupeEnMotsTexteVide(t *testing.T) {
	VerifierDecoupe(t, "", []string{})
}

// Tests de la fonction AssembleEnTexte //
func VerifierAssemble(t *testing.T, entree []string, attendu string) {
	resultat := regles.AssembleEnTexte(entree)
	if resultat != attendu {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestAssembleEnTexteListeVide(t *testing.T) {
	VerifierAssemble(t, []string{}, "")
}

func TestAssembleEnTexteCasNormal(t *testing.T) {
	VerifierAssemble(t, []string{"Un", "cas", "normal"}, "Un cas normal")
}

func TestAssembleEnTexteUnSeulMot(t *testing.T) {
	VerifierAssemble(t, []string{"Bonjour"}, "Bonjour")
}

// Tests de la fonction DecoupeEnLignes //
func VerifierDecoupeEnLigne(t *testing.T, entree string, attendu []string) {
	resultat := regles.DecoupeEnLignes(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestDecoupeEnLigneCasNormal(t *testing.T) {
	VerifierDecoupeEnLigne(t, "Hello\nworld", []string{"Hello", "world"})
}

func TestDecoupeEnLigneSautDeLigneFin(t *testing.T) {
	VerifierDecoupeEnLigne(t, "Hello\nworld\n", []string{"Hello", "world", ""})
}

func TestDecoupeEnLigneCasLigneVide(t *testing.T) {
	VerifierDecoupeEnLigne(t, "Hello\n\nworld", []string{"Hello", "", "world"})
}

// Tests de la fonction AssembleEnLignes //
func VerifierAssembleEnLigne(t *testing.T, entree []string, attendu string) {
	resultat := regles.AssembleEnLignes(entree)
	if resultat != attendu {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestAssembleEnLigneCasNormal(t *testing.T) {
	VerifierAssembleEnLigne(t, []string{"Hello", "world"}, "Hello\nworld")
}

func TestAssembleEnLigneCasLigneVide(t *testing.T) {
	VerifierAssembleEnLigne(t, []string{"Hello", "", "world"}, "Hello\n\nworld")
}

func TestAssembleEnLigneSautDeLigneFin(t *testing.T) {
	VerifierAssembleEnLigne(t, []string{"Hello", "world", ""}, "Hello\nworld\n")
}
