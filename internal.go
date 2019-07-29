package gomnist

import (
	"image/color"
	"sync"

	"github.com/petar/GoMNIST"
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

func set2Mat(s *GoMNIST.Set, normalization bool, oneHot bool) (data mat.Matrix, labels mat.Matrix) {
	d := mat.NewDense(len(s.Images), s.NRow*s.NRow, nil)
	var l *mat.Dense
	if oneHot {
		l = mat.NewDense(len(s.Labels), 10, nil)
	} else {
		l = mat.NewDense(len(s.Labels), 1, nil)
	}

	var wg sync.WaitGroup
	for i := 0; i < len(s.Images); i++ {
		wg.Add(1)
		go func(i int) {
			image, label := s.Get(i)
			b := image.Bounds()

			imageVec := make([]float64, 0, s.NRow*s.NRow)
			for n := 0; n < b.Max.Y; n++ {
				for m := 0; m < b.Max.X; m++ {
					var v float64
					if normalization {
						v = float64(image.At(n, m).(color.Gray).Y) / 255
					} else {
						v = float64(image.At(n, m).(color.Gray).Y)
					}
					imageVec = append(imageVec, v)
				}
			}

			var labelVec []float64
			if oneHot {
				labelVec = make([]float64, 10, 10)
				labelVec[int(label)] = 1
			} else {
				labelVec = make([]float64, 1, 1)
				labelVec = []float64{float64(label)}
			}

			d.SetRow(i, imageVec)
			l.SetRow(i, labelVec)

			wg.Done()
		}(i)
	}

	wg.Wait()
	return d, l
}
