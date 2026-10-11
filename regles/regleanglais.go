package regles

// =============================== //
// =========== A ET AN =========== //
// =============================== //

// Fonction permettant le traitement des règles a et an //
func TraitementAEtAn(listemots []string) []string {
	nouvelleListe := []string{}
	voyelles := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"h": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
		"H": true,
	}
	for i := 0; i < len(listemots); i++ {
		if i+1 < len(listemots) && listemots[i] == "a" && voyelles[string(listemots[i+1][0])] {
			nouvelleListe = append(nouvelleListe, "an")
		} else if i+1 < len(listemots) && listemots[i] == "A" && voyelles[string(listemots[i+1][0])] {
			nouvelleListe = append(nouvelleListe, "An")
		} else {
			nouvelleListe = append(nouvelleListe, listemots[i])
		}
	}
	return nouvelleListe
}
