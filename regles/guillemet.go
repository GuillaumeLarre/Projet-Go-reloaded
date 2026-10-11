package regles

// ================================== //
// =========== GUILLEMETS =========== //
// ================================== //

// Fonction permettant d'appliquer les règles de guillemet //
func TraitementGuillemets(listeMots []string) []string {
	guillemet := map[string]bool{
		"'":  true,
		"\"": true,
	}
	nouvelleListe := []string{}
	estOuvrant := false
	ponct := ""
	for i := 0; i < len(listeMots); i++ {
		if guillemet[string(listeMots[i][0])] {
			if estOuvrant == false {
				if len(listeMots[i]) > 1 && ponct == "" {
					nouvelleListe = append(nouvelleListe, listeMots[i])
					if guillemet[string(listeMots[i][len(listeMots[i])-1])] {
						estOuvrant = false
						continue
					}
					estOuvrant = true
					continue
				} else if len(listeMots[i]) > 1 && ponct != "" {
					nouvelleListe = append(nouvelleListe, ponct+listeMots[i])
					ponct = ""
					estOuvrant = true
					continue
				} else {
					ponct = listeMots[i]
					continue
				}
			} else {
				nouvelleListe[len(nouvelleListe)-1] = nouvelleListe[len(nouvelleListe)-1] + listeMots[i]
				estOuvrant = false
				continue
			}
		}
		if estOuvrant == false && ponct != "" {
			if guillemet[string(listeMots[i][len(listeMots[i])-1])] {
				nouvelleListe = append(nouvelleListe, ponct+listeMots[i])
				ponct = ""
			} else {
				nouvelleListe = append(nouvelleListe, ponct+listeMots[i])
				ponct = ""
				estOuvrant = true
			}
		} else if estOuvrant && guillemet[string(listeMots[i][len(listeMots[i])-1])] {
			nouvelleListe = append(nouvelleListe, listeMots[i])
			estOuvrant = false
		} else {
			nouvelleListe = append(nouvelleListe, listeMots[i])
		}
	}
	if ponct != "" {
		nouvelleListe = append(nouvelleListe, ponct)
	}
	return nouvelleListe
}
