package main

func (c *cycle) everyLabelTwice() bool {
	for i := range labels {
		counter := 0
		for j := range c.labels {
			if c.labels[j] == labels[i] {
				counter = counter + 1
			}
		}
		if counter < 2 {
			return false
		}
	}
	return true
}

func (c *cycle) partitionSet(index int) bool {
	var min int
	if min = index; min > c.length-index {
		min = c.length - index
	}
	return false
}
