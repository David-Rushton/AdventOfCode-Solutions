package main

func getLoadouts() []loadout {
	result := []loadout{}

	for _, weapon := range getWeapons() {
		for _, armour := range getArmours() {
			for _, leftRing := range getRings() {
				for _, rightRing := range getRings() {
					if leftRing != rightRing {
						result = append(result, loadout{weapon, armour, leftRing, rightRing})
					}
				}
			}
		}
	}

	return result
}

func getWeapons() []weapon {
	return []weapon{
		{name: "Dagger", cost: 8, damage: 4, armour: 0},
		{name: "Shortsword", cost: 10, damage: 5, armour: 0},
		{name: "Warhammer", cost: 25, damage: 6, armour: 0},
		{name: "Longsword", cost: 40, damage: 7, armour: 0},
		{name: "Greataxe", cost: 74, damage: 8, armour: 0},
	}
}

func getArmours() []armour {
	return []armour{
		{name: "Leather", cost: 13, damage: 0, armour: 1},
		{name: "Chainmail", cost: 31, damage: 0, armour: 2},
		{name: "Splintmail", cost: 53, damage: 0, armour: 3},
		{name: "Bandedmail", cost: 75, damage: 0, armour: 4},
		{name: "Platemail", cost: 102, damage: 0, armour: 5},
		{name: "None", cost: 0, damage: 0, armour: 0},
	}
}

func getRings() []ring {
	return []ring{
		{name: "Damage +1", cost: 25, damage: 1, armour: 0},
		{name: "Damage +2", cost: 50, damage: 2, armour: 0},
		{name: "Damage +3", cost: 100, damage: 3, armour: 0},
		{name: "Defense +1", cost: 20, damage: 0, armour: 1},
		{name: "Defense +2", cost: 40, damage: 0, armour: 2},
		{name: "Defense +3", cost: 80, damage: 0, armour: 3},
		{name: "No Ring 1", cost: 0, damage: 0, armour: 0},
		{name: "No Ring 2", cost: 0, damage: 0, armour: 0},
	}
}
