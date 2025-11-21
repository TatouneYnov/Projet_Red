package red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SpellBook(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║            📚 LIVRE DE SORT 📚           ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	fmt.Println("Choisissez le sort que vous souhaitez apprendre :")
	fmt.Println("[1] 🔥 Hadouken (Street Fighter) - Attaque de feu")
	fmt.Println("[2] ⚡ Éclair de Zeus (God of War) - Attaque électrique")
	fmt.Println("[3] 🧊 Blizzard (Final Fantasy) - Attaque de glace")
	fmt.Println("[4] 💫 FUS RO DAH (Skyrim) - Souffle puissant")
	fmt.Println("[5] 🌪️ Tornade (Zelda) - Attaque de vent")

	fmt.Print("Votre choix : ")
	input, _ := reader.ReadString('\n')
	choix := strings.TrimSpace(input)

	switch choix {
	case "1":
		learnSpell(c, "Hadouken (Street Fighter)")
	case "2":
		learnSpell(c, "Éclair de Zeus (God of War)")
	case "3":
		learnSpell(c, "Blizzard (Final Fantasy)")
	case "4":
		learnSpell(c, "FUS RO DAH (Skyrim)")
	case "5":
		learnSpell(c, "Tornade (Zelda)")
	default:
		fmt.Println("❌ Choix invalide. Aucun sort appris.")
	}
}

func learnSpell(c *Character, spellName string) {
	for _, spell := range c.Skill {
		if spell == spellName {
			fmt.Printf("� Vous connaissez déjà le sort « %s ».\n", spellName)
			return
		}
	}

	c.Skill = append(c.Skill, spellName)
	fmt.Printf("📘 Vous avez appris le sort « %s » !\n", spellName)
}

func UseSpecificSpellBook(c *Character, bookName string) {
	var spellToLearn string

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║            📚 LIVRE DE SORT 📚           ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	switch bookName {
	case "Livre de Sort : Hadouken":
		spellToLearn = "Hadouken (Street Fighter)"
		fmt.Println("🔥 Vous ouvrez le livre de Hadouken...")
		fmt.Println("   \"Technique de combat de rue utilisant l'énergie spirituelle\"")
	case "Livre de Sort : Éclair de Zeus":
		spellToLearn = "Éclair de Zeus (God of War)"
		fmt.Println("⚡ Vous ouvrez le livre d'Éclair de Zeus...")
		fmt.Println("   \"Invoquez la foudre des dieux de l'Olympe\"")
	case "Livre de Sort : Blizzard":
		spellToLearn = "Blizzard (Final Fantasy)"
		fmt.Println("🧊 Vous ouvrez le livre de Blizzard...")
		fmt.Println("   \"Magie élémentaire de glace dévastatrice\"")
	case "Livre de Sort : FUS RO DAH":
		spellToLearn = "FUS RO DAH (Skyrim)"
		fmt.Println("💫 Vous ouvrez le livre de FUS RO DAH...")
		fmt.Println("   \"Ancien cri de dragon aux pouvoirs dévastateurs\"")
	case "Livre de Sort : Tornade":
		spellToLearn = "Tornade (Zelda)"
		fmt.Println("🌪️ Vous ouvrez le livre de Tornade...")
		fmt.Println("   \"Manipulation des vents et des tempêtes\"")
	default:
		fmt.Printf("❌ Livre inconnu : %s\n", bookName)
		return
	}

	fmt.Printf("🔍 DEBUG: spellToLearn = '%s'\n", spellToLearn)
	fmt.Printf("\n📖 Souhaitez-vous apprendre « %s » ? (o/n) : ", spellToLearn)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(strings.ToLower(input))

	if choice == "o" || choice == "oui" {
		learnSpell(c, spellToLearn)
	} else {
		fmt.Println("📚 Vous refermez le livre sans rien apprendre.")
	}
}
