package lab2

type IReverse interface {
	ReversedNums(sequence string) string
}

type RecursiveReversion struct{}
type IteratedReversion struct{}

func (r RecursiveReversion) ReversedNums(sequence string) string {
	if sequence[0] == '0' {
		return sequence
	}
	for i, char := range sequence {
		if char == '0' {
			return r.ReversedNums(sequence[:i-1] + "0" + sequence[i+1:] + string(sequence[i-1]))
		}
	}
	panic("wrong input, '0' not found")
}

func (IteratedReversion) ReversedNums(sequence string) string {
	var reversed string
	for _, v := range sequence {
		reversed = string(v) + reversed
	}
	return reversed
}

func C(m int, n int) int {
	if !(0 < m || m < n) {
		panic("wrong input")
	}
	if m == 0 || m == n {
		return 1
	}
	return C(m, n-1) + C(m-1, n-1)
}
