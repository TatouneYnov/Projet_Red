package red

import "fmt"

func displayCombatInterface(c *Character, m *Monster, tourCombat int, isTraining bool) {
	playerHealthPercent := float64(c.PvActuelle) / float64(c.PvMax) * 100
	monsterHealthPercent := float64(m.PvActuelle) / float64(m.PvMax) * 100

	playerHealthBar := ""
	monsterHealthBar := ""

	for i := 0; i < 20; i++ {
		if float64(i)*5 < playerHealthPercent {
			playerHealthBar += "█"
		} else {
			playerHealthBar += "░"
		}

		if float64(i)*5 < monsterHealthPercent {
			monsterHealthBar += "█"
		} else {
			monsterHealthBar += "░"
		}
	}

	DisplayCombatHeader(tourCombat, isTraining)

	fmt.Println("\n┌─────────────────────────────┐  ┌─────────────────────────────┐")
	fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │  👹 \033[1m%-10s\033[0m              │\n", c.Nom, c.Niveau, m.Nom)
	fmt.Printf("│    \033[32m%-20s\033[0m     │  │    \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
	fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │    PV: %-3d/%-3d (%.0f%%)        │\n",
		c.PvActuelle, c.PvMax, playerHealthPercent,
		m.PvActuelle, m.PvMax, monsterHealthPercent)
	fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║            📜 DÉROULEMENT 📜             ║")
	fmt.Println("╚══════════════════════════════════════════╝")
}
