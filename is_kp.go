package main

import (
	"sort"
)

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
	// never do this!!!
	//sort.Strings(labels)
	labelsSet := make(map[string]bool)
	for _, label := range labels {
		labelsSet[label] = true
	}
	var ret []string
	for label := range labelsSet {
		ret = append(ret, label)
	}
	sort.Strings(ret)
	return ret
}

func (c *cycle) partition(half int) bool {
	firstHalf := labelsOrderSet(c.labels[:half])
	secondHalf := labelsOrderSet(c.labels[half:c.length])
	lengthOfFirst := len(firstHalf)
	lengthOfSecond := len(secondHalf)
	count := 0
	indexFirst := 0
	indexSecond := 0
	for {
		if count > 1 {
			return true
		} else if indexFirst > lengthOfFirst-1 {
			return false
		} else if indexSecond > lengthOfSecond-1 {
			return false
		}
		if firstHalf[indexFirst] == secondHalf[indexSecond] {
			count = count + 1
			indexFirst = indexFirst + 1
			indexSecond = indexSecond + 1
		} else if firstHalf[indexFirst] < secondHalf[indexSecond] {
			indexFirst = indexFirst + 1
		} else if firstHalf[indexFirst] > secondHalf[indexSecond] {
			indexSecond = indexSecond + 1
		}

	}
}
