package red

import (
	"fmt"
	"math"
)

func CalculateExpForLevel(niveau int) int {
	return int(100 * math.Pow(float64(niveau), 1.5))
}

func CheckLevelUp(c *Character) {
	if c.ExpActuelle >= c.ExpMax {
		excessExp := c.ExpActuelle - c.ExpMax

		oldLevel := c.Niveau
		c.Niveau++

		c.ExpTotal += c.ExpMax

		c.ExpActuelle = excessExp

		c.ExpMax = CalculateExpForLevel(c.Niveau)

		previousMax := c.PvMax
		c.PvMax += 20

		c.PvActuelle = c.PvMax

		c.PointsStats += 3

		fmt.Println("\n🎉 \033[33mFÉLICITATIONS! Vous avez atteint le niveau\033[0m", c.Niveau, "!")
		fmt.Println("⬆️ Vos points de vie maximum ont augmenté :", previousMax, "→", c.PvMax)
		fmt.Println("💫 Vos PV ont été complètement restaurés !")
		fmt.Println("🔶 Vous avez gagné 3 points de statistiques à distribuer !")

		if c.Niveau == 3 {
			c.Skill = append(c.Skill, "Charge")
			fmt.Println("✨ Vous avez appris un nouveau sort : \033[35mCharge\033[0m !")
		} else if c.Niveau == 5 {
			c.Skill = append(c.Skill, "Coup Puissant")
			fmt.Println("✨ Vous avez appris un nouveau sort : \033[35mCoup Puissant\033[0m !")
		} else if c.Niveau == 7 {
			c.Skill = append(c.Skill, "Onde de Choc")
			fmt.Println("✨ Vous avez appris un nouveau sort : \033[35mOnde de Choc\033[0m !")
		} else if c.Niveau == 10 {
			c.Skill = append(c.Skill, "Tempête d'Acier")
			fmt.Println("✨ Vous avez appris un nouveau sort : \033[35mTempête d'Acier\033[0m !")
		}

		if c.Niveau > oldLevel+1 {
			fmt.Printf("🌟 \033[33mIncroyable ! Vous avez gagné %d niveaux d'un coup !\033[0m\n", c.Niveau-oldLevel)
		}

		if c.ExpActuelle >= c.ExpMax {
			CheckLevelUp(c)
		}
	}
}

func DisplayExpBar(c *Character) string {
	expPercent := float64(c.ExpActuelle) / float64(c.ExpMax) * 100

	expBar := ""
	for i := 0; i < 20; i++ {
		if float64(i)*5 < expPercent {
			expBar += "█"
		} else {
			expBar += "░"
		}
	}

	return fmt.Sprintf("\033[36m%s\033[0m %d/%d (%.1f%%)", expBar, c.ExpActuelle, c.ExpMax, expPercent)
}
