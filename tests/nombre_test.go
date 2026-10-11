package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

// ================================== //
// =========== TESTS HEXA =========== //
// ================================== //

// Fonction générale pour tester ConvertirHexaEnNombre //
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

// ================================= //
// =========== TESTS BIN =========== //
// ================================= //

// Fonction générale pour tester ConvertirBinEnNombre //
func VerificationConvertirBin(t *testing.T, entree string, attendu string) {
	resultat := regles.ConvertirBinEnNombre(regles.DecoupeEnMots(entree))
	if !slices.Equal(resultat, regles.DecoupeEnMots(attendu)) {
		t.Errorf("Resultat : %q, Attendu : %q", regles.AssembleEnTexte(resultat), attendu)
	}
}

func TestConvertirBinCasNormal(t *testing.T) {
	VerificationConvertirBin(t, "101 (bin) files were added", "5 files were added")
}

func TestConvertirBinMarqueurPremierePosition(t *testing.T) {
	VerificationConvertirBin(t, "(bin) files were added", "(bin) files were added")
}

func TestConvertirBinPlusieursMarqueurs(t *testing.T) {
	VerificationConvertirBin(t, "101 (bin) files and 11 (bin) folders", "5 files and 3 folders")
}

func TestConvertirBinMotNonBin(t *testing.T) {
	VerificationConvertirBin(t, "hello (bin) world", "hello (bin) world")
}

func TestConvertirBinMarqueurFin(t *testing.T) {
	VerificationConvertirBin(t, "1010 (bin) files and 10 (bin)", "10 files and 2")
}

func TestConvertirBinNombreNonBin(t *testing.T) {
	VerificationConvertirBin(t, "102 (bin) world", "102 (bin) world")
}
