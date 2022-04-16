# rands

Random string generator in Go

## as CLI

### install

```bash
% go install github.com/m-mizutani/rands/cmd/rands@latest
```

### Usage
No option (default length is 12 and char set consists of alphabet and number)
```bash
% srand
5cGQwT01CPuu
```

Change string length
```bash
% srand -l 128
Khovsqks6MkuHeDjxBiNHs8wRNZtGDh7ia3gNF8EKJ6rJYPWlHg52VLWEaD9E4Cib2ww3sAKvvqRoOQGeBlWlD06roqyAC2QsdCJYpHEv1gDa9b8ic5FE7NoIvdkuhLO
```

Choose characters to be used for random string (multi-byte chars are also acceptable)
```bash
% srand --chars="晴雨曇"
晴晴曇曇晴曇晴曇曇晴雨晴
```

Use predefined character set
```bash
% srand --use-lowers --use-numbers --use-marks
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
r := srand.New()
fmt.Println(r.NewString(12))

// use global variable
fmt.Println(srand.NewString())

// With charset
fmt.Println(srand.New(srand.WithCharSet("ABCDEFG")).NewString())

// With seed
fmt.Println(srand.New(srand.WithSeed(666)).NewString(12))
// Must output: uwSttwcqbmIh
```
