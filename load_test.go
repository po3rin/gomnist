package gomnist_test

import (
	"testing"

	"github.com/po3rin/gomnist"
)

func TestLoadMat(t *testing.T) {
	trainData, trainLabels, testData, testLabels, err := gomnist.LoadMat("./data")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if r, c := trainData.Dims(); r != 60000 || c != 784 {
		t.Fatalf("want = [r: 60000 c: 784], got = [r: %v, c: %v]", r, c)
	}
	if r, c := trainLabels.Dims(); r != 60000 || c != 1 {
		t.Fatalf("want = [r: 60000 c: 1], got = [r: %v, c: %v]", r, c)
	}
	if r, c := testData.Dims(); r != 10000 || c != 784 {
		t.Fatalf("want = [r: 10000 c: 784], got = [r: %v, c: %v]", r, c)
	}
	if r, c := testLabels.Dims(); r != 10000 || c != 1 {
		t.Fatalf("want = [r: 10000 c: 1], got = [r: %v, c: %v]", r, c)
	}

	if v := trainData.At(0, 135); v != 55 {
		t.Fatalf("want: trainData.At(0, 135) = 55, got: %v", v)
	}
	if v := trainLabels.At(0, 0); v != 5 {
		t.Fatalf("want: trainLabels.At(0, 0) = 5, got: %v", v)
	}
	if v := testData.At(0, 175); v != 84 {
		t.Fatalf("want: trainLabels.At(0, 175) = 84, got: %v", v)
	}
	if v := testLabels.At(0, 0); v != 7 {
		t.Fatalf("want: testLabels.At(0, 0) = 7, got: %v", v)
	}
}
