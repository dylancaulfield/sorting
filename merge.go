package sorting

func Merge(values *[]interface{}, greaterFn func(interface{}, interface{}) bool) {

	if len(*values) < 2 {
		return
	}

	if len(*values) == 2 {

		if greaterFn((*values)[0], (*values)[1]) {

			(*values)[0], (*values)[1] = (*values)[1], (*values)[0]
			return

		} else {

			return

		}

	}

	firstHalf := (*values)[:len(*values)/2]
	secondHalf := (*values)[len(*values)/2:]

	Merge(&firstHalf, greaterFn)
	Merge(&secondHalf, greaterFn)

	merged := mergeTogether(&firstHalf, &secondHalf, greaterFn)

	*values = merged

}

func mergeTogether(firstHalf, secondHalf *[]interface{}, greaterFn func(interface{}, interface{}) bool) []interface{} {

	var i, j = 0, 0

	merged := make([]interface{}, len(*firstHalf)+len(*secondHalf))

	for i+j < len(*firstHalf)+len(*secondHalf) {

		if i == len(*firstHalf) {

			for j = j; j < len(*secondHalf); j++ {
				merged[i+j] = (*secondHalf)[j]
			}

			break

		} else if j == len(*secondHalf) {

			for i = i; i < len(*firstHalf); i++ {
				merged[i+j] = (*firstHalf)[i]
			}

			break
		}

		if greaterFn((*firstHalf)[i], (*secondHalf)[j]) {

			merged[i+j] = (*secondHalf)[j]

			j++

		} else {

			merged[i+j] = (*firstHalf)[i]

			i++

		}

	}

	return merged

}
