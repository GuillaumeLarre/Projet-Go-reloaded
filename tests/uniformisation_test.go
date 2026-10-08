package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

// ========================================== //
// ========== TESTS UNIFORMISATION ========== //
// ========================================== //

// Tests de la fonction Uniformisation //
func VerificationUniformisation(t *testing.T, entree []string, marqueur string, attendu []string) {
	resultat := regles.Uniformisation(entree, marqueur)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestUniformisationMarqueurCoupeParEspace(t *testing.T) {
	VerificationUniformisation(t, []string{"(up,", "2)"}, "(up,", []string{"(up, 2)"})
}

func TestUniformisationMarqueurSansEspace(t *testing.T) {
	VerificationUniformisation(t, []string{"(up,2)"}, "(up,", []string{"(up, 2)"})
}

func TestUniformisationMarqueurSansIndice(t *testing.T) {
	VerificationUniformisation(t, []string{"go", "(up)"}, "(up,", []string{"go", "(up)"})
}

func TestUniformisationMarqueurOublieIndice(t *testing.T) {
	VerificationUniformisation(t, []string{"so", "exciting", "(up,"}, "(up,", []string{"so", "exciting", "(up,"})
}

func TestUniformisationSansMarqueur(t *testing.T) {
	VerificationUniformisation(t, []string{"(Voici,", "un", "exemple)"}, "(up,", []string{"(Voici,", "un", "exemple)"})
}

// Test de l'uniformisation des marqueurs //
func VerifierTestUniformiserMarqueurs(t *testing.T, entree []string, attendu []string) {
	resultat := regles.UniformiserMarqueurs(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestUniformiserMarqueursAvecEspace(t *testing.T) {
	VerifierTestUniformiserMarqueurs(t, []string{"so", "exciting", "(up,", "2)"}, []string{"so", "exciting", "(up, 2)"})
}

func TestUniformiserMarqueursSansEspace(t *testing.T) {
	VerifierTestUniformiserMarqueurs(t, []string{"so", "exciting", "(up,2)"}, []string{"so", "exciting", "(up, 2)"})
}

func TestUniformiserMarqueursDifferents(t *testing.T) {
	VerifierTestUniformiserMarqueurs(t, []string{"so", "low", "(low,", "2)", "and", "so", "(up,1)"}, []string{"so", "low", "(low, 2)", "and", "so", "(up, 1)"})
}

func TestUniformiserMarqueursSansNombre(t *testing.T) {
	VerifierTestUniformiserMarqueurs(t, []string{"so", "low", "(low)"}, []string{"so", "low", "(low)"})
}
