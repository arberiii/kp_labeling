package main

func main() {
	c := newCycle(6)
	c.labelFirstTwoEdges()
	c.nextLabel(1)
}
