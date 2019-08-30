package classify

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

var (
	graphFile  = "../model/tensorflow_inception_graph.pb"
	labelsFile = "../model/imagenet_comp_graph_label_strings.txt"
)

// Label struct to get result.
type Label struct {
	Label       string
	Probability float32
}

//Labels jkfjnf
type Labels []Label

// Len knfndkfnv
func (l Labels) Len() int { return len(l) }

//Swap jknfjkdnd
func (l Labels) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

//Less jknfjkdnd
func (l Labels) Less(i, j int) bool { return l[i].Probability > l[j].Probability }

func normalizeImage(body io.ReadCloser) (*tf.Tensor, error) {
	var buf bytes.Buffer
	io.Copy(&buf, body)
	t, err := tf.NewTensor(buf.String())
	if err != nil {
		return nil, err
	}
	graph, input, output, err := getNormalizedGraph()
	if err != nil {
		return nil, err
	}

	session, err := tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	normalized, err := session.Run(map[tf.Output]*tf.Tensor{
		input: t,
	},
		[]tf.Output{
			output,
		}, nil)
	if err != nil {
		return nil, err
	}

	return normalized[0], nil
}

func getNormalizedGraph() (*tf.Graph, tf.Output, tf.Output, error) {
	s := op.NewScope()
	input := op.Placeholder(s, tf.String)
	decode := op.DecodeJpeg(s, input, op.DecodeJpegChannels(3))

	output := op.Sub(s,
		op.ResizeBilinear(s,
			op.ExpandDims(s,
				op.Cast(s, decode, tf.Float),
				op.Const(s.SubScope("make_batch"), int32(0))),
			op.Const(s.SubScope("size"), []int32{224, 224})),
		op.Const(s.SubScope("mean"), float32(117)))
	graph, err := s.Finalize()

	return graph, input, output, err
}

func loadGraphLabels() (*tf.Graph, []string, error) {
	model, err := ioutil.ReadFile(graphFile)
	if err != nil {
		return nil, nil, err
	}

	g := tf.NewGraph()
	if err = g.Import(model, ""); err != nil {

		return nil, nil, err
	}

	f, err := os.Open(labelsFile)
	if err != err {
		return nil, nil, err
	}
	defer f.Close()

	var labels []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}
	return g, labels, nil

}

// Predict does the actual prediction
func Predict(url string) []Label {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Unable o get an image: %v\n", err)
	}
	defer resp.Body.Close()

	modelGraph, labels, err := loadGraphLabels()
	if err != nil {
		log.Fatalf("Unable to load graph: %v\n", err)
	}
	session, err := tf.NewSession(modelGraph, nil)
	if err != nil {
		log.Fatalf("Unable to init session: %v\n", err)
	}
	defer session.Close()

	tensor, err := normalizeImage(resp.Body)
	if err != nil {
		log.Fatalf("Unable to normalize image: %v\n", err)
	}

	result, err := session.Run(map[tf.Output]*tf.Tensor{
		modelGraph.Operation("input").Output(0): tensor,
	},
		[]tf.Output{
			modelGraph.Operation("output").Output(0),
		}, nil)
	if err != nil {
		log.Fatalf("Unable to inference: %v\n", err)
	}

	return getTopFiveLabels(labels, result[0].Value().([][]float32)[0])

}

func getTopFiveLabels(labels []string, probabilities []float32) []Label {
	var results []Label
	for i, p := range probabilities {
		if i >= len(labels) {
			break
		}
		results = append(results, Label{
			Label:       labels[i],
			Probability: p})
	}
	sort.Sort(Labels(results))
	return results[:5]
}
