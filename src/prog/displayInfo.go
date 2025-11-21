package red

import "fmt"

func DisplayInfo(c Character) {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║          👤 FICHE PERSONNAGE 👤          ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	healthPercent := float64(c.PvActuelle) / float64(c.PvMax) * 100
	healthBar := ""
	for i := 0; i < 20; i++ {
		if float64(i)*5 < healthPercent {
			healthBar += "█"
		} else {
			healthBar += "░"
		}
	}

	var healthColor string
	if healthPercent >= 70 {
		healthColor = "\033[32m"
	} else if healthPercent >= 30 {
		healthColor = "\033[33m"
	} else {
		healthColor = "\033[31m"
	}

	fmt.Println("\n🏷️  Informations générales :")
	fmt.Println("┌──────────────────────────────────────────┐")
	fmt.Printf("│  📛 Nom     : %-25s  │\n", c.Nom)
	fmt.Printf("│  ⚔️  Classe  : %-25s  │\n", c.Classe)
	fmt.Printf("│  ⭐ Niveau  : %-25d  │\n", c.Niveau)
	fmt.Printf("│  💰 Argent  : %-25d  │\n", c.Argent)
	fmt.Printf("│  🧠 Expérience : %d / %d                 │\n", c.ExpActuelle, c.ExpMax)
	fmt.Printf("│  ⚡ Initiative : %-21d   │\n", c.Initiative)
	fmt.Println("│  (Détermine qui commence le combat)      │")
	fmt.Printf("│  🔵 Mana     : %d / %d                   │\n", c.Mana, c.ManaMax)
	fmt.Println("└──────────────────────────────────────────┘")

	fmt.Println("\n❤️  Étatq de santé :")
	fmt.Println("┌──────────────────────────────────────────┐")
	fmt.Printf("│ %s%s\033[0m                     │\n", healthColor, healthBar)
	fmt.Printf("│ Points de vie : %s%d/%d (%.0f%%)\033[0m%s          │\n",
		healthColor, c.PvActuelle, c.PvMax, healthPercent,
		repeatSpaces(max(0, 15-len(fmt.Sprintf("%d/%d (%.0f%%)", c.PvActuelle, c.PvMax, healthPercent)))))
	fmt.Println("└──────────────────────────────────────────┘")

	fmt.Println("\n👕 Équipement :")
	fmt.Println("┌──────────────────────────────────────────┐")

	teteBonus := ""
	if c.Equipement.Tete == "Casquette de Mario" {
		teteBonus = " (+10❤️)"
	}
	fmt.Printf("│ 🧢 Tête   : %-25s%s    │\n", c.Equipement.Tete, teteBonus)

	torseBonus := ""
	if c.Equipement.Torse == "Tunique de Link" {
		torseBonus = " (+25❤️)"
	}
	fmt.Printf("│ 👚 Torse  : %-25s%s    │\n", c.Equipement.Torse, torseBonus)

	piedsBonus := ""
	bootBonus := GetBootsInitiativeBonus(c.Equipement.Pieds)
	if bootBonus > 0 {
		piedsBonus = fmt.Sprintf(" (+%d⚡)", bootBonus)
	}
	fmt.Printf("│ � Pieds  : %-25s%s     │\n", c.Equipement.Pieds, piedsBonus)

	armeBonus := ""
	weaponBonus := GetWeaponDamageBonus(c.Equipement.Arme)
	if weaponBonus > 0 {
		armeBonus = fmt.Sprintf(" (+%d💥)", weaponBonus)
	}
	fmt.Printf("│ ⚔️  Arme   : %-25s%s    │\n", c.Equipement.Arme, armeBonus)
	fmt.Println("└──────────────────────────────────────────┘")
}

func repeatSpaces(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += " "
	}
	return result
}
