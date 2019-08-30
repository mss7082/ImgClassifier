package main

import (
	"fmt"
	"github.com/mss7082/ImgClassifier/classifier/"
)

func main() {
	prediction := classifier.predict("http://yesofcorsa.com/wp-content/uploads/2015/10/1372_rabbit.jpg")
	for _, l := range prediction {
		fmt.Printf("label: %s, probability: %.2f%%\n", l.Label, l.Probability*100)
	}
}
