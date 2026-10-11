package tests

import (
	"go-reloaded/regles"
	"testing"
)

// ======================================== //
// =========== TESTS TRAITEMENT =========== //
// ======================================== //

// Fonction générale pour tester //
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

// ======================================== //
// =========== EXEMPLES DU SUJET ========== //
// ======================================== //

func TestSujetPlusieursRegles(t *testing.T) {
	VerifierTraitement(t,
		"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
		"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.")
}

func TestSujetUsageHexBin(t *testing.T) {
	VerifierTraitement(t,
		"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
		"Simply add 66 and 2 and you will see the result is 68.")
}

func TestSujetUsageArticle(t *testing.T) {
	VerifierTraitement(t,
		"There is no greater agony than bearing a untold story inside you.",
		"There is no greater agony than bearing an untold story inside you.")
}

func TestSujetUsagePonctuation(t *testing.T) {
	VerifierTraitement(t,
		"Punctuation tests are ... kinda boring ,what do you think ?",
		"Punctuation tests are... kinda boring, what do you think?")
}

func TestSujetRegleHex(t *testing.T) {
	VerifierTraitement(t,
		"1E (hex) files were added",
		"30 files were added")
}

func TestSujetRegleBin(t *testing.T) {
	VerifierTraitement(t,
		"It has been 10 (bin) years",
		"It has been 2 years")
}

func TestSujetRegleUp(t *testing.T) {
	VerifierTraitement(t,
		"Ready, set, go (up) !",
		"Ready, set, GO!")
}

func TestSujetRegleLow(t *testing.T) {
	VerifierTraitement(t,
		"I should stop SHOUTING (low)",
		"I should stop shouting")
}

func TestSujetRegleCap(t *testing.T) {
	VerifierTraitement(t,
		"Welcome to the Brooklyn bridge (cap)",
		"Welcome to the Brooklyn Bridge")
}

func TestSujetRegleUpAvecNombre(t *testing.T) {
	VerifierTraitement(t,
		"This is so exciting (up, 2)",
		"This is SO EXCITING")
}

func TestSujetReglePonctuation(t *testing.T) {
	VerifierTraitement(t,
		"I was sitting over there ,and then BAMM !!",
		"I was sitting over there, and then BAMM!!")
}

func TestSujetReglePonctuationGroupe(t *testing.T) {
	VerifierTraitement(t,
		"I was thinking ... You were right",
		"I was thinking... You were right")
}

func TestSujetRegleGuillemetsUnMot(t *testing.T) {
	VerifierTraitement(t,
		"I am exactly how they describe me: ' awesome '",
		"I am exactly how they describe me: 'awesome'")
}

func TestSujetRegleGuillemetsPlusieursMots(t *testing.T) {
	VerifierTraitement(t,
		"As Elton John said: ' I am the most well-known homosexual in the world '",
		"As Elton John said: 'I am the most well-known homosexual in the world'")
}

func TestSujetRegleArticle(t *testing.T) {
	VerifierTraitement(t,
		"There it was. A amazing rock!",
		"There it was. An amazing rock!")
}

// ======================================================= //
// =========== GRILLE DE CORRECTION OFFICIELLE =========== //
// ======================================================= //

func TestGrilleCasse(t *testing.T) {
	VerifierTraitement(t,
		"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
		"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?")
}

func TestGrilleConversions(t *testing.T) {
	VerifierTraitement(t,
		"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
		"I have to pack 5 outfits. Packed 26 just to be sure")
}

func TestGrillePonctuation(t *testing.T) {
	VerifierTraitement(t,
		"Don not be sad ,because sad backwards is das . And das not good",
		"Don not be sad, because sad backwards is das. And das not good")
}

func TestGrilleGuillemetsEtArticle(t *testing.T) {
	VerifierTraitement(t,
		"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
		"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'")
}
