package regles

import (
	"fmt"
	"strconv"
	"strings"
)

// ============================================= //
// =========== MARQUEURS AVEC INDICE =========== //
// ============================================= //

// Uniformise l'écriture de différents marqueurs //
// DecoupeEnMots découpe à chaque espace et peut séparer un indice en deux //
func UniformisationMarqueurs(listeMots []string, marqueur string) []string {
	nouvelleListe := []string{}
	i := 0
	for i < len(listeMots) {
		if i+1 < len(listeMots) && listeMots[i] == marqueur {
			nouvelleListe = append(nouvelleListe, AssembleEnTexte([]string{listeMots[i], listeMots[i+1]}))
			i += 2
		} else if len(listeMots[i]) > len(marqueur) && listeMots[i][:len(marqueur)] == marqueur {
			nouvelleListe = append(nouvelleListe, AssembleEnTexte([]string{listeMots[i][:len(marqueur)], listeMots[i][len(marqueur):]}))
			i++
		} else {
			nouvelleListe = append(nouvelleListe, listeMots[i])
			i++
		}
	}
	return nouvelleListe
}

// Permet d'éviter de perdre une ponctuation collée au marqueur ex : (bin)! //
func UniformisationApresMarqueur(listeMots []string, marqueur string) []string {
	nouvelleListe := []string{}
	for i := 0; i < len(listeMots); i++ {
		position := strings.Index(listeMots[i], ")")
		if len(listeMots[i]) > len(marqueur) && position >= 0 && listeMots[i][:len(marqueur)] == marqueur {
			nouvelleListe = append(nouvelleListe, listeMots[i][:position+1])
			if listeMots[i][position+1:] != "" {
				nouvelleListe = append(nouvelleListe, listeMots[i][position+1:])
			}
		} else {
			nouvelleListe = append(nouvelleListe, listeMots[i])
		}
	}
	return nouvelleListe
}

// Appel l'uniformisation pour tous les marqueurs //
func UniformiserMarqueurs(mots []string) []string {
	nouvelleListe := UniformisationMarqueurs(mots, "(up,")
	nouvelleListe = UniformisationMarqueurs(nouvelleListe, "(low,")
	nouvelleListe = UniformisationMarqueurs(nouvelleListe, "(cap,")
	nouvelleListe = UniformisationApresMarqueur(nouvelleListe, "(up,")
	nouvelleListe = UniformisationApresMarqueur(nouvelleListe, "(low,")
	nouvelleListe = UniformisationApresMarqueur(nouvelleListe, "(cap,")
	return nouvelleListe
}

// Fonction globale pour appliquer le traitement des marqueurs Up, Low ou Cap //
func AppliquerMarqueurCasse(liste []string, marqueur string, fonction func(string) string) []string {
	nouvelleListe := []string{}
	i := 0
	for i < len(liste) {
		if len(liste[i]) > len(marqueur) && liste[i][:len(marqueur)] == marqueur {
			if liste[i][len(marqueur)] == ')' {
				if len(nouvelleListe) > 0 {
					nouvelleListe[len(nouvelleListe)-1] = fonction(nouvelleListe[len(nouvelleListe)-1])
					i++
				} else {
					fmt.Println("Erreur : la conversion n'est pas possible, aucun mot à convertir")
					nouvelleListe = append(nouvelleListe, liste[i])
					i++
					continue
				}
			} else if liste[i][len(marqueur)] != ')' {
				prefixe := strings.TrimPrefix(liste[i], marqueur+", ")
				suffixe := strings.TrimSuffix(prefixe, ")")
				n, err := strconv.Atoi(suffixe)
				if err != nil {
					fmt.Println("Erreur :", err)
					nouvelleListe = append(nouvelleListe, liste[i])
					i++
					continue
				}
				if n > len(nouvelleListe) || n <= 0 {
					fmt.Printf("Erreur : La conversion n'est pas possible l'indice n'a pas une valeur correcte (%d)\n", n)
					nouvelleListe = append(nouvelleListe, liste[i])
					i++
					continue
				}
				for indice := 1; indice <= n; indice++ {
					nouvelleListe[len(nouvelleListe)-indice] = fonction(nouvelleListe[len(nouvelleListe)-indice])
				}
				i++
			}
		} else {
			nouvelleListe = append(nouvelleListe, liste[i])
			i++
		}
	}
	return nouvelleListe
}

// ========================== //
// =========== UP =========== //
// ========================== //

// Convertie le marqueur (up) //
func ConvertirListeMotsEnUp(listeMots []string) []string {
	return AppliquerMarqueurCasse(listeMots, "(up", strings.ToUpper)
}

// =========================== //
// =========== LOW =========== //
// =========================== //

// Convertie le marqueur (low) //
func ConvertirListeMotsEnLow(listeMots []string) []string {
	return AppliquerMarqueurCasse(listeMots, "(low", strings.ToLower)
}

// =========================== //
// =========== CAP =========== //
// =========================== //

// Fonction pour appliquer (cap) sur un mot //
func ConversionCap(mot string) string {
	if mot == "" {
		return ""
	}
	motLower := strings.ToLower(mot)
	nouveauMot := ""
	for i, r := range motLower {
		if i == 0 {
			nouveauMot += strings.ToUpper(string(r))
		} else {
			nouveauMot += string(r)
		}
	}
	return nouveauMot
}

// Converie le marqueur (cap) //
func ConvertirListeMotsEnCap(listeMots []string) []string {
	return AppliquerMarqueurCasse(listeMots, "(cap", ConversionCap)
}
