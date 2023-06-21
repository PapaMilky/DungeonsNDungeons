package utils

type Inventory struct {
	Size       int      `json:"inventorySize"`
	Helm       Item     `json:"helm"`
	Chestplate Item     `json:"chestplate"`
	Leggings   Item     `json:"leggings"`
	Boots      Item     `json:"boots"`
	Rucksack   Rucksack `json:"rucksack"`
	Hand       Item     `json:"hand"`
	Slots      []Item   `json:"inventory"`
}

type Rucksack struct {
	Size   int    `json:"size"`
	Rarity int    `json:"rarity"`
	Slots  []Item `json:"slots"`
	Gizmos []Item `json:"gizmos"`
}

type Item struct {
	Name          string      `json:"name"`
	Lore          string      `json:"lore"`
	Type          int         `json:"type"`
	ItemID        int         `json:"itemID"`
	DurabilityMax int         `json:"durabilityMax"`
	Durability    int         `json:"durability"`
	StackSize     int         `json:"stackSize"`
	Quantity      int         `json:"quantity"`
	Skills        []ItemSkill `json:"skills"`
	Modifiers     []Modifier  `json:"modifiers"`
	Stats         Stats       `json:"stats"`
}

type Modifier struct {
	Type  int `json:"type"`
	Level int `json:"level"`
}

type ItemSkill struct {
	Type  int `json:"type"`
	Level int `json:"level"`
}
