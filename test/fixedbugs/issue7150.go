// errorcheck

// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// issue 7150: array index out of bounds error off by one

package main

func main() {
	_ = [0]int{-1: 50}              // ERROR "array index must be non-negative integer constant"
	_ = [0]int{0: 0}                // ERROR "array index 0 out of bounds \[0:0\]"
	_ = [0]int{5: 25}               // ERROR "array index 5 out of bounds \[0:0\]"
	_ = [10]int{2: 10, 15: 30}      // ERROR "array index 15 out of bounds \[0:10\]"
	_ = [10]int{5: 5, 1: 1, 12: 12} // ERROR "array index 12 out of bounds \[0:10\]"
}
