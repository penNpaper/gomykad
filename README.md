# gomykad

![Go builds](https://github.com/ShiraazMoollatjie/gomykad/workflows/Go%20builds/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ShiraazMoollatjie/gomykad?style=flat-square)](https://goreportcard.com/report/github.com/ShiraazMoollatjie/gomykad)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/ShiraazMoollatjie/gomykad/pkg/mykad)

A package to generate, validate and represent Malaysian Identity Card (MyKAD) numbers in Go.

# Import
To import the package use:
```
 go get github.com/ShiraazMoollatjie/gomykad/pkg/mykad
```

# Representation
You can represent a MyKAD using a formatted or unformatted NRIC number. A represented MyKAD number can be assumed to 
be validated.

## Formatted NRIC
```go
m, err := NewMyKAD("721212-24-3221")
```

## Unformatted NRIC
```go
m, err := NewMyKAD("721212243221")
```

# Validation
Validates a MyKAD NRIC number. 

```go
err := Validate("721212-24-3221")
```

# Generation
It is also possible to generate random MyKAD numbers. This is very useful if you need to test random MyKAD numbers. 

```go
m := Generate()
```
