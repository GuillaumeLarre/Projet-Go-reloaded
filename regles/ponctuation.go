package regles

import "strings"

// ================================== //
// =========== PONCTUATION=========== //
// ================================== //

// Fonction permettant d'appliquer les règles de ponctuation //
func TraitementPonctuation(listeMots []string) []string {
	nouvelleListe := []string{}
	ponctuationCible := map[string]bool{
		",": true,
		".": true,
		"!": true,
		"?": true,
		":": true,
		";": true,
		"…": true,
	}
	for i := 0; i < len(listeMots); i++ {
		ponct := ""
		motprefix := listeMots[i]
		for _, r := range listeMots[i] {
			if ponctuationCible[string(r)] {
				ponct += string(r)
			} else {
				break
			}
		}
		if len(nouvelleListe) >= 1 && ponct != "" {
			nouvelleListe[len(nouvelleListe)-1] += ponct
			motprefix = strings.TrimPrefix(listeMots[i], ponct)
		}
		if motprefix != "" {
			nouvelleListe = append(nouvelleListe, motprefix)
		}
	}
	return nouvelleListe
}
