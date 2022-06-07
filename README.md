# kigen: Go 1.18+ Generics Libraries, for GoKi framework

[![Go Report Card](https://goreportcard.com/badge/github.com/goki/kigen)](https://goreportcard.com/report/github.com/goki/kigen)
[![Go Reference](https://pkg.go.dev/badge/github.com/goki/kigen.svg)](https://pkg.go.dev/github.com/goki/kigen)
[![CI](https://github.com/goki/kigen/actions/workflows/ci.yml/badge.svg)](https://github.com/goki/kigen/actions/workflows/ci.yml)
[![Codecov](https://codecov.io/gh/goki/kigen/branch/master/graph/badge.svg?token=Hw5cInAxY3)](https://codecov.io/gh/goki/kigen)

This collection of Generics libraries takes advantage of the new generic type parameters introduced in Go 1.18.

# ordmap: ordered map

Implements an ordered map that retains the order of items added to a slice, while also providing fast key-based map lookup of items.

# dedupe: deduplicate (uniquify) any slice with comparable elements

Implements a de-duplication function for any comparable slice type, efficiently using a map to check for duplicates.  The original order of items is preserved.


