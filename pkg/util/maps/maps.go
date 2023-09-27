/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// replacing this with `https://pkg.go.dev/golang.org/x/exp/maps` should be considered
// when `x/exp/maps` graduates to stable.
package maps

import (
	"maps"
)

// Merge merges a and b while resolving the conflicts by calling commonKeyValue
func Merge[K comparable, V any, S ~map[K]V](a, b S, commonKeyValue func(a, b V) V) S {
	if a == nil {
		return maps.Clone(b)
	}

	ret := maps.Clone(a)

	for k, v := range b {
		if _, found := a[k]; found {
			ret[k] = commonKeyValue(a[k], v)
		} else {
			ret[k] = v
		}
	}
	return ret
}

// Merge returns the intersection of a and b with the values generated by calling commonKeyValue
func Intersect[K comparable, V any, M ~map[K]V](a, b M, commonKeyValue func(a, b V) V) M {
	if a == nil || b == nil {
		return nil
	}

	ret := make(M)

	for k, v := range b {
		if _, found := a[k]; found {
			ret[k] = commonKeyValue(a[k], v)
		}
	}
	return ret
}

// Merge merges a and b keeping the values in a in case of conflict
func MergeKeepFirst[K comparable, V any, S ~map[K]V](a, b S) S {
	return Merge(a, b, func(v, _ V) V { return v })
}

// Contains returns true if a contains all the keys in b with the same value
func Contains[K, V comparable, A ~map[K]V, B ~map[K]V](a A, b B) bool {
	for k, bv := range b {
		if av, found := a[k]; !found || av != bv {
			return false
		}
	}
	return true
}

// Keys returns a slice containing the m keys
func Keys[K comparable, V any, M ~map[K]V](m M) []K {
	if m == nil {
		return nil
	}
	ret := make([]K, 0, len((m)))

	for k := range m {
		ret = append(ret, k)
	}
	return ret
}
