package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

func VerificationConvertirHexa(t *testing.T, entree string, attendu string) {
	resultat := regles.ConvertirHexaEnNombre(regles.DecoupeEnMots(entree))
	if !slices.Equal(resultat, regles.DecoupeEnMots(attendu)) {
		t.Errorf("Resultat : %q, Attendu : %q", regles.AssembleEnTexte(resultat), attendu)
	}
}

func TestConvertirHexaCasNormal(t *testing.T) {
	VerificationConvertirHexa(t, "1E (hex) files were added", "30 files were added")
}

func TestConvertirHexaMarqueurPremierePosition(t *testing.T) {
	VerificationConvertirHexa(t, "(hex) files were added", "(hex) files were added")
}

func TestConvertirHexaPlusieursMarqueurs(t *testing.T) {
	VerificationConvertirHexa(t, "1E (hex) files and FF (hex) folders", "30 files and 255 folders")
}

func TestConvertirHexaMotNonHexa(t *testing.T) {
	VerificationConvertirHexa(t, "hello (hex) world", "hello (hex) world")
}

func TestConvertirHexaMarqueurFin(t *testing.T) {
	VerificationConvertirHexa(t, "1E (hex) files and FF (hex)", "30 files and 255")
}
