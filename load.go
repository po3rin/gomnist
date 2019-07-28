// Package gomnist lets you to load the MNIST data set for use with gonum package.
package gomnist

import (
	"image/color"
	"sync"

	"github.com/petar/GoMNIST"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/mat"
)

// LoadData gets Mnist data using petar/GoMNIST package.
func load(rootPath string) (train *GoMNIST.Set, test *GoMNIST.Set, err error) {
	trainSet, testSet, err := GoMNIST.Load(rootPath)
	if err != nil {
		return nil, nil, err
	}
	return trainSet, testSet, nil
}

func set2Mat(s *GoMNIST.Set) (data mat.Matrix, labels mat.Matrix) {
	d := mat.NewDense(len(s.Images), s.NRow*s.NRow, nil)
	l := mat.NewDense(len(s.Labels), 1, nil)

	var wg sync.WaitGroup
	for i := 0; i < len(s.Images); i++ {
		wg.Add(1)
		go func(i int) {
			image, label := s.Get(i)
			b := image.Bounds()

			imageVec := make([]float64, 0, s.NRow*s.NRow)
			for n := 0; n < b.Max.Y; n++ {
				for m := 0; m < b.Max.X; m++ {
					v := float64(image.At(n, m).(color.Gray).Y)
					imageVec = append(imageVec, v)
				}
			}

			d.SetRow(i, imageVec)
			l.SetRow(i, []float64{float64(label)})

			wg.Done()
		}(i)
	}

	wg.Wait()
	return d, l
}

// LoadMat loads MNIST data for gonum matrix.
func LoadMat(rootPath string) (
	trainData mat.Matrix,
	trainLabels mat.Matrix,
	testData mat.Matrix,
	testLabels mat.Matrix,
	err error,
) {
	trainSet, testSet, err := load(rootPath)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "gomnist: failed to load mnist data")
	}

	trainData, trainLabels = set2Mat(trainSet)
	testData, testLabels = set2Mat(testSet)
	return trainData, trainLabels, testData, testLabels, nil
}
