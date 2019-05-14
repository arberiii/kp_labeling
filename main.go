package main

func main() {
	c := newCycle(8)
	c.labelFirstTwoEdges()
	c.nextLabel(1)
}
