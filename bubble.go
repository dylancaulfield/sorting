package sorting

func Bubble(sortables []interface{}, greaterFn func(interface{}, interface{}) bool) []interface{} {

	var temp interface{}
	swapped := 1

	for swapped > 0 {

		swapped = 0

		for i := 0; i < len(sortables)-1; i++ {

			if greaterFn(sortables[i], sortables[i+1]) {

				temp = sortables[i]
				sortables[i] = sortables[i+1]
				sortables[i+1] = temp

				swapped++
			}

		}

	}

	return sortables

}
