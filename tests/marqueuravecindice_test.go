package tests

import (
	"go-reloaded/regles"
	"slices"
	"strings"
	"testing"
)

// =================================================== //
// =========== TESTS MARQUEURS AVEC INDICE =========== //
// =================================================== //

// Fonction générale pour tester UniformisationMarqueurs //
func VerifierUniformisationMarqueurs(t *testing.T, entree []string, marqueur string, attendu []string) {
	resultat := regles.UniformisationMarqueurs(entree, marqueur)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestAppliquerUniformisationMarqueursCap(t *testing.T) {
	VerifierUniformisationMarqueurs(t, []string{"Bonjour", "(cap,", "1)"}, "(cap,", []string{"Bonjour", "(cap, 1)"})
}

func TestAppliquerUniformisationMarqueursLow(t *testing.T) {
	VerifierUniformisationMarqueurs(t, []string{"Bonjour", "(low,", "2)"}, "(low,", []string{"Bonjour", "(low, 2)"})
}

func TestAppliquerUniformisationMarqueursUp(t *testing.T) {
	VerifierUniformisationMarqueurs(t, []string{"Bonjour", "(up,", "11)"}, "(up,", []string{"Bonjour", "(up, 11)"})
}

// Fonction générale pour tester UniformisationApresMarqueurs //
func VerifierUniformisationApresMarqueurs(t *testing.T, entree []string, marqueur string, attendu []string) {
	resultat := regles.UniformisationApresMarqueur(entree, marqueur)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestUniformisationApresMarqueurPonctuationCollee(t *testing.T) {
	VerifierUniformisationApresMarqueurs(t, []string{"go", "(up)!"}, "(up", []string{"go", "(up)", "!"})
}

func TestUniformisationApresMarqueurAvecNombreEtGroupe(t *testing.T) {
	VerifierUniformisationApresMarqueurs(t, []string{"so", "exciting", "(up, 2)?!"}, "(up", []string{"so", "exciting", "(up, 2)", "?!"})
}

func TestUniformisationApresMarqueurSeulInchange(t *testing.T) {
	VerifierUniformisationApresMarqueurs(t, []string{"go", "(up)"}, "(up", []string{"go", "(up)"})
}

func TestUniformisationApresMarqueurHex(t *testing.T) {
	VerifierUniformisationApresMarqueurs(t, []string{"1E", "(hex)."}, "(hex", []string{"1E", "(hex)", "."})
}

// Fonction générale pour tester AppliquerMarqueurCasse //
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

// ================================ //
// =========== TESTS UP =========== //
// ================================ //

// Fonction générale pour tester ConvertirListeMotsEnUp //
func VerifierConvertirListeMotsEnUp(t *testing.T, entree []string, attendu []string) {
	resultat := regles.ConvertirListeMotsEnUp(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestAppliquerConvertirListeMotsEnUpSansIndice(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "un", "exemple", "(up)"}, []string{"Voici", "un", "EXEMPLE"})
}

func TestAppliquerConvertirListeMotsEnUpAvecIndice(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "un", "exemple", "(up, 2)"}, []string{"Voici", "UN", "EXEMPLE"})
}

func TestAppliquerConvertirListeMotsEnUpAvecIndiceTropLong(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "un", "exemple", "(up, 6)"}, []string{"Voici", "un", "exemple", "(up, 6)"})
}

func TestAppliquerConvertirListeMotsEnUpAvecIndiceNegatif(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "un", "exemple", "(up, -1)"}, []string{"Voici", "un", "exemple", "(up, -1)"})
}

func TestAppliquerConvertirListeMotsEnUpMarqueurDebut(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"(up)", "Voici", "un", "exemple"}, []string{"(up)", "Voici", "un", "exemple"})
}

func TestAppliquerConvertirListeMotsEnUpMarqueurInvalide(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "un", "exemple", "(up, 3x)"}, []string{"Voici", "un", "exemple", "(up, 3x)"})
}

func TestAppliquerConvertirListeMotsEnUpAvecIndiceZero(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "un", "exemple", "(up, 0)"}, []string{"Voici", "un", "exemple", "(up, 0)"})
}

func TestAppliquerConvertirListeMotsEnUpPlusieursMarqueurs(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "(up)", "un", "exemple", "(up, 2)"}, []string{"VOICI", "UN", "EXEMPLE"})
}

// ================================= //
// =========== TESTS LOW =========== //
// ================================= //

// Fonction générale pour tester ConvertirListeMotsEnLow //
func VerifierConvertirListeMotsEnLow(t *testing.T, entree []string, attendu []string) {
	resultat := regles.ConvertirListeMotsEnLow(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestAppliquerConvertirListeMotsEnLowSansIndice(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "UN", "EXEMPLE", "(low)"}, []string{"VOICI", "UN", "exemple"})
}

func TestAppliquerConvertirListeMotsEnLowAvecIndice(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 2)"}, []string{"VOICI", "un", "exemple"})
}

func TestAppliquerConvertirListeMotsEnLowAvecIndiceTropLong(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 6)"}, []string{"VOICI", "UN", "EXEMPLE", "(low, 6)"})
}

func TestAppliquerConvertirListeMotsEnLowAvecIndiceNegatif(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "UN", "EXEMPLE", "(low, -1)"}, []string{"VOICI", "UN", "EXEMPLE", "(low, -1)"})
}

func TestAppliquerConvertirListeMotsEnLowMarqueurDebut(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"(low)", "VOICI", "UN", "EXEMPLE"}, []string{"(low)", "VOICI", "UN", "EXEMPLE"})
}

func TestAppliquerConvertirListeMotsEnLowMarqueurInvalide(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 3x)"}, []string{"VOICI", "UN", "EXEMPLE", "(low, 3x)"})
}

func TestAppliquerConvertirListeMotsEnLowAvecIndiceZero(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 0)"}, []string{"VOICI", "UN", "EXEMPLE", "(low, 0)"})
}

func TestAppliquerConvertirListeMotsEnLowPlusieursMarqueurs(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "(low)", "UN", "EXEMPLE", "(low, 2)"}, []string{"voici", "un", "exemple"})
}

// ================================= //
// =========== TESTS CAP =========== //
// ================================= //

// Fonction générale pour tester ConversionCap //
func VerifierConversionCap(t *testing.T, entree string, attendu string) {
	resultat := regles.ConversionCap(entree)
	if resultat != attendu {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestConversionCapAvecAccent(t *testing.T) {
	VerifierConversionCap(t, "élan", "Élan")
}

func TestConversionCapMotAvecMaj(t *testing.T) {
	VerifierConversionCap(t, "bRIDGE", "Bridge")
}

func TestConversionCapCasNormal(t *testing.T) {
	VerifierConversionCap(t, "bridge", "Bridge")
}

func TestConversionCapMotVide(t *testing.T) {
	VerifierConversionCap(t, "", "")
}

// Fonction générale pour tester ConvertirListeMotsEnCap //
func VerifierConvertirListeMotsEnCap(t *testing.T, entree []string, attendu []string) {
	resultat := regles.ConvertirListeMotsEnCap(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestConvertirListeMotsEnCapSansIndice(t *testing.T) {
	VerifierConvertirListeMotsEnCap(t, []string{"Bonjour", "voici", "un", "exemple", "(cap)"}, []string{"Bonjour", "voici", "un", "Exemple"})
}

func TestConvertirListeMotsEnCapPlusieursIndices(t *testing.T) {
	VerifierConvertirListeMotsEnCap(t, []string{"Bonjour", "vOICI", "(cap, 1)", "UN", "exemple", "(cap, 2)"}, []string{"Bonjour", "Voici", "Un", "Exemple"})
}

func TestConvertirListeMotsEnCapIndiceInvalide(t *testing.T) {
	VerifierConvertirListeMotsEnCap(t, []string{"Bonjour", "voici", "un", "exemple", "(cap, -2)"}, []string{"Bonjour", "voici", "un", "exemple", "(cap, -2)"})
}
