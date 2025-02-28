<p align="center">
  <img src="./static/gomnist.png" width="480px">
  <h3 align="center">gomnist</h3>
  <p align="center">This package lets you to load the MNIST data set for use with gonum package. It is useful when implementing, for example, deep learning using the gonum package.</p>
</p>

---
<img src="https://img.shields.io/badge/go-v1.12-blue.svg"/> [![CircleCI](https://circleci.com/gh/po3rin/gomnist.svg?style=shield)](https://circleci.com/gh/po3rin/llb2dot) [![GolangCI](https://golangci.com/badges/github.com/po3rin/gomnist.svg)](https://golangci.com) [![Maintainability](https://api.codeclimate.com/v1/badges/7c29dd3d56a623ba729e/maintainability)](https://codeclimate.com/github/po3rin/gomnist/maintainability)

## What's MNIST ??

<img src="./static/mnist.png">

[THE MNIST DATABASE](http://yann.lecun.com/exdb/mnist/)

 > The MNIST database of handwritten digits, available from this page, has a training set of 60,000 examples, and a test set of 10,000 examples. It is a subset of a larger set available from NIST. The digits have been size-normalized and centered in a fixed-size image. It is a good database for people who want to try learning techniques and pattern recognition methods on real-world data while spending minimal efforts on preprocessing and formatting.

## Quick Start

First of all, you should download mnist file.
[THE MNIST DATABASE](http://yann.lecun.com/exdb/mnist/)

```bash
# exmple
.
└── testdata
    ├── t10k-images-idx3-ubyte.gz
    ├── t10k-labels-idx1-ubyte.gz
    ├── train-images-idx3-ubyte.gz
    └── train-labels-idx1-ubyte.gz
```

Using gomnist, you can get MNIST data as gonum matrix.

```go
package main

import "github.com/po3rin/gomnist"

func main() {
    // first arg is target dir has mnist file.
    l := gomnist.NewLoader("./testdata")

    // Do !!
    mnist, err := l.Load()
    if err != nil {
      // error handling ...
    }

    // type MNIST struct {
    //   TrainData   mat.Matrix
    //   TrainLabels mat.Matrix
    //   TestData    mat.Matrix
    //   TestLabels  mat.Matrix
    // }
    _ = mnist.TrainData.At(0, 135)
    // 55
}
```

## Options

gomnist options is implimented as Functional Option Pattern.

#### Normalization

Normalization Option is whether to normalize the input image value to a value between 0 and 1 (Default false)

```go
package main

import "github.com/po3rin/gomnist"

func main() {
    l := gomnist.NewLoader("./testdata", gomnist.Normalization(true))
    mnist, err := l.Load()
    if err != nil {
      // error handling ...
    }
    _ = mnist.TrainData.At(0, 135)
    // 0.21568627450980393
}
```

#### One-Hot Label

OneHotLabel Option is whether to set one-hot label.

```go
package main

import "github.com/po3rin/gomnist"

func main() {
    l := gomnist.NewLoader("./testdata", gomnist.OneHotLabel(true))
    mnist, err := l.Load()
    if err != nil {
      // error handling ...
    }
    _ = mnist.TrainLabel
    // ⎡0 0 0 0 0 1 0 0 0 0 0⎤
    // ⎢          .          ⎥
    // ⎣          .          ⎦
}
```

## Dimension

#### Default

(Number of images) * (Total number of pixels : 28*28)
* trainData:   60000 - 784
* testData:    10000 - 784

(Number of images) * (Handwritten digits value)
* trainLabels: 60000 - 1
* testLabels:  10000 - 1

#### One-Hot

(Number of images) * (Handwritten digits value)
* trainLabels: 60000 - 10
* testLabels:  10000 - 10

## TODO
* Download if mnist file do not exits
