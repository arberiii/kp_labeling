package main

type cycle []int

func createCycle(n int) cycle {
	c := make([]int, n)
	return c
}

func (c cycle) labelFirstTwoEdges() {
	c[0] = 1
	c[1] = 2
	return
}