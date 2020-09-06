package util

type MinIntHeap []int

func (a MinIntHeap) Len() int {
	return len(a)
}

func (a MinIntHeap) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a MinIntHeap) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a *MinIntHeap) Push(value interface{}) {
	*a = append(*a, value.(int))
}

func (a *MinIntHeap) Pop() interface{} {
	old := *a
	value := old[len(old)-1]
	*a = old[:len(old)-1]

	return value
}
