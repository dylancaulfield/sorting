package sorting

func Merge(values *[]interface{}, l int, r int, greaterFn func(interface{}, interface{}) bool) {

	if r-l < 2 {
		return
	}

	m := (r - l) / 2

	Merge(values, l, m, greaterFn)
	Merge(values, m+1, r, greaterFn)

	mergeTogether(values, l, r, greaterFn)

}

func mergeTogether(values *[]interface{}, l int, r int, greaterFn func(interface{}, interface{}) bool) {

	m := (r + l) / 2

	sizeLeft := m - l + 1
	sizeRight := r - m

	left := make([]interface{}, sizeLeft)
	right := make([]interface{}, sizeRight)

	for i := 0; i < sizeLeft; i++ {

		left[i] = (*values)[i+l]

	}

	for i := 0; i < sizeRight; i++ {

		right[i] = (*values)[i+m+1]

	}

	i := 0
	j := 0
	k := l

	for i < sizeLeft && j < sizeRight {

		if greaterFn(left[i], right[j]) {

			(*values)[k] = right[j]

			j++
			k++
		} else {

			(*values)[k] = left[i]

			i++
			k++
		}

	}

	for i < sizeLeft {

		(*values)[k] = left[i]

		i++
		k++
	}

	for j < sizeRight {

		(*values)[k] = right[j]

		j++
		k++
	}

	return
}
