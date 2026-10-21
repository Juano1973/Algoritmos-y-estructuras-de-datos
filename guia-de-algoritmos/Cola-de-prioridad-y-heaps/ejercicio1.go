package main

func esHeap(arr []int) bool {

	largo := len(arr)

	if largo < 2 {
		return true
	}

	for i := 0; i < largo; i++ {
		hijo_izq := (2 * i) + 1
		hijo_der := (2 * i) + 2

		if hijo_izq < largo {
			if arr[hijo_izq] < arr[i] {
				return false
			}
		}

		if hijo_der < largo {
			if arr[hijo_der] < arr[i] {
				return false
			}
		}
	}

	return true
}
