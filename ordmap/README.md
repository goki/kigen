# ordmap: ordered map using Go 1.18 generics

[![Go Reference](https://pkg.go.dev/badge/github.com/goki/kigen.svg)](https://pkg.go.dev/github.com/goki/kigen/ordmap)

Package `ordmap` implements an ordered map that retains the order of items added to a slice, while also providing fast key-based map lookup of items, using the Go 1.18 generics system.

The implementation is fully visible and the API provides a minimal subset of methods, compared to other implementations that are heavier, so that additional functionality can be added as needed.  Iteration can be performed directly on the `Order` using standard Go `range` function.

The slice structure holds the Key and Val for items as they are added, enabling direct updating of the corresponding map, which holds the index into the slice.

