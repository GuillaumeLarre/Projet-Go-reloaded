package tests

import (
	"go-reloaded/regles"
	"slices"
	"strings"
	"testing"
)

// ================================== //
// =========== TESTS HEXA =========== //
// ================================== //

// Tests de la fonction ConvertirHexaEnNombre //
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

// Tests de la fonction ConvertirBinEnNombre //
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

// ================================ //
// =========== TESTS UP =========== //
// ================================ //

func VerifierAppliquerMarqueurCasse(t *testing.T, entree []string, marqueur string, fonction func(string) string, attendu []string) {
	resultat := regles.AppliquerMarqueurCasse(entree, marqueur, fonction)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestAppliquerMarqueurCasseUpSansIndice(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"Voici", "un", "exemple", "(up)"}, "(up", strings.ToUpper, []string{"Voici", "un", "EXEMPLE"})
}

func TestAppliquerMarqueurCasseUpAvecIndice(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"Voici", "un", "exemple", "(up, 2)"}, "(up", strings.ToUpper, []string{"Voici", "UN", "EXEMPLE"})
}

func TestAppliquerMarqueurCasseUpAvecIndiceTropLong(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"Voici", "un", "exemple", "(up, 6)"}, "(up", strings.ToUpper, []string{"Voici", "un", "exemple", "(up, 6)"})
}

func TestAppliquerMarqueurCasseUpAvecIndiceNegatif(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"Voici", "un", "exemple", "(up, -1)"}, "(up", strings.ToUpper, []string{"Voici", "un", "exemple", "(up, -1)"})
}

func TestAppliquerMarqueurCasseUpMarqueurDebut(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"(up)", "Voici", "un", "exemple"}, "(up", strings.ToUpper, []string{"(up)", "Voici", "un", "exemple"})
}

func TestAppliquerMarqueurCasseUpMarqueurInvalide(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"Voici", "un", "exemple", "(up, 3x)"}, "(up", strings.ToUpper, []string{"Voici", "un", "exemple", "(up, 3x)"})
}

func TestAppliquerMarqueurCasseUpAvecIndiceZero(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"Voici", "un", "exemple", "(up, 0)"}, "(up", strings.ToUpper, []string{"Voici", "un", "exemple", "(up, 0)"})
}

func TestAppliquerMarqueurCasseUpPlusieursMarqueurs(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"Voici", "(up)", "un", "exemple", "(up, 2)"}, "(up", strings.ToUpper, []string{"VOICI", "UN", "EXEMPLE"})
}

// ================================= //
// =========== TESTS LOW =========== //
// ================================= //

func TestAppliquerMarqueurCasseLowSansIndice(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"VOICI", "UN", "EXEMPLE", "(low)"}, "(low", strings.ToLower, []string{"VOICI", "UN", "exemple"})
}

func TestAppliquerMarqueurCasseLowAvecIndice(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 2)"}, "(low", strings.ToLower, []string{"VOICI", "un", "exemple"})
}

func TestAppliquerMarqueurCasseLowAvecIndiceTropLong(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 6)"}, "(low", strings.ToLower, []string{"VOICI", "UN", "EXEMPLE", "(low, 6)"})
}

func TestAppliquerMarqueurCasseLowAvecIndiceNegatif(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"VOICI", "UN", "EXEMPLE", "(low, -1)"}, "(low", strings.ToLower, []string{"VOICI", "UN", "EXEMPLE", "(low, -1)"})
}

func TestAppliquerMarqueurCasseLowMarqueurDebut(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"(low)", "VOICI", "UN", "EXEMPLE"}, "(low", strings.ToLower, []string{"(low)", "VOICI", "UN", "EXEMPLE"})
}

func TestAppliquerMarqueurCasseLowMarqueurInvalide(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 3x)"}, "(low", strings.ToLower, []string{"VOICI", "UN", "EXEMPLE", "(low, 3x)"})
}

func TestAppliquerMarqueurCasseLowAvecIndiceZero(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 0)"}, "(low", strings.ToLower, []string{"VOICI", "UN", "EXEMPLE", "(low, 0)"})
}

func TestAppliquerMarqueurCasseLowPlusieursMarqueurs(t *testing.T) {
	VerifierAppliquerMarqueurCasse(t, []string{"VOICI", "(low)", "UN", "EXEMPLE", "(low, 2)"}, "(low", strings.ToLower, []string{"voici", "un", "exemple"})
}
