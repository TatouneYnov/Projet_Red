package red

import (
	"fmt"
	"time"
)

func IsDead(c *Character) bool {
	if c.PvActuelle <= 0 {
		fmt.Println("\n💀 ═══════════════════════════════════════ 💀")
		fmt.Println("     Vous êtes tombé au combat...")
		fmt.Println("     Vos forces vous abandonnent...")
		fmt.Println("💀 ═══════════════════════════════════════ 💀")

		fmt.Print("\n💫 ")
		for i := 0; i < 3; i++ {
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println()

		c.PvActuelle = c.PvMax / 2

		fmt.Println("\n✨ ═══════════════════════════════════════ ✨")
		fmt.Println("     Une force mystérieuse vous ranime...")
		fmt.Printf("     Vous vous réveillez avec %d/%d PV !\n", c.PvActuelle, c.PvMax)
		fmt.Println("✨ ═══════════════════════════════════════ ✨")

		fmt.Println("\n🔮 \"Les dieux vous accordent une seconde chance...\"")

		return true
	}

	return false
}
