package main

import "fmt"

func main() {
	items := [][]string{[]string{"phone", "blue", "pixel"}, []string{"computer", "silver", "lenovo"}, []string{"phone", "gold", "iphone"}}

	ruleKey := "color"
	ruleValue := "silver"
	total := countMatches(items, ruleKey, ruleValue)
	fmt.Println(total)

}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	total := 0
	for i, _ := range items {
		switch ruleKey {
		case "type":
			if items[i][0] == ruleValue {
				total++
			}
		case "color":
			if items[i][1] == ruleValue {
				total++
			}
		case "name":
			if items[i][2] == ruleValue {
				total++
			}

		}
	}

	return total
}
