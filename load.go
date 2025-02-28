// Package gomnist lets you to load the MNIST data set for use with gonum package.
package gomnist

import (
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/mat"
)

// MNIST has all data & labels.
type MNIST struct {
	TrainData   mat.Matrix
	TrainLabels mat.Matrix
	TestData    mat.Matrix
	TestLabels  mat.Matrix
}

// Loader has loader setting.
type Loader struct {
	RootPath      string
	Normalization bool
	OneHot        bool
}

// OptionFunc for set loader options.
type OptionFunc func(l *Loader)

// Normalization is optional function for set normalization config.
func Normalization(normalization bool) func(l *Loader) {
	return func(l *Loader) {
		l.Normalization = normalization
	}
}

// OneHotLabel is optional function to get one-hot labels.
func OneHotLabel(oneHot bool) func(l *Loader) {
	return func(l *Loader) {
		l.OneHot = oneHot
	}
}

// NewLoader inits loader with options.
func NewLoader(rootPath string, ops ...OptionFunc) *Loader {
	l := &Loader{
		RootPath: rootPath,
	}

	for _, f := range ops {
		f(l)
	}

	return l
}

// Load loads MNIST data for gonum matrix.
func (l *Loader) Load() (MNIST, error) {
	trainSet, testSet, err := load(l.RootPath)
	if err != nil {
		return MNIST{}, errors.Wrap(err, "gomnist: failed to load mnist data")
	}

	trainData, trainLabels := set2Mat(trainSet, l.Normalization, l.OneHot)
	testData, testLabels := set2Mat(testSet, l.Normalization, l.OneHot)

	return MNIST{
		TrainData:   trainData,
		TrainLabels: trainLabels,
		TestData:    testData,
		TestLabels:  testLabels,
	}, nil
}
