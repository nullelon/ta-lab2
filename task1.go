package lab2

func ReverseDigitsSequence(sequence string) string {
	if sequence[0] == '0' {
		return sequence
	}
	for i, c := range sequence {
		if c == '0' {
			return ReverseDigitsSequence(sequence[1:i] + "0" + string(sequence[0]) + sequence[i+1:])
		}
	}
	panic("wrong input")
}

func TaskTwoFunction(m, n int) int {
	if m == 0 || m == n {
		return 1
	}
	if 0 < m && m < n {
		return TaskTwoFunction(m, n-1) + TaskTwoFunction(m-1, n-1)
	}
	panic("wrong input")
}
