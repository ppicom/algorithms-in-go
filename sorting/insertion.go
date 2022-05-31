package sorting

func InsertionSort(keys []int) []int {
	for j := 1; j < len(keys); j++ {
		key := keys[j]
		i := j - 1

		for ;i > 0 && keys[i] < key; i--{
			keys[i+1] = keys[i]
		}

		keys[i+1] = key
	}

	return keys
}
