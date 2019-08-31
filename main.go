package main

import (
	"fmt"

	"imagepredict/classify"
)

func main() {
	prediction := classify.Predict("https://hillrag.com/wp-content/uploads/2017/11/rabbit-pic.jpg")
	for _, l := range prediction {
		fmt.Printf("label: %s, probability: %.2f%%\n", l.Label, l.Probability*100)
	}
}
