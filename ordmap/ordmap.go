// Copyright (c) 2022, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
package ordmap implements an ordered map that retains the order of items
added to a slice, while also providing fast key-based map lookup of items,
using the Go 1.18 generics system.

The implementation is fully visible and the API provides a minimal
subset of methods, compared to other implementations that are heavier,
so that additional functionality can be added as needed.

The slice structure holds the Key and Val for items as they are added,
enabling direct updating of the corresponding map, which holds the
index into the slice.
*/
package ordmap

import "golang.org/x/exp/slices"

// KeyVal represents the Key and Value
type KeyVal[K comparable, V any] struct {
	Key K
	Val V
}

// Map is a generic ordered map that combines the order of a slice
// and the fast key lookup of a map.  A map stores an index
// into a slice that has the value and key associated with the value.
type Map[K comparable, V any] struct {
	Order []*KeyVal[K, V] `desc:"ordered list of values and associated keys -- in order added"`
	Map   map[K]int       `desc:"key to index mapping"`
}

// New returns a new ordered map
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		Map: make(map[K]int),
	}
}

// Add adds a new value for given key
func (om *Map[K, V]) Add(key K, val V) {
	om.Map[key] = len(om.Order)
	om.Order = append(om.Order, &KeyVal[K, V]{Key: key, Val: val})
}

// InsertAtIdx inserts value with key at given index
func (om *Map[K, V]) InsertAtIdx(idx int, key K, val V) {
	om.Map[key] = idx
	om.Order = slices.Insert(om.Order, idx, &KeyVal[K, V]{Key: key, Val: val})
}

// ValByKey returns value based on Key, along with bool reflecting
// presence of key.
func (om *Map[K, V]) ValByKey(key K) (V, bool) {
	idx, ok := om.Map[key]
	if ok {
		return om.Order[idx].Val, ok
	}
	var zv V
	return zv, false
}

// IdxByKey returns index of given Key, along with bool reflecting
// presence of key.
func (om *Map[K, V]) IdxByKey(key K) (int, bool) {
	idx, ok := om.Map[key]
	return idx, ok
}

// ValByIdx returns value at given index, in ordered slice.
func (om *Map[K, V]) ValByIdx(idx int) V {
	return om.Order[idx].Val
}

// KeyByIdx returns key for given index, in ordered slice.
func (om *Map[K, V]) KeyByIdx(idx int) K {
	return om.Order[idx].Key
}

// Len returns the number of items in the map
func (om *Map[K, V]) Len() int {
	return len(om.Order)
}

// DeleteIdx deletes item at given index
func (om *Map[K, V]) DeleteIdx(idx int) {
	kv := om.Order[idx]
	delete(om.Map, kv.Key)
	slices.Delete(om.Order, idx, idx+1)
}

// DeleteKey deletes item by given key, returns true if found
func (om *Map[K, V]) DeleteKey(key K) bool {
	idx, ok := om.Map[key]
	if !ok {
		return false
	}
	slices.Delete(om.Order, idx, idx+1)
	return true
}
