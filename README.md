gopt
===

[![Build Status](https://travis-ci.org/tatsy/gopt.svg?branch=master)](https://travis-ci.org/tatsy/gopt)
[![Coverage Status](https://coveralls.io/repos/github/tatsy/gopt/badge.svg)](https://coveralls.io/github/tatsy/gopt)

> Physically based path tracer implemented with Go.

## Usage

```sh
export GOPATH=`pwd`
go build ./src/main.go
./main -i [Scene JSON file]
```

## Result

#### Gopher

<img src="./results/gopher.jpg" alt="gopher.jpg" width="500"/>

#### Cornell box

<img src="./results/cbox.jpg" alt="cbox.jpg" width="500"/>

The scene courtesy of Mitsuba renderer (W. Jakob 2010).

## Copyright

MIT License 2017 (c) Tatsuya Yatagawa (tatsy)
