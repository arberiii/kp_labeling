package main

import "sort"

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

// given a path it will sort and find labels
func labelsOrderSet(labels []string) []string {
	sort.Strings(labels)
	var ret []string
	ret = append(ret, labels[0])
	for _, label := range labels {
		if ret[len(ret)-1] != label {
			ret = append(ret, label)
		}
	}
	return ret
}
