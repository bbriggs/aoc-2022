package util

type Elf struct {
	Calories int
}

type Elves []Elf

// SortElvesByCalories sorts the elves by calories
func SortElvesByCalories(elves Elves) {
	// Use literally any sorting algorithm you want because this is just AoC

	// I'm using a bubble sort because it's easy to implement and I'm lazy
	for i := 0; i < len(elves); i++ {
		for j := 0; j < len(elves)-1; j++ {
			if elves[j].Calories > elves[j+1].Calories {
				elves[j], elves[j+1] = elves[j+1], elves[j]
			}
		}
	}
}
