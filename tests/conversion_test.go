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

func VerifierConvertirListeMotsEnUp(t *testing.T, entree []string, attendu []string) {
	resultat := regles.ConvertirListeMotsEnUp(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestConvertirListeMotsEnUpCasNormal(t *testing.T) {
	VerifierConvertirListeMotsEnUp(t, []string{"Voici", "un", "exemple", "(up, 2)"}, []string{"Voici", "UN", "EXEMPLE"})
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

func VerifierConvertirListeMotsEnLow(t *testing.T, entree []string, attendu []string) {
	resultat := regles.ConvertirListeMotsEnLow(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Resultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestConvertirListeMotsEnLowCasNormal(t *testing.T) {
	VerifierConvertirListeMotsEnLow(t, []string{"VOICI", "UN", "EXEMPLE", "(low, 2)"}, []string{"VOICI", "un", "exemple"})
}

// ================================= //
// =========== TESTS CAP =========== //
// ================================= //

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

// =================================== //
// =========== PONCTUATION =========== //
// =================================== //

func VerifierTraitementPonctuation(t *testing.T, entree []string, attendu []string) {
	resultat := regles.TraitementPonctuation(entree)
	if !slices.Equal(resultat, attendu) {
		t.Errorf("Résultat : %q, Attendu : %q", resultat, attendu)
	}
}

func TestTraitementPonctuationCasNormal(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici,", "un", "exemple"}, []string{"Voici,", "un", "exemple"})
}

func TestTraitementPonctuationPonctuationApresMot(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici", ",un", "exemple"}, []string{"Voici,", "un", "exemple"})
}

func TestTraitementPonctuationPlusieursPonctuations(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici", ",,un", "exemple"}, []string{"Voici,,", "un", "exemple"})
}

func TestTraitementPonctuationPlusieursCas(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"Voici", "?un", "exemple", "."}, []string{"Voici?", "un", "exemple."})
}

func TestTraitementPonctuationDebutMot(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"!Voici", "un", "exemple"}, []string{"!Voici", "un", "exemple"})
}

func TestTraitementPonctuationMultiplesCas(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"%µVoici,", "un", "...exemple", "?", "fin", "!?"}, []string{"%µVoici,", "un...", "exemple?", "fin!?"})
}

func TestTraitementPonctuationSujet(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"then", "BAMM", "!!"}, []string{"then", "BAMM!!"})
}

func TestTraitementPonctuationSymbolleSpecifique(t *testing.T) {
	VerifierTraitementPonctuation(t, []string{"then", "BAMM", "…"}, []string{"then", "BAMM…"})
}

// ================================== //
// =========== GUILLEMETS =========== //
// ================================== //

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

// ============================ //
// =========== A/AN =========== //
// ============================ //

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
