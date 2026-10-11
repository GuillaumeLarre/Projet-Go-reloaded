package tests

import (
	"go-reloaded/regles"
	"slices"
	"testing"
)

// ======================================== //
// =========== TESTS GUILLEMETS =========== //
// ======================================== //

// Fonction générale pour tester TraitementGuillemets //
func VerifierTraitementGuillemets(t *testing.T, entree []string, attendu []string) {
	resultat := regles.TraitementGuillemets(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestTraitementGuillemetsCasNormal(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"'Aucun", "problèmes", "ici'"}, []string{"'Aucun", "problèmes", "ici'"})
}

func TestTraitementGuillemetsEspaceAvantEtApres(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"'", "Aucun", "problèmes", "ici", "'"}, []string{"'Aucun", "problèmes", "ici'"})
}

func TestTraitementGuillemetsEspaceAvant(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"'", "Aucun", "problèmes", "ici'"}, []string{"'Aucun", "problèmes", "ici'"})
}

func TestTraitementGuillemetPlusieursCasSepares(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"'", "hello", "'", "and", "'", "world", "'"}, []string{"'hello'", "and", "'world'"})
}

func TestTraitementGuillemetPlusieursCasMelange(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"'", "hello'", "and", "'world", "'"}, []string{"'hello'", "and", "'world'"})
}

func TestTraitementGuillemetUnSeul(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"I", "said", "yes", "'"}, []string{"I", "said", "yes", "'"})
}

func TestTraitementGuillemetColle(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"'hello", "world'", "and", "'", "x", "'"}, []string{"'hello", "world'", "and", "'x'"})
}

func TestTraitementGuillemetColleAvantApres(t *testing.T) {
	VerifierTraitementGuillemets(t, []string{"He", "said", "'yes'", "and", "'", "no", "'"}, []string{"He", "said", "'yes'", "and", "'no'"})
}
