package main

import "fmt"

func main() {
	c := newCycle(8)
	c.labelFirstTwoEdges()
	c.nextLabel(1)
	fmt.Println(c)
	fmt.Println(labelsOrderSet(c.labels))
	fmt.Println(c)
	fmt.Println(c.partition(2))
}
