package red

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func StatsMenu(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()

		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║          📊 STATISTIQUES 📊              ║")
		fmt.Println("╠══════════════════════════════════════════╣")

		fmt.Printf("║  💪 Force        : %-20d  ║\n", c.Force)
		fmt.Printf("║  🏃 Agilité      : %-20d  ║\n", c.Agilite)
		fmt.Printf("║  🧠 Intelligence : %-20d  ║\n", c.Intelligence)
		fmt.Println("╠══════════════════════════════════════════╣")

		if c.PointsStats > 0 {
			fmt.Printf("║  🔶 Points disponibles : \033[33m%-13d\033[0m ║\n", c.PointsStats)
			fmt.Println("╠══════════════════════════════════════════╣")
			fmt.Println("║                                          ║")
			fmt.Println("║  \033[36m[1]\033[0m 💪 Ajouter un point en Force      ║")
			fmt.Println("║  \033[36m[2]\033[0m 🏃 Ajouter un point en Agilité    ║")
			fmt.Println("║  \033[36m[3]\033[0m 🧠 Ajouter un point en Intelligence║")
			fmt.Println("║  \033[31m[4]\033[0m 🔙 Retourner au menu principal   ║")
		} else {
			fmt.Println("║  🔒 Aucun point disponible               ║")
			fmt.Println("╠══════════════════════════════════════════╣")
			fmt.Println("║                                          ║")
			fmt.Println("║  \033[31m[1]\033[0m 🔙 Retourner au menu principal      ║")
		}

		fmt.Println("║                                          ║")
		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║  💡 Chaque niveau gagné vous donne 3     ║")
		fmt.Println("║     points à distribuer dans vos stats.  ║")
		fmt.Println("║  📈 Bénéfices :                          ║")
		fmt.Println("║     - Force: Augmente vos dégâts         ║")
		fmt.Println("║     - Agilité: Chance d'esquive          ║")
		fmt.Println("║     - Intelligence: Efficacité des sorts ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Print("\n🔹 Votre choix: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		if c.PointsStats > 0 {
			switch choice {
			case "1":
				c.Force++
				c.PointsStats--
				fmt.Println("\n💪 Vous avez augmenté votre Force !")
				fmt.Println("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			case "2":
				c.Agilite++
				c.PointsStats--
				fmt.Println("\n🏃 Vous avez augmenté votre Agilité !")
				fmt.Println("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			case "3":
				c.Intelligence++
				c.PointsStats--
				fmt.Println("\n🧠 Vous avez augmenté votre Intelligence !")
				fmt.Println("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			case "4":
				return
			default:
				fmt.Println("\n❌ Choix invalide. Veuillez choisir entre 1 et 4.")
				fmt.Println("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			}
		} else {
			if choice == "1" {
				return
			} else {
				fmt.Println("\n❌ Choix invalide. Veuillez choisir 1 pour retourner au menu.")
				fmt.Println("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			}
		}
	}
}

func ApplyStatBonuses(c *Character, baseValue int, statType string) int {
	var finalValue int

	switch statType {
	case "damage":
		forceBonus := int(float64(baseValue) * (float64(c.Force) * 0.02))

		weaponBonus := GetWeaponDamageBonus(c.Equipement.Arme)

		finalValue = baseValue + forceBonus + weaponBonus
	case "dodge":
		finalValue = c.Agilite / 2
	case "spell":
		intBonus := int(float64(baseValue) * (float64(c.Intelligence) * 0.03))

		var weaponBonus int
		if c.Equipement.Arme == "Bâton Magique" {
			weaponBonus = GetWeaponDamageBonus(c.Equipement.Arme) / 2
		}

		finalValue = baseValue + intBonus + weaponBonus
	default:
		finalValue = baseValue
	}

	return finalValue
}

func TryDodge(c *Character) bool {

	dodgeChance := ApplyStatBonuses(c, 0, "dodge")

	if dodgeChance > 30 {
		dodgeChance = 30
	}

	roll := rand.Intn(100) + 1

	return roll <= dodgeChance
}
