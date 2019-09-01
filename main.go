package main

import (
	"imagepredict/components"

	"github.com/gopherjs/vecty"
)

func main() {
	// prediction := classify.Predict("https://hillrag.com/wp-content/uploads/2017/11/rabbit-pic.jpg")
	// for _, l := range prediction {
	// 	fmt.Printf("label: %s, probability: %.2f%%\n", l.Label, l.Probability*100)
	// }

	vecty.SetTitle("Image Classifier | Vecty")
	p := &components.PageView{
		// Labels: prediction,
	}
	vecty.RenderBody(p)
}
