package regles

import (
	"fmt"
	"strconv"
)

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

func ConvertirHexaEnNombre(listeMots []string) []string {
	return Conversion(listeMots, "(hex)", 16)
}

func ConvertirBinEnNombre(listeMots []string) []string {
	return Conversion(listeMots, "(bin)", 2)
}

func Fusion(listeMots []string, marqueur string) []string {
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
