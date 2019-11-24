package sorting

func Bubble(sortables []Sortable) []Sortable {

	var temp Sortable
	swapped := 1

	for swapped > 0 {

		swapped = 0

		for i := 0; i < len(sortables)-1; i++ {

			if sortables[i].Greater(sortables[i], sortables[i+1]) {

				temp = sortables[i]
				sortables[i] = sortables[i+1]
				sortables[i+1] = temp

				swapped++
			}

		}

	}

	return sortables

}
