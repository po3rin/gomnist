package gomnist_test

import (
	"testing"

	"github.com/po3rin/gomnist"
)

func TestLoadMat(t *testing.T) {

	type mat struct {
		r, c, i, j int
		v          float64
	}
	tests := []struct {
		name          string
		normalization bool
		onehot        bool
		trainData     mat
		trainLabels   mat
		testData      mat
		testLabels    mat
	}{
		{
			name: "normal",
			trainData: mat{
				r: 60000,
				c: 784,
				i: 0,
				j: 135,
				v: 55,
			},
			trainLabels: mat{
				r: 60000,
				c: 1,
				i: 0,
				j: 0,
				v: 5,
			},
			testData: mat{
				r: 10000,
				c: 784,
				i: 0,
				j: 175,
				v: 84,
			},
			testLabels: mat{
				r: 10000,
				c: 1,
				i: 0,
				j: 0,
				v: 7,
			},
		},
		{
			name:          "with normalization",
			normalization: true,
			trainData: mat{
				r: 60000,
				c: 784,
				i: 0,
				j: 135,
				v: 0.21568627450980393,
			},
			trainLabels: mat{
				r: 60000,
				c: 1,
				i: 0,
				j: 0,
				v: 5,
			},
			testData: mat{
				r: 10000,
				c: 784,
				i: 0,
				j: 175,
				v: 0.32941176470588235,
			},
			testLabels: mat{
				r: 10000,
				c: 1,
				i: 0,
				j: 0,
				v: 7,
			},
		},
		{
			name:   "with one-hot label",
			onehot: true,
			trainData: mat{
				r: 60000,
				c: 784,
				i: 0,
				j: 135,
				v: 55,
			},
			trainLabels: mat{
				r: 60000,
				c: 10,
				i: 0,
				j: 5,
				v: 1,
			},
			testData: mat{
				r: 10000,
				c: 784,
				i: 0,
				j: 175,
				v: 84,
			},
			testLabels: mat{
				r: 10000,
				c: 10,
				i: 0,
				j: 7,
				v: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := gomnist.NewLoader(
				"./testdata",
				gomnist.Normalization(tt.normalization),
				gomnist.OneHotLabel(tt.onehot),
			)
			mnist, err := l.Load()

			trainData := mnist.TrainData
			trainLabels := mnist.TrainLabels
			testData := mnist.TestData
			testLabels := mnist.TestLabels

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if r, c := trainData.Dims(); r != tt.trainData.r || c != tt.trainData.c {
				t.Errorf("want = [r: %v c: %v], got = [r: %v, c: %v]", tt.trainData.r, tt.trainData.c, r, c)
			}
			if r, c := trainLabels.Dims(); r != tt.trainLabels.r || c != tt.trainLabels.c {
				t.Errorf("want = [r: %v c: %v], got = [r: %v, c: %v]", tt.testLabels.r, tt.testLabels.c, r, c)
			}
			if r, c := testData.Dims(); r != tt.testData.r || c != tt.testData.c {
				t.Errorf("want = [r: %v c: %v], got = [r: %v, c: %v]", tt.testData.r, tt.testData.c, r, c)
			}
			if r, c := testLabels.Dims(); r != tt.testLabels.r || c != tt.testLabels.c {
				t.Errorf("want = [r: %v c: %v], got = [r: %v, c: %v]", tt.testLabels.r, tt.testLabels.c, r, c)
			}

			if v := trainData.At(tt.trainData.i, tt.trainData.j); v != tt.trainData.v {
				t.Errorf("want: trainData.At(%v, %v) = %v, got: %v", tt.trainData.i, tt.trainData.j, tt.trainData.v, v)
			}
			if v := trainLabels.At(tt.trainLabels.i, tt.trainLabels.j); v != tt.trainLabels.v {
				t.Errorf("want: trainLabels.At(%v, %v) = %v, got: %v", tt.trainLabels.i, tt.trainLabels.j, tt.trainLabels.v, v)
			}
			if v := testData.At(tt.testData.i, tt.testData.j); v != tt.testData.v {
				t.Errorf("want: trainLabels.At(%v, %v) = %v, got: %v", tt.testData.i, tt.testData.j, tt.testData.v, v)
			}
			if v := testLabels.At(tt.testLabels.i, tt.testLabels.j); v != tt.testLabels.v {
				t.Errorf("want: testLabels.At(%v, %v) = %v, got: %v", tt.testLabels.i, tt.testLabels.j, tt.testLabels.v, v)
			}
		})
	}
}
