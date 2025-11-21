package red

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CharacterTurn(c *Character, m *Monster, reader *bufio.Reader, tourCombat int, isTraining bool) int {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║           🎮 VOS ACTIONS 🎮              ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Printf("║   🔵 Mana : %d / %d%22s║\n", c.Mana, c.ManaMax, "")
	fmt.Println("║   \033[36m[1]\033[0m 👊 Attaquer                        ║")
	fmt.Println("║   \033[32m[2]\033[0m 🎒 Inventaire                      ║")

	if len(c.Skill) > 0 {
		fmt.Println("║   \033[33m[3]\033[0m 🔮 Utiliser un sort                ║")
	}

	fmt.Println("║   \033[31m[4]\033[0m 🚪 Quitter le combat               ║")

	fmt.Println("╚══════════════════════════════════════════╝")

	FlushInputBuffer(reader)

	fmt.Print("\nVotre choix : ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erreur de lecture : ", err)
		return -1
	}

	choice := strings.TrimSpace(input)

	switch choice {
	case "1":
		baseDamage := 8
		var dommage int
		if c.Mana > 0 {
			dommage = ApplyStatBonuses(c, baseDamage, "damage")
			c.Mana -= 1
			fmt.Printf("🔵 Mana restant : %d / %d\n", c.Mana, c.ManaMax)
			if m.Nom != "Bokoblin Rouge (Zelda)" && m.Nom != "Sandbag Smash Bros" && rand.Intn(100) < 5 {
				ClearScreen()
				if isTraining {
					displayCombatInterface(c, m, tourCombat, true)
				} else {
					DisplayCombatHeader(tourCombat, false)
					DisplayHealthBarsQuest(c, m)
				}
				fmt.Println("\n┌──────────────────────────────────────────┐")
				fmt.Printf("│ 👊 \033[1m%s\033[0m tente d'attaquer              │\n", c.Nom)
				fmt.Printf("│ 💨 \033[33mMais %s esquive l'attaque !\033[0m     │\n", m.Nom)
				fmt.Println("└──────────────────────────────────────────┘")
			} else {
				m.PvActuelle -= dommage
				if m.PvActuelle < 0 {
					m.PvActuelle = 0
				}
				ClearScreen()
				if isTraining {
					displayCombatInterface(c, m, tourCombat, true)
				} else {
					DisplayCombatHeader(tourCombat, false)
					DisplayHealthBarsQuest(c, m)
				}
				fmt.Println("\n┌──────────────────────────────────────────┐")
				fmt.Printf("│ � \033[1m%s\033[0m attaque                              │\n", c.Nom)
				if dommage > baseDamage {
					bonusDamage := dommage - baseDamage
					fmt.Printf("│ 💥 Inflige \033[31m%d\033[0m points de dégâts (\033[33m+%d\033[0m)      │\n", dommage, bonusDamage)
				} else {
					fmt.Printf("│ 💥 Inflige \033[31m%d\033[0m points de dégâts            │\n", dommage)
				}
				fmt.Println("└──────────────────────────────────────────┘")
			}
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
			return 1
		} else {
			fmt.Println("❌ Pas assez de mana pour Coup de poing !")
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
			return 0
		}

	case "2":
		inventoryUsed := ShowCombatInventory(c, m, reader, tourCombat, isTraining)
		if inventoryUsed {
			return 1
		} else {
			return 0
		}

	case "3":
		if len(c.Skill) > 0 {
			spellResult := UseSpell(c, m, reader, tourCombat, isTraining)
			if spellResult {
				return 1
			} else {
				return 0
			}
		} else {
			fmt.Println("\n❌ Vous n'avez pas encore appris de sorts !")
			return -1
		}

	case "4":
		fmt.Println("\n📌 Vous décidez de quitter le combat.")

		if tourCombat > 0 && m.Nom == "Sandbag Smash Bros" {
			fmt.Println("\nℹ️ Le mannequin est un outil d'entraînement uniquement et ne donne ni expérience ni récompense.")
			fmt.Println("   Vous avez pratiqué vos techniques de combat pendant " + fmt.Sprintf("%d", tourCombat) + " tours.")
		} else {
			fmt.Println("\nℹ️ Vous vous éloignez prudemment du combat.")

			if strings.Contains(m.Nom, "Boss") && m.Nom != "Sandbag Smash Bros" {
				fmt.Println("\n⚠️ Fuir un boss n'est pas sans conséquence...")
				damage := c.PvMax / 10
				if damage < 1 {
					damage = 1
				}
				c.PvActuelle -= damage
				if c.PvActuelle < 1 {
					c.PvActuelle = 1
				}
				fmt.Printf("\n💥 Vous subissez %d points de dégâts en fuyant !\n", damage)
			}
		}

		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		reader.ReadString('\n')

		return 2

	default:
		fmt.Println("\n❌ Choix invalide ! Veuillez choisir une option valide.")
		return -1
	}
}

func ShowCombatInventory(c *Character, m *Monster, reader *bufio.Reader, tourCombat int, isTraining bool) bool {
	if len(c.Inventaire) == 0 {
		fmt.Println("\n🎒 Votre inventaire est vide !")
		fmt.Print("Appuyez sur Entrée pour revenir au combat...")
		FlushInputBuffer(reader)
		reader.ReadString('\n')

		ClearScreen()
		return false
	}

	ClearScreen()

	DisplayCombatHeader(tourCombat, isTraining)

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

	fmt.Println("┌─────────────────────────────┐  ┌─────────────────────────────┐")
	fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │ 👹 \033[1m%-10s\033[0m           │\n", c.Nom, c.Niveau, m.Nom)
	fmt.Printf("│    \033[32m%-20s\033[0m     │  │    \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
	fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │    PV: %-3d/%-3d (%.0f%%)        │\n",
		c.PvActuelle, c.PvMax, playerHealthPercent,
		m.PvActuelle, m.PvMax, monsterHealthPercent)
	fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║           🎒 INVENTAIRE 🎒               ║")
	fmt.Println("╠══════════════════════════════════════════╣")

	items := make([]string, 0, len(c.Inventaire))
	for item := range c.Inventaire {
		items = append(items, item)
	}

	for i, item := range items {
		fmt.Printf("║  \033[33m[%d]\033[0m %s (x%d)\n", i+1, item, c.Inventaire[item])
	}

	fmt.Println("║  \033[31m[0]\033[0m Retour au combat")
	fmt.Println("╚══════════════════════════════════════════╝")

	fmt.Print("\nChoix (numéro de l'objet, 0 pour retour) : ")

	FlushInputBuffer(reader)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erreur de lecture : ", err)
		return false
	}

	choice := strings.TrimSpace(input)
	if choice == "0" {
		ClearScreen()
		return false
	}

	var index int
	fmt.Sscanf(choice, "%d", &index)

	if choice != "0" && (index < 1 || index > len(items)) {
		fmt.Println("\n❌ Choix invalide !")
		time.Sleep(500 * time.Millisecond)
		ClearScreen()
		return ShowCombatInventory(c, m, reader, tourCombat, isTraining)
	}

	selectedItem := items[index-1]
	return UseItemInCombat(c, m, selectedItem, reader, tourCombat)
}

func UseItemInCombat(c *Character, m *Monster, itemName string, reader *bufio.Reader, tourCombat int) bool {
	quantity, exists := c.Inventaire[itemName]
	if !exists || quantity <= 0 {
		fmt.Println("\n❌ Cet objet n'est plus disponible !")
		return false
	}

	switch itemName {
	case "Potion de vie":
		soin := 20
		ancienPv := c.PvActuelle
		c.PvActuelle += soin
		if c.PvActuelle > c.PvMax {
			c.PvActuelle = c.PvMax
		}

		ClearScreen()
		DisplayCombatHeader(tourCombat, true)
		fmt.Println("\n┌──────────────────────────────────────────┐")
		fmt.Printf("│ 🧪 \033[1m%s\033[0m utilise \033[32mPotion de vie\033[0m        │\n", c.Nom)
		fmt.Printf("│ ❤️  Récupère \033[32m%d\033[0m points de vie          │\n", c.PvActuelle-ancienPv)
		fmt.Println("└──────────────────────────────────────────┘")
	case "Potion de mana", "Potion Bleue de Mana":
		soinMana := 30
		ancienMana := c.Mana
		c.Mana += soinMana
		if c.Mana > c.ManaMax {
			c.Mana = c.ManaMax
		}
		ClearScreen()
		DisplayCombatHeader(tourCombat, true)
		fmt.Println("\n┌──────────────────────────────────────────┐")
		fmt.Printf("│ 🔵 %s utilise %s           │\n", c.Nom, itemName)
		fmt.Printf("│ 🔋 Récupère %d points de mana          │\n", c.Mana-ancienMana)
		fmt.Println("└──────────────────────────────────────────┘")

		ClearScreen()

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

		DisplayCombatHeader(tourCombat, true)

		fmt.Println("\n┌─────────────────────────────┐  ┌─────────────────────────────┐")
		fmt.Printf("│ 🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │ 👹 \033[1m%-10s\033[0m  │\n", c.Nom, c.Niveau, m.Nom)
		fmt.Printf("│ \033[32m%-20s\033[0m     │  │ \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
		fmt.Printf("│ PV: %-3d/%-3d (%.0f%%)        │  │   PV: %-3d/%-3d (%.0f%%)        │\n",
			c.PvActuelle, c.PvMax, playerHealthPercent,
			m.PvActuelle, m.PvMax, monsterHealthPercent)
		fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

	case "Poison":
		dommage := 15
		m.PvActuelle -= dommage
		if m.PvActuelle < 0 {
			m.PvActuelle = 0
		}

		ClearScreen()

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

		DisplayCombatHeader(tourCombat, true)

		fmt.Println("\n┌─────────────────────────────┐  ┌─────────────────────────────┐")
		fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │ 👹 \033[1m%-10s\033[0m    │\n", c.Nom, c.Niveau, m.Nom)
		fmt.Printf("│    \033[32m%-20s\033[0m     │  │ \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
		fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │   PV: %-3d/%-3d (%.0f%%)        │\n",
			c.PvActuelle, c.PvMax, playerHealthPercent,
			m.PvActuelle, m.PvMax, monsterHealthPercent)
		fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║            📜 DÉROULEMENT 📜             ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		fmt.Println()
		fmt.Println("┌──────────────────────────────────────────┐")
		fmt.Printf("│ ☠️  \033[1m%s\033[0m utilise \033[35mPoison\033[0m               │\n", c.Nom)
		fmt.Printf("│ 💥 Inflige \033[33m%d\033[0m points de dégâts à \033[31m%s\033[0m │\n", dommage, m.Nom)
		fmt.Println("└──────────────────────────────────────────┘")

	default:
		ClearScreen()

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

		DisplayCombatHeader(tourCombat, true)

		fmt.Println("\n┌─────────────────────────────┐  ┌─────────────────────────────┐")
		fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │ 👹 \033[1m%-10s\033[0m              │\n", c.Nom, c.Niveau, m.Nom)
		fmt.Printf("│    \033[32m%-20s\033[0m     │  │ \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
		fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │   PV: %-3d/%-3d (%.0f%%)        │\n",
			c.PvActuelle, c.PvMax, playerHealthPercent,
			m.PvActuelle, m.PvMax, monsterHealthPercent)
		fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║            📜 DÉROULEMENT 📜             ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		fmt.Println()
		fmt.Println("┌──────────────────────────────────────────┐")
		fmt.Printf("│ ❓ \033[1m%s\033[0m utilise \033[33m%s\033[0m              │\n", c.Nom, itemName)
		fmt.Printf("│ 🤔 Mais rien ne se passe...               │\n")
		fmt.Println("└──────────────────────────────────────────┘")
		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		reader.ReadString('\n')

		ClearScreen()
		return false
	}

	c.Inventaire[itemName]--
	if c.Inventaire[itemName] <= 0 {
		delete(c.Inventaire, itemName)
	}

	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	reader.ReadString('\n')

	ClearScreen()

	return true
}

func DisplayCombatHeader(tourCombat int, isTraining bool) {
	if tourCombat <= 0 {
		tourCombat = 1
	}

	if isTraining {
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Printf("║      ⚔️  COMBAT D'ENTRAÎNEMENT - TOUR %d ⚔️  ║\n", tourCombat)
		fmt.Println("╚══════════════════════════════════════════╝")
	} else {
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Printf("║        ⚔️  QUÊTE PRINCIPALE - TOUR %d ⚔️    ║\n", tourCombat)
		fmt.Println("╚══════════════════════════════════════════╝")
	}
}

func UseSpell(c *Character, m *Monster, reader *bufio.Reader, tourCombat int, isTraining bool) bool {
	if len(c.Skill) == 0 {
		fmt.Println("\n🔮 Vous n'avez pas encore appris de sorts !")
		fmt.Print("Appuyez sur Entrée pour revenir au combat...")
		reader.ReadString('\n')
		ClearScreen()
		return false
	}

	ClearScreen()

	DisplayCombatHeader(tourCombat, isTraining)

	playerHealthPercent := float64(c.PvActuelle) / float64(c.PvMax) * 100
	monsterHealthPercent := float64(m.PvActuelle) / float64(m.PvMax) * 100
	playerManaPercent := float64(c.Mana) / float64(c.ManaMax) * 100

	playerHealthBar := ""
	monsterHealthBar := ""
	playerManaBar := ""

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

		if float64(i)*5 < playerManaPercent {
			playerManaBar += "█"
		} else {
			playerManaBar += "░"
		}
	}

	fmt.Println("┌─────────────────────────────┐  ┌─────────────────────────────┐")
	fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │ 👹 \033[1m%-10s\033[0m           │\n", c.Nom, c.Niveau, m.Nom)
	fmt.Printf("│    \033[32m%-20s\033[0m     │  │    \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
	fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │   PV: %-3d/%-3d (%.0f%%)         │\n",
		c.PvActuelle, c.PvMax, playerHealthPercent,
		m.PvActuelle, m.PvMax, monsterHealthPercent)
	fmt.Printf("│    \033[36m%-20s\033[0m     │  │                             │\n", playerManaBar)
	fmt.Printf("│    MP: %-3d/%-3d (%.0f%%)        │  │                             │\n",
		c.Mana, c.ManaMax, playerManaPercent)
	fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║            🔮 LIVRE DE SORTS 🔮          ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	fmt.Println("\nChoisissez un sort à lancer :")

	for i, spell := range c.Skill {
		fmt.Printf("[%d] %s\n", i+1, getSpellDisplay(spell))
	}

	fmt.Printf("[%d] 🔙 Retour au combat\n", len(c.Skill)+1)

	fmt.Print("\nVotre choix: ")
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)

	var choiceNum int
	fmt.Sscanf(choice, "%d", &choiceNum)

	if choiceNum < 1 || choiceNum > len(c.Skill)+1 {
		fmt.Println("\n❌ Choix invalide ! Retour au combat...")
		fmt.Print("Appuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
		ClearScreen()
		return false
	}

	if choiceNum == len(c.Skill)+1 {
		fmt.Println("\n🔙 Retour au combat...")
		ClearScreen()
		return false
	}

	spellName := c.Skill[choiceNum-1]
	manaAvant := c.Mana
	damage := castSpell(spellName, c, m)
	if manaAvant != c.Mana {
		fmt.Printf("🔵 Mana restant : %d / %d\n", c.Mana, c.ManaMax)
	}

	fmt.Println("\n┌──────────────────────────────────────────┐")
	fmt.Printf("│ 🔮 \033[1m%s\033[0m lance \033[33m%s\033[0m !  │\n", c.Nom, spellName)
	fmt.Printf("│ 💥 Inflige \033[31m%d\033[0m points de dégâts à \033[1m%s\033[0m  │\n", damage, m.Nom)
	fmt.Println("└──────────────────────────────────────────┘")

	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	reader.ReadString('\n')
	ClearScreen()

	return true
}

func getSpellDisplay(spellName string) string {
	switch spellName {
	case "Hadouken (Street Fighter)":
		return "🔥 Hadouken (Street Fighter) - Boule de feu puissante (15 dégâts, 8 MP)"
	case "Éclair de Zeus (God of War)":
		return "⚡ Éclair de Zeus (God of War) - Frappe électrique (20 dégâts, 12 MP)"
	case "Blizzard (Final Fantasy)":
		return "🧊 Blizzard (Final Fantasy) - Attaque de glace (18 dégâts, 10 MP)"
	case "FUS RO DAH (Skyrim)":
		return "💫 FUS RO DAH (Skyrim) - Cri puissant (25 dégâts, 15 MP)"
	case "Tornade (Zelda)":
		return "🌪️ Tornade (Zelda) - Attaque de vent (17 dégâts, 9 MP)"
	case "Coup de poing":
		return "👊 Coup de poing - Attaque basique (8 dégâts, 3 MP)"
	default:
		return "❓ " + spellName + " - Sort inconnu (10 dégâts, 5 MP)"
	}
}

func castSpell(spellName string, c *Character, m *Monster) int {
	var baseDamage int
	var manaCost int

	switch spellName {
	case "Hadouken (Street Fighter)":
		baseDamage = 15
		manaCost = 8
	case "Éclair de Zeus (God of War)":
		baseDamage = 20
		manaCost = 12
	case "Blizzard (Final Fantasy)":
		baseDamage = 18
		manaCost = 10
	case "FUS RO DAH (Skyrim)":
		baseDamage = 25
		manaCost = 15
	case "Tornade (Zelda)":
		baseDamage = 17
		manaCost = 9
	case "Coup de poing":
		baseDamage = 8
		manaCost = 3
	default:
		baseDamage = 10
		manaCost = 5
	}

	if c.Mana < manaCost {
		fmt.Printf("❌ Pas assez de mana pour lancer %s ! (Coût: %d, Actuel: %d)\n", spellName, manaCost, c.Mana)
		fmt.Print("Appuyez sur Entrée pour continuer...")
		return 0
	}

	c.Mana -= manaCost

	damage := ApplyStatBonuses(c, baseDamage, "spell")

	m.PvActuelle -= damage
	if m.PvActuelle < 0 {
		m.PvActuelle = 0
	}

	return damage
}

func DisplayHealthBarsQuest(c *Character, m *Monster) {
	playerHealthPercent := float64(c.PvActuelle) / float64(c.PvMax) * 100
	monsterHealthPercent := float64(m.PvActuelle) / float64(m.PvMax) * 100
	playerManaPercent := float64(c.Mana) / float64(c.ManaMax) * 100

	playerHealthBar := ""
	monsterHealthBar := ""
	playerManaBar := ""

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

		if float64(i)*5 < playerManaPercent {
			playerManaBar += "█"
		} else {
			playerManaBar += "░"
		}
	}

	fmt.Println("┌─────────────────────────────┐  ┌─────────────────────────────┐")
	fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │    👹 \033[1m%-20s\033[0m  │\n", c.Nom, c.Niveau, m.Nom)
	fmt.Printf("│    \033[32m%-20s\033[0m     │  │    \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
	fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │    PV: %-3d/%-3d (%.0f%%)        │\n",
		c.PvActuelle, c.PvMax, playerHealthPercent,
		m.PvActuelle, m.PvMax, monsterHealthPercent)
	fmt.Printf("│    \033[36m%-20s\033[0m     │  │                             │\n", playerManaBar)
	fmt.Printf("│    MP: %-3d/%-3d (%.0f%%)        │  │                             │\n",
		c.Mana, c.ManaMax, playerManaPercent)
	fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")
}
