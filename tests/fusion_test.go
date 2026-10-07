package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

func VerificationFusion(t *testing.T, entree []string, marqueur string, attendu []string) {
	resultat := regles.Fusion(entree, marqueur)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestFusionMarqueurCoupeParEspace(t *testing.T) {
	VerificationFusion(t, []string{"(up,", "2)"}, "(up,", []string{"(up, 2)"})
}

func TestFusionMarqueurSansEspace(t *testing.T) {
	VerificationFusion(t, []string{"(up,2)"}, "(up,", []string{"(up, 2)"})
}

func TestFusionMarqueurSansIndice(t *testing.T) {
	VerificationFusion(t, []string{"go", "(up)"}, "(up,", []string{"go", "(up)"})
}

func TestFusionMarqueurOublieIndice(t *testing.T) {
	VerificationFusion(t, []string{"so", "exciting", "(up,"}, "(up,", []string{"so", "exciting", "(up,"})
}

func TestFusionSansMarqueur(t *testing.T) {
	VerificationFusion(t, []string{"(Voici,", "un", "exemple)"}, "(up,", []string{"(Voici,", "un", "exemple)"})
}
