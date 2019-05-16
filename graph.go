package main

import "os"

type cycle struct {
	labels []string
	length int
}

var labels = [...]string{"1", "2", "3", "4"}

func newCycle(n int) cycle {
	c := make([]string, n)
	return cycle{labels: c, length: n}
}

// up to the labels it is the same if we name the first two edges with 1 and 2 respectively
func (c *cycle) labelFirstTwoEdges() {
	c.labels[0] = labels[0]
	c.labels[1] = labels[1]
	return
}

// fmt.Print(cycle) calls this method
func (c *cycle) String() string {
	ret := c.labels[0]
	for i := 1; i < c.length; i++ {
		ret = ret + ", " + c.labels[i]
	}
	return ret + "\n"
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

func (c *cycle) nextLabel(current int) {
	if current == c.length-1 {
		if c.labels[0] == c.labels[c.length-1] {
			return
		}
		if !c.everyLabelTwice() {
			return
		}

		// not quite an improvement but still
		for i := 2; i < 6; i++ {
			if !c.partition(i) {
				return
			}
		}
		c.writeToFile()
		return
	}
	for i := range labels {
		if labels[i] == c.labels[current] {
			continue
		} else {
			c.labels[current+1] = labels[i]
			c.nextLabel(current + 1)
		}

	}
}
