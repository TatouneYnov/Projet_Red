package red

import "fmt"

func CanAddItem(c *Character, quantity int) bool {
	total := 0
	for _, q := range c.Inventaire {
		total += q
	}
	return total+quantity <= c.InventorySize
}

func AddInventory(c *Character, itemName string, quantity int) {
	if c.Inventaire == nil {
		c.Inventaire = make(map[string]int)
	}
	if itemName == "Mana" {
		fmt.Println("❌ Le mana ne peut pas être ajouté à l'inventaire. Utilisez une potion ou un mécanisme dédié.")
		return
	}
	if !CanAddItem(c, quantity) {
		if itemName != "Mana" {
			fmt.Printf("❌ Inventaire plein ! Vous ne pouvez pas avoir plus de %d items.\n", c.InventorySize)
		}
		return
	}
	c.Inventaire[itemName] += quantity
	if quantity == 1 {
		fmt.Printf("✅ %s ajouté à votre inventaire !\n", itemName)
	} else {
		fmt.Printf("✅ %d x %s ajoutés à votre inventaire !\n", quantity, itemName)
	}
}

func RemoveInventory(c *Character, itemName string, quantity int) bool {
	if c.Inventaire == nil {
		return false
	}

	currentQuantity, exists := c.Inventaire[itemName]
	if !exists || currentQuantity < quantity {
		return false
	}

	c.Inventaire[itemName] -= quantity

	if c.Inventaire[itemName] <= 0 {
		delete(c.Inventaire, itemName)
	}

	if quantity == 1 {
		fmt.Printf("➖ %s retiré de votre inventaire.\n", itemName)
	} else {
		fmt.Printf("➖ %d x %s retirés de votre inventaire.\n", quantity, itemName)
	}

	return true
}

func GetItemQuantity(c *Character, itemName string) int {
	if c.Inventaire == nil {
		return 0
	}
	return c.Inventaire[itemName]
}
