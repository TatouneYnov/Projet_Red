package red

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func MainQuestFight(c *Character) {
	if c.MainQuest.Completed {
		fightFinalBoss(c)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	c.Mana = c.ManaMax

	monsterIndex := c.MainQuest.Progression + 1
	monster := InitQuestMonster(monsterIndex)

	ClearScreen()

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║           ⚔️ QUÊTE PRINCIPALE ⚔️          ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	fmt.Printf("\n👹 Un %s apparaît devant vous !\n", monster.Nom)
	fmt.Println("⚔️ Préparez-vous au combat !")

	tourCombat := 1
	reader.ReadString('\n')

	for {
		ClearScreen()

		DisplayCombatHeader(tourCombat, false)

		DisplayHealthBars(c, &monster)

		result := CharacterTurn(c, &monster, reader, tourCombat, false)

		if result == -1 {
			continue
		} else if result == 0 {
			continue
		} else if result == 2 {
			return
		}

		if monster.PvActuelle <= 0 {
			break
		}

		fmt.Println("\n👹 Tour du monstre...")
		time.Sleep(1 * time.Second)

		QuestMonsterPattern(tourCombat, monster, c, monsterIndex)

		if c.PvActuelle <= 0 {
			break
		}

		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		reader.ReadString('\n')

		tourCombat++
	}

	if c.PvActuelle <= 0 {
		fmt.Println("\n☠️ Défaite !")
		fmt.Printf("Vous avez été vaincu par le %s...\n", monster.Nom)
		fmt.Println("Vous vous réveillez plus tard, affaibli mais vivant...")
		c.PvActuelle = c.PvMax / 4
	} else if monster.PvActuelle <= 0 {
		fmt.Println("\n🏆 Victoire !")
		fmt.Printf("Vous avez vaincu le %s !\n", monster.Nom)

		c.MainQuest.Progression++

		if c.MainQuest.Progression >= 10 {
			c.MainQuest.Completed = true
			fmt.Println("\n🎖️ FÉLICITATIONS ! Vous avez terminé la quête principale !")
			fmt.Println("   Vous avez vaincu les 10 monstres légendaires et êtes désormais un héros reconnu !")
			bonusOr := 1000
			bonusExp := 500
			c.Argent += bonusOr
			c.ExpActuelle += bonusExp
			fmt.Printf("\n�� Récompense spéciale : %d pièces d'or !\n", bonusOr)
			fmt.Printf("✨ Bonus d'expérience : %d points !\n", bonusExp)

			fmt.Println("\n⚠️ Mais attendez... Une présence mystérieuse se fait sentir...")
			fmt.Println("🖥️ Un adversaire ultime est apparu : Print Alphabet, le boss final des exercices d'initiation Go !")
			fmt.Println("👨‍💻 Prouvez votre maîtrise du code en l'affrontant lorsque vous serez prêt !")
		} else {
			expGain := monster.Exp
			goldGain := 30 + 10*monsterIndex

			c.ExpActuelle += expGain
			c.Argent += goldGain

			fmt.Printf("\n✨ Vous gagnez %d points d'expérience !\n", expGain)
			fmt.Printf("💰 Vous trouvez %d pièces d'or !\n", goldGain)

			CheckLevelUp(c)
		}
	}

	if c.PvActuelle > c.PvMax {
		c.PvActuelle = c.PvMax
	}
}

func fightFinalBoss(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	c.Mana = c.ManaMax

	monster := InitFinalBoss()

	ClearScreen()

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║           🔥 BOSS FINAL 🔥               ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	fmt.Println("\n⚠️ \033[31mATTENTION\033[0m ⚠️")
	fmt.Println("👨‍💻 Vous faites face au redoutable \033[1mPrint Alphabet\033[0m, le boss final des exercices d'initiation Go !")
	fmt.Println("💻 \033[36mfunc main() {\033[0m")
	fmt.Println("💻 \033[36m    for i := 'A'; i <= 'Z'; i++ {\033[0m")
	fmt.Println("💻 \033[36m        fmt.Print(string(i))\033[0m")
	fmt.Println("💻 \033[36m}\033[0m")

	tourCombat := 1
	reader.ReadString('\n')

	for {
		ClearScreen()

		DisplayCombatHeader(tourCombat, false)

		DisplayHealthBars(c, &monster)

		result := CharacterTurn(c, &monster, reader, tourCombat, false)

		if result == -1 {
			continue
		} else if result == 0 {
			continue
		}

		if monster.PvActuelle <= 0 {
			break
		}

		fmt.Println("\n👹 Tour du boss final...")
		time.Sleep(1 * time.Second)

		QuestMonsterPattern(tourCombat, monster, c, 11)

		if c.PvActuelle <= 0 {
			break
		}

		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		reader.ReadString('\n')

		tourCombat++
	}

	if c.PvActuelle <= 0 {
		fmt.Println("\n☠️ Défaite !")
		fmt.Printf("Vous avez été vaincu par %s...\n", monster.Nom)
		fmt.Println("\n💻 \033[36mfatal error: runtime: out of memory\033[0m")
		fmt.Println("Vous vous réveillez plus tard, affaibli mais vivant...")
		c.PvActuelle = c.PvMax / 4
	} else if monster.PvActuelle <= 0 {
		fmt.Println("\n🏆 \033[33mVICTOIRE ULTIME !\033[0m 🏆")
		fmt.Printf("Vous avez vaincu \033[1m%s\033[0m, le boss final !\n", monster.Nom)

		if !c.MainQuest.FinalBossDefeated {
			c.MainQuest.FinalBossDefeated = true

			bonusOr := 5000
			bonusExp := 2000

			c.Argent += bonusOr
			c.ExpActuelle += bonusExp

			fmt.Println("\n🎖️ \033[33mFÉLICITATIONS SUPRÊMES !\033[0m 🎖️")
			fmt.Println("🧠 Vous avez prouvé votre maîtrise absolue du code !")
			fmt.Println("👑 Vous êtes désormais une légende vivante, respectée par tous les développeurs !")
			fmt.Printf("\n💰 Trésor légendaire : \033[33m%d pièces d'or\033[0m !\n", bonusOr)
			fmt.Printf("✨ Illumination mentale : \033[33m%d points d'expérience\033[0m !\n", bonusExp)

			fmt.Println("\n💻 \033[36mfmt.Println(\"Félicitations, maître du code !\")\033[0m")
			fmt.Println("💻 \033[36m}\033[0m")
		} else {
			bonusOr := 1000
			bonusExp := 500

			c.Argent += bonusOr
			c.ExpActuelle += bonusExp

			fmt.Println("\n🎖️ \033[33mVICTOIRE RENOUVELÉE !\033[0m 🎖️")
			fmt.Println("👑 Vous avez encore une fois prouvé votre valeur face au boss final !")
			fmt.Printf("\n💰 Récompense : \033[33m%d pièces d'or\033[0m !\n", bonusOr)
			fmt.Printf("✨ Expérience : \033[33m%d points\033[0m !\n", bonusExp)
			fmt.Println("\n💻 \033[36mBoss final vaincu une nouvelle fois !\033[0m")
		}

		CheckLevelUp(c)
	}

	if c.PvActuelle > c.PvMax {
		c.PvActuelle = c.PvMax
	}
}

func DisplayQuestStatus(c Character) {
	ClearScreen()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║         📜 QUÊTE PRINCIPALE 📜            ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	if c.MainQuest.Completed && c.MainQuest.FinalBossDefeated {
		fmt.Println("\n🏆 QUÊTE TERMINÉE ! Vous avez vaincu tous les monstres et le boss final !")
		fmt.Println("👑 Vous êtes désormais une légende vivante, respectée par tous !")

		finalBoss := InitFinalBoss()
		fmt.Printf("\n🔥 Boss Final: %s (PV: %d) (déjà vaincu)\n", finalBoss.Nom, finalBoss.PvMax)
		fmt.Println("🔄 Vous pouvez toujours l'affronter à nouveau pour plus de récompenses !")
		return
	}

	if c.MainQuest.Completed {
		fmt.Println("\n🏆 Quête principale terminée ! Tous les monstres ont été vaincus !")
		finalBoss := InitFinalBoss()
		fmt.Printf("⚠️ Mais un boss final est apparu : %s (PV: %d) !\n", finalBoss.Nom, finalBoss.PvMax)
		fmt.Println("👨‍💻 Affrontez-le pour prouver votre maîtrise totale du code !")
		return
	}

	fmt.Printf("\n📊 Progression: %d/10 monstres vaincus\n", c.MainQuest.Progression)
	nextMonster := InitQuestMonster(c.MainQuest.Progression + 1)
	fmt.Printf("\n🎯 Prochain adversaire: %s (PV: %d)\n", nextMonster.Nom, nextMonster.PvMax)

	var finalBoss Monster = InitFinalBoss()
	fmt.Printf("\n🔥 Boss Final : %s (PV: %d)\n", finalBoss.Nom, finalBoss.PvMax)

	fmt.Println("\n📝 Liste des monstres à vaincre:")
	fmt.Println()

	for i := 1; i <= 10; i++ {
		monster := InitQuestMonster(i)
		if i <= c.MainQuest.Progression {
			fmt.Printf("✅ Monstre %d: %s (PV: %d)\n", i, monster.Nom, monster.PvMax)
		} else if i == c.MainQuest.Progression+1 {
			fmt.Printf("⚔️ Monstre %d: %s (PV: %d) (prochain)\n", i, monster.Nom, monster.PvMax)
		} else {
			fmt.Printf("❌ Monstre %d: %s (PV: %d)\n", i, monster.Nom, monster.PvMax)
		}
	}

	var bossFinal Monster = InitFinalBoss()
	if c.MainQuest.Completed && c.MainQuest.FinalBossDefeated {
		fmt.Printf("\n✅ Boss Final: %s (PV: %d) (déjà vaincu)\n", bossFinal.Nom, bossFinal.PvMax)
	} else if c.MainQuest.Completed {
		fmt.Printf("\n⚔️ Boss Final: %s (PV: %d) (disponible)\n", bossFinal.Nom, bossFinal.PvMax)
	} else {
		fmt.Printf("\n❌ Boss Final: %s (PV: %d) (après tous les monstres)\n", bossFinal.Nom, bossFinal.PvMax)
	}
}

func DisplayHealthBars(c *Character, m *Monster) {
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
