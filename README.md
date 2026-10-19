# go-reloaded

Projet scolaire d'édition et d'autocorrection de texte écrit en Go, avec uniquement la bibliothèque standard.

Le programme lit un fichier texte, applique une série de règles de transformation, puis écrit le texte corrigé dans un second fichier.

## Utilisation

```
go run . <fichier_entree> <fichier_sortie>
```

Pour lancer les tests :

```
go test ./...
```

## Règles à suivre

| Marqueur | Effet |
| --- | --- |
| `(hex)` | convertit le mot précédent d'hexadécimal en décimal |
| `(bin)` | convertit le mot précédent de binaire en décimal |
| `(up)` / `(up, N)` | majuscules sur le ou les N mots précédents |
| `(low)` / `(low, N)` | minuscules sur le ou les N mots précédents |
| `(cap)` / `(cap, N)` | majuscule initiale sur le ou les N mots précédents |
| Ponctuation `. , ! ? : ;` | collée au mot précédent, groupes conservés |
| Guillemets collés aux mots qu'ils entourent |
| `a` devient `an` devant une voyelle ou un `h` |

## Organisation du projet

```
go-reloaded/
├── main.go                     lecture des arguments et du fichier, Appelle du traitement, écriture
├── regles/                     package contenant toute la logique de transformation
│   ├── decoupe.go              découpage et recollage (lignes et mots)
│   ├── guillemet.go            règles des guillemets
│   ├── marqueuravecindice.go   règles (up), (low) et (cap)
│   ├── nombre.go               règles (hex) et (bin)
│   ├── ponctuation.go          règles de ponctuation
│   ├── regleanglais.go         règle a / an
│   └── traitement.go           orchestre l'ensemble des règles de conversion
└── tests/                      tests unitaires, un fichier par thème
```

`main.go` ne contient aucune logique de transformation : il lit, appelle `regles.Traitement`, puis écrit. Toute la logique est dans le package `regles`, ce qui permet de la tester entièrement sans passer par des fichiers.

## Fonctionnement : une chaîne de traitement

Le texte est d'abord découpé en lignes. Chaque ligne est traitée **indépendamment**, puis les lignes sont recollées. Pour chaque ligne :

```
découpage en mots
  → uniformisation des marqueurs (up, low, cap)
  → (hex) → (bin)
  → (up) → (low) → (cap)
  → ponctuation
  → a / an
  → guillemets
  → recollage des mots
```

Chaque règle reçoit une liste de mots et renvoie une nouvelle liste de mots (`[]string → []string`). Elles peuvent donc s'enchaîner librement, et chacune peut être testée seule.

L'ordre suit une règle simple : **une étape qui prépare les données passe avant toutes les étapes qui en dépendent.**

- L'uniformisation passe en premier : selon les espaces tapés par l'utilisateur, un marqueur comme `(up, 2)` arrive après découpage en un ou deux morceaux (`"(up,"` + `"2)"`, ou `"(up,2)"`). Elle les ramène tous à une forme unique, `"(up, 2)"`, pour que les règles de casse n'aient qu'une seule forme à reconnaître.
- La ponctuation passe après toutes les règles à marqueur : sinon, dans `go (up) !`, le `!` se collerait au marqueur.
- Les guillemets passent après la ponctuation, pour que `raincoat . '` donne bien `raincoat.'`.
- `a` / `an` passe avant les guillemets, tant que le `a` est encore un mot isolé.

## Gestion des erreurs

Le programme ne plante jamais à cause du contenu du texte. Quand un marqueur ne peut pas être appliqué, il **reste dans le texte tel quel**, et un avertissement s'affiche dans le terminal. Le fichier de sortie ne contient que du texte, jamais de message d'erreur.

Cas concernés :

- marqueur en début de ligne, sans mot à transformer : `(hex) files`, `(up) hello` ;
- mot qui n'est pas un nombre valide dans la base demandée : `hello (hex)`, `102 (bin)` ;
- nombre invalide, nul ou négatif : `(up, 3x)`, `(up, 0)`, `(up, -1)` ;
- nombre plus grand que le nombre de mots disponibles : `hello (up, 5)`.

## Choix d'interprétation du sujet

Le sujet laisse une part d’interprétation ; voici mes partis pris :

- **Lignes indépendantes** : une règle ne peut pas agir sur la ligne précédente. Les lignes vides sont conservées. Ce choix a été confirmé par un examinateur.
- **`(cap)`** : première lettre en majuscule, le reste du mot en minuscules (`bRIDGE` → `Bridge`).
- **`(up, N)` avec moins de N mots** : la règle n'est pas appliquée du tout, le marqueur reste avec un avertissement.
- **Groupes de ponctuation** : toute suite composée uniquement des signes `. , ! ? : ;` (et `…`) est considérée comme un groupe, et collée d'un bloc au mot précédent. Cela couvre `...`, `!?` et `!!`, tous présents dans le sujet.
- **Ponctuation en début de ligne** : laissée telle quelle, faute de mot précédent.
- **Guillemets** : `'` et `"` sont pris en charge. Un guillemet n'est reconnu qu'en première ou dernière position d'un mot ; une apostrophe au milieu d'un mot (`I'm`, `don't`) est ignorée. Les guillemets sont associés par paires, de gauche à droite.
- **`a` / `an`** : la majuscule est conservée (`A` → `An`). Le mot suivant est testé sur sa première lettre, en minuscule ou en majuscule.

## Limites connues et pistes d'amélioration

**Marqueurs**

- Seules les formes `(up, N)` et `(up,N)` sont reconnues. Des variantes comme `(up , 2)` ou `( up, 2)` ne le sont pas et restent dans le texte.
- L'uniformisation n'est pas idempotente : appliquée deux fois, elle ajouterait un espace en trop (`(up,  2)`). Elle ne doit être appelée qu'une fois, juste après le découpage en mots, ce qui est le cas dans la chaîne actuelle.

**Ponctuation**

- La ponctuation collée au milieu d'un mot (`there,and`) n'est pas séparée. La distinguer de cas légitimes comme `3.5` ou `e.g.` demanderait des règles supplémentaires.
- Un nombre qui commence par un point (`.5`) est coupé par la règle de ponctuation.

**Guillemets**

- Un guillemet orphelin (nombre impair sur une ligne) n'est pas géré : le sujet garantit qu'ils vont toujours par paires.
- Les guillemets mélangés (`'` fermé par `"`) ou imbriqués (une citation dans une citation) ne sont pas pris en charge.
- Les guillemets typographiques `‘ ’ “ ”` ne sont pas reconnus.
- Une apostrophe en bord de mot (`students'`, `'cause`) est prise pour un guillemet.

**`a` / `an`**

- Si un guillemet sépare le `a` de son mot (`a 'apple'`), la règle ne s'applique pas.
- La règle suit le sujet à la lettre : tout mot commençant par `h` déclenche `an`, y compris `house`. `an` n'est jamais transformé en `a`.

**Ordre d'application des marqueurs**

Les marqueurs sont traités par type, dans l'ordre de la chaîne ((hex), (bin), (up), (low), (cap)), et non dans l'ordre où ils apparaissent dans le texte. Pendant le passage d'une règle, les marqueurs des autres types sont encore présents dans la ligne et sont comptés comme des mots.

- `hello (cap) (up)` donne `hello (CAP)` au lieu de `HELLO` : la règle (up) transforme le marqueur (cap), qui n'est ensuite plus reconnu.
- `one two (low, 2) three (up, 3)` donne `one TWO (LOW, 2) THREE` au lieu de `ONE TWO THREE`.

La correction consisterait à remplacer les règles de marqueurs par une seule passe de gauche à droite, qui applique et retire chaque marqueur à sa position.

**Comptage des mots par les marqueurs**

Un marqueur avec nombre compte tous les éléments qui le précèdent, y compris la ponctuation isolée et les guillemets, qui ne sont rattachés aux mots qu'après les règles de marqueurs.

- `hello , world (up, 2)` donne `hello, WORLD` au lieu de `HELLO, WORLD`.
- `go ! (up)` transforme le `!` au lieu de `go`.

Ces cas n'apparaissent ni dans le sujet ni dans la grille de correction. Ils seront traités dans go-reforged, qui impose un traitement des modificateurs de gauche à droite et ne compte que les vrais mots.