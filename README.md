# rands [![gosec](https://github.com/m-mizutani/rands/actions/workflows/gosec.yml/badge.svg)](https://github.com/m-mizutani/rands/actions/workflows/gosec.yml) [![trivy](https://github.com/m-mizutani/rands/actions/workflows/trivy.yml/badge.svg)](https://github.com/m-mizutani/rands/actions/workflows/trivy.yml) [![lint](https://github.com/m-mizutani/rands/actions/workflows/lint.yml/badge.svg)](https://github.com/m-mizutani/rands/actions/workflows/lint.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/m-mizutani/rands.svg)](https://pkg.go.dev/github.com/m-mizutani/rands)

Random string generator in Go

## as CLI

### install

```bash
% go install github.com/m-mizutani/rands/cmd/rands@latest
```

### Usage
No option (default length is 12 and char set consists of alphabet and number)
```bash
% rands
5cGQwT01CPuu
```

Change string length
```bash
% rands -l 128
Khovsqks6MkuHeDjxBiNHs8wRNZtGDh7ia3gNF8EKJ6rJYPWlHg52VLWEaD9E4Cib2ww3sAKvvqRoOQGeBlWlD06roqyAC2QsdCJYpHEv1gDa9b8ic5FE7NoIvdkuhLO
```

Choose characters to be used for random string (multi-byte chars are also acceptable)
```bash
% rands --chars="晴雨曇"
晴晴曇曇晴曇晴曇曇晴雨晴
```

Use predefined character set
```bash
% rands --use-lowers --use-numbers --use-marks
-,6rf*_-htnm
```

## as Package

### Install

```bash
% go get github.com/m-mizutani/rands
```

### Usage

```go
// create instance and generate string
r := rands.New()
fmt.Println(r.NewString(12))

// use global variable
fmt.Println(rands.NewString())

// With charset
fmt.Println(rands.New(rands.WithCharSet("ABCDEFG")).NewString())

// With seed
fmt.Println(rands.New(rands.WithSeed(666)).NewString(12))
// Must output: uwSttwcqbmIh
```

## Test

```bash
# unit test
% go test ./test

# fuzzing
% go test -fuzz Fuzz -fuzztime=10s ./test/
```

## License

Apache License 2.0
