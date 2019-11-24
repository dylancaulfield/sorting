package sorting

func Bubble(values *[]interface{}, greaterFn func(interface{}, interface{}) bool) {

	var temp interface{}
	swapped := 1

	for swapped > 0 {

		swapped = 0

		for i := 0; i < len(*values)-1; i++ {

			if greaterFn((*values)[i], (*values)[i+1]) {

				temp = (*values)[i]
				(*values)[i] = (*values)[i+1]
				(*values)[i+1] = temp

				swapped++
			}

		}

	}

	return

}
