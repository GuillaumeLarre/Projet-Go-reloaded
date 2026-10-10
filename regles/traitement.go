package regles

// Permet de traiter plusieurs lignes séparemment //
func Traitement(texte string) string {
	lignes := DecoupeEnLignes(texte)
	for i := 0; i < len(lignes); i++ {
		mots := DecoupeEnMots(lignes[i])
		mots = UniformiserMarqueurs(mots)
		mots = ConvertirHexaEnNombre(mots)
		mots = ConvertirBinEnNombre(mots)
		mots = ConvertirListeMotsEnUp(mots)
		mots = ConvertirListeMotsEnLow(mots)
		mots = ConvertirListeMotsEnCap(mots)
		mots = TraitementPonctuation(mots)
		mots = TraitementAEtAn(mots)
		mots = TraitementGuillemets(mots)
		lignes[i] = AssembleEnTexte(mots)
	}
	texteTraite := AssembleEnLignes(lignes)
	return texteTraite
}
