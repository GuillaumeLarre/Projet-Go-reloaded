package regles

import (
	"fmt"
	"strconv"
	"strings"
)

// ================================== //
// =========== CONVERSION =========== //
// ================================== //

// Fonction globale pour parcourir des listes de mots //
// Et appliquer les conversions voulues //
func Conversion(listeMots []string, marqueur string, base int) []string {
	var i int
	nouvelleListe := []string{}
	for i < len(listeMots) {
		if listeMots[i] == marqueur && i == 0 {
			fmt.Printf("Erreur : La conversion ne peut être faite que si %s précède un mot\n", marqueur)
			nouvelleListe = append(nouvelleListe, listeMots[i])
			i++
		} else if i+1 < len(listeMots) && listeMots[i+1] == marqueur {
			conversion, err := strconv.ParseInt(listeMots[i], base, 64)
			texte := strconv.FormatInt(conversion, 10)
			if err != nil {
				fmt.Printf("Erreur: La conversion n'est pas posssible, le mot qui précède %s n'est pas convertible\n", marqueur)
				nouvelleListe = append(nouvelleListe, listeMots[i])
				nouvelleListe = append(nouvelleListe, listeMots[i+1])
				i += 2
			} else {
				nouvelleListe = append(nouvelleListe, texte)
				i += 2
			}
		} else {
			nouvelleListe = append(nouvelleListe, listeMots[i])
			i++
		}
	}
	return nouvelleListe
}

// Convertie un hexadécimal en nombre //
func ConvertirHexaEnNombre(listeMots []string) []string {
	return Conversion(listeMots, "(hex)", 16)
}

// Convertie un binaire en nombre //
func ConvertirBinEnNombre(listeMots []string) []string {
	return Conversion(listeMots, "(bin)", 2)
}

// Uniformise l'écriture de différents marqueurs //
// DecoupeEnMots découpe à chaque espace et peut séparer un indice en deux //
func Uniformisation(listeMots []string, marqueur string) []string {
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

// Appel l'uniformisation pour tous les marqueurs //
func UniformiserMarqueurs(mots []string) []string {
	nouvelleListe := Uniformisation(mots, "(up,")
	nouvelleListe = Uniformisation(nouvelleListe, "(low,")
	nouvelleListe = Uniformisation(nouvelleListe, "(cap,")
	return nouvelleListe
}

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

// Fonction permettant de convertir avec le marqueur (up) //
func ConvertirListeMotsEnUp(listeMots []string) []string {
	return AppliquerMarqueurCasse(listeMots, "(up", strings.ToUpper)
}

// Fonction permettant de convertir le marqueur (low) //
func ConvertirListeMotsEnLow(listeMots []string) []string {
	return AppliquerMarqueurCasse(listeMots, "(low", strings.ToLower)
}

// Fonction permettant de convertir le marqueur (cap) //
func ConvertirListeMotsEnCap(listeMots []string) []string {
	return AppliquerMarqueurCasse(listeMots, "(cap", ConversionCap)
}

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
					estOuvrant = true
					continue
				} else if len(listeMots[i]) > 1 && ponct != "" {
					nouvelleListe = append(nouvelleListe, ponct+listeMots[i])
					ponct = ""
					estOuvrant = true
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
