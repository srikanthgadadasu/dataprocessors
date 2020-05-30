// Copyright 2020 The Nadi Data Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import(
	"fmt"
	"github.com/nadidata/dataprocessors/src/readcsv"
)

func main(){
	Ipfile := "testfiles/udemy_courses.csv"
	rows := readcsv.Csvtodataframe(Ipfile)
	fmt.Println(rows[0]) //sample record
	fmt.Println(rows[0].Course_title)
}
