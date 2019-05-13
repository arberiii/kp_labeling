package main

import "os"

type cycle struct {
	labels []string
	length int
}

func createCycle(n int) cycle {
	c := make([]string, n)
	return cycle{labels: c, length: n}
}

func (c *cycle) labelFirstTwoEdges() {
	c.labels[0] = ""
	c.labels[1] = ""
	return
}

func (c *cycle) String() string {
	ret := c.labels[0]
	for i := 1; i < c.length; i++ {
		ret = ret + ", " + c.labels[i]
	}
	return ret
}

func (c *cycle) writeToFile() {
	f, err := os.OpenFile("cycle", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(c.String()); err != nil {
		panic(err)
	}
}
