package main

func main() {
	c := newCycle(4)
	c.labelFirstTwoEdges()
	c.nextLabel(1)
}
