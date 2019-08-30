package main

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/example/todomvc/components"
	"github.com/mss7082/ImgClassifier/classify"
)

func main() {
	prediction := classify.Predict("http://yesofcorsa.com/wp-content/uploads/2015/10/1372_rabbit.jpg")
	for _, l := range prediction {
		fmt.Printf("label: %s, probability: %.2f%%\n", l.Label, l.Probability*100)
	}
	vecty.SetTitle("Image Classifier")
	p := &components.PageView{}
}
