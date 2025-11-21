package red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func InitCharacter(nom string, classe string, niveau int, points_de_vie_max int, points_de_vie_actuelle int, potions int) Character {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║         📊 PERSONNAGE CRÉÉ 📊            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Printf("📛 Nom          : %s\n", nom)
	fmt.Printf("🎭 Classe       : %s\n", classe)
	fmt.Printf("⭐ Niveau       : %d\n", niveau)
	fmt.Printf("❤️  Points de vie : %d/%d\n", points_de_vie_actuelle, points_de_vie_max)
	fmt.Printf("🧪 Potions      : %d\n", potions)
	fmt.Println("⚔️  Sorts appris : Coup de Poing")

	return Character{
		Nom:          nom,
		Classe:       classe,
		Niveau:       niveau,
		PvMax:        points_de_vie_max,
		PvActuelle:   points_de_vie_actuelle,
		ExpActuelle:  0,
		ExpMax:       100,
		ExpTotal:     0,
		Force:        5,
		Agilite:      5,
		Intelligence: 5,
		PointsStats:  0,
		Inventaire:   map[string]int{"Potion de vie": potions},
		Skill:        []string{"Coup de Poing"},
		Argent:       100,
		Equipement: Equipment{
			Tete:       "Aucun",
			Torse:      "Vêtements simples",
			Pieds:      "Sandales usées",
			Arme:       "Mains nues",
			Initiative: 6,
		},
		InventorySize: 10,
		UpgradesUsed:  0,
		Initiative:    6,
		MainQuest: MainQuest{
			Progression:       0,
			Completed:         false,
			FinalBossDefeated: false,
		},
		Mana:    50,
		ManaMax: 50,
	}
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatName(name string) string {
	if len(name) == 0 {
		return ""
	}
	name = strings.ToLower(name)
	return strings.ToUpper(string(name[0])) + name[1:]
}

func CharacterCreation() Character {
	reader := bufio.NewReader(os.Stdin)
	var nom string

	for {
		fmt.Print("Entrez le nom du personnage : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if isAlpha(input) && len(input) > 0 {
			nom = formatName(input)
			break
		} else {
			fmt.Println("Le nom doit contenir uniquement des lettres.")
		}
	}

	var classe string
	var pvMax int

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║           🎭 CHOIX DE CLASSE 🎭          ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  \033[36m[1]\033[0m 👨 Humain    (100 PV)               ║")
	fmt.Println("║  \033[32m[2]\033[0m 🧝 Elfe      (80 PV)                ║")
	fmt.Println("║  \033[33m[3]\033[0m ⚒️  Nain      (120 PV)               ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	for {
		fmt.Print("\n🔹 Choisissez votre classe (1-3) : ")
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			classe = "Humain"
			pvMax = 100
			fmt.Println("\n✨ Vous avez choisi la classe Humain !")
			fmt.Println("   Race équilibrée avec 100 points de vie.")
		case "2":
			classe = "Elfe"
			pvMax = 80
			fmt.Println("\n✨ Vous avez choisi la classe Elfe !")
			fmt.Println("   Race agile avec 80 points de vie.")
		case "3":
			classe = "Nain"
			pvMax = 120
			fmt.Println("\n✨ Vous avez choisi la classe Nain !")
			fmt.Println("   Race robuste avec 120 points de vie.")
		default:
			fmt.Println("❌ Choix invalide. Veuillez choisir 1, 2 ou 3.")
			continue
		}
		break
	}

	niveau := 1
	pvActuelle := pvMax / 2
	potions := 3

	return InitCharacter(nom, classe, niveau, pvMax, pvActuelle, potions)
}
