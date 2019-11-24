package sorting

func Bubble(values *[]interface{}, greaterFn func(interface{}, interface{}) bool) {

	swapped := 1

	for swapped > 0 {

		swapped = 0

		for i := 0; i < len(*values)-1; i++ {

			if greaterFn((*values)[i], (*values)[i+1]) {

				(*values)[i], (*values)[i+1] = (*values)[i+1], (*values)[i]

				swapped++
			}

		}

	}

	return

}
