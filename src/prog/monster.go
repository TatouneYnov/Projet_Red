package red

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	Nom        string
	PvMax      int
	PvActuelle int
	Attaque    int
	Exp        int
	Initiative int
	Argent     int
}

func InitGoblin() Monster {
	gobelin := Monster{
		Nom:        "Gobelin",
		PvMax:      35,
		PvActuelle: 35,
		Attaque:    4,
		Exp:        15,
		Initiative: 6,
		Argent:     8,
	}

	return gobelin
}

func InitQuestMonster(monsterIndex int) Monster {
	switch monsterIndex {
	case 1:
		return Monster{
			Nom:        "Goomba (Mario)",
			PvMax:      40,
			PvActuelle: 40,
			Attaque:    5,
			Exp:        20,
			Initiative: 4,
			Argent:     10,
		}
	case 2:
		return Monster{
			Nom:        "Slime (Minecraft)",
			PvMax:      55,
			PvActuelle: 55,
			Attaque:    6,
			Exp:        30,
			Initiative: 5,
			Argent:     15,
		}
	case 3:
		return Monster{
			Nom:        "Creeper (Minecraft)",
			PvMax:      70,
			PvActuelle: 70,
			Attaque:    8,
			Exp:        40,
			Initiative: 7,
			Argent:     20,
		}
	case 4:
		return Monster{
			Nom:        "Teemo (LoL)",
			PvMax:      80,
			PvActuelle: 80,
			Attaque:    9,
			Exp:        50,
			Initiative: 8,
			Argent:     25,
		}
	case 5:
		return Monster{
			Nom:        "Widowmaker (Overwatch)",
			PvMax:      100,
			PvActuelle: 100,
			Attaque:    12,
			Exp:        70,
			Initiative: 9,
			Argent:     35,
		}
	case 6:
		return Monster{
			Nom:        "Bowser (Mario)",
			PvMax:      130,
			PvActuelle: 130,
			Attaque:    15,
			Exp:        90,
			Initiative: 6,
			Argent:     45,
		}
	case 7:
		return Monster{
			Nom:        "Terminator T-800",
			PvMax:      150,
			PvActuelle: 150,
			Attaque:    16,
			Exp:        110,
			Initiative: 7,
			Argent:     55,
		}
	case 8:
		return Monster{
			Nom:        "Lynel (Zelda BOTW)",
			PvMax:      180,
			PvActuelle: 180,
			Attaque:    18,
			Exp:        140,
			Initiative: 8,
			Argent:     70,
		}
	case 9:
		return Monster{
			Nom:        "Thanos (Marvel)",
			PvMax:      220,
			PvActuelle: 220,
			Attaque:    22,
			Exp:        180,
			Initiative: 10,
			Argent:     90,
		}
	case 10:
		return Monster{
			Nom:        "Bowser Final Boss",
			PvMax:      280,
			PvActuelle: 280,
			Attaque:    25,
			Exp:        250,
			Initiative: 8,
			Argent:     120,
		}
	default:
		return Monster{
			Nom:        "Monstre Inconnu",
			PvMax:      50,
			PvActuelle: 50,
			Attaque:    10,
			Exp:        30,
			Initiative: 6,
			Argent:     15,
		}
	}
}

func InitFinalBoss() Monster {
	return Monster{
		Nom:        "Print Alphabet",
		PvMax:      400,
		PvActuelle: 400,
		Attaque:    35,
		Exp:        400,
		Initiative: 12,
		Argent:     200,
	}
}

func InitTrainingDummy() Monster {
	return Monster{
		Nom:        "Sandbag Smash Bros",
		PvMax:      9999,
		PvActuelle: 9999,
		Attaque:    0,
		Exp:        1,
		Initiative: 0,
	}
}

func QuestMonsterPattern(tour int, monster Monster, c *Character, monsterIndex int) {
	damage := monster.Attaque

	fmt.Printf("\n⚔️ %s attaque !\n", monster.Nom)

	switch monsterIndex {
	case 5, 10, 11:
		if rand.Intn(100) < 20 {
			damage = int(float64(damage) * 1.5)
			fmt.Printf("🔥 %s utilise une attaque spéciale !\n", monster.Nom)
		}
	}

	c.PvActuelle -= damage

	if c.PvActuelle < 0 {
		c.PvActuelle = 0
	}

	fmt.Printf("💥 Vous subissez %d points de dégâts !\n", damage)
}

func TrainingDummyPattern(tour int, dummy Monster, c *Character) {
	fmt.Printf("\n🎯 %s se tient immobile devant vous.\n", dummy.Nom)
	fmt.Println("   C'est un parfait outil d'entraînement qui ne riposte pas.")
}

func GoblinPattern(tour int, gobelin Monster, c *Character) {
	var degats int
	var isSpecialAttack bool

	if tour%3 == 0 {
		degats = gobelin.Attaque * 2
		isSpecialAttack = true
	} else {
		degats = gobelin.Attaque
		isSpecialAttack = false
	}

	c.PvActuelle -= degats
	if c.PvActuelle < 0 {
		c.PvActuelle = 0
	}

	fmt.Println()
	fmt.Println("┌──────────────────────────────────────────┐")
	if isSpecialAttack {
		fmt.Printf("│ � \033[1m%s\033[0m utilise \033[31mAttaque puissante\033[0m !   │\n", gobelin.Nom)
		fmt.Printf("│ 💀 Inflige \033[31m%d\033[0m points de dégâts à \033[1m%s\033[0m  │\n", degats, c.Nom)
	} else {
		fmt.Printf("│ 👹 \033[1m%s\033[0m attaque \033[1m%s\033[0m                     │\n", gobelin.Nom, c.Nom)
		fmt.Printf("│ 🗡️  Inflige \033[33m%d\033[0m points de dégâts            │\n", degats)
	}
	fmt.Println("└──────────────────────────────────────────┘")
}
