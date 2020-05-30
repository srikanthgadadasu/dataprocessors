// Copyright 2020 The Nadi Data Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package readcsv

import(
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"time"
	"strconv"
	"bufio"
)

//Iptype is input csv file schema (created for sample file(udemy_courses.csv) located in testfiles directory). 
//This needs to be moved to another file and pass the file and pass it as parameter to Csvtodataframe function
type Iptype struct{
	Course_id string
	Course_title string
	Url	string
	Is_paid	bool
	Price float64
	Num_subscribers int	
	Num_reviews	int
	Num_lectures int
	Level string
	Content_duration float64	
	Published_timestamp time.Time
	Subject string
}

//Csvtodataframe function converts csv data into data frame format using the schema
func Csvtodataframe(Ipfile string) []Iptype {

	csvfile, err := os.Open(Ipfile)
	checkerror(err)
	defer csvfile.Close()
	
	// Skip first row (line)
    row1, err := bufio.NewReader(csvfile).ReadSlice('\n')
    checkerror(err)
	_, err = csvfile.Seek(int64(len(row1)), io.SeekStart)
	checkerror(err)
	
	//read remaining lines
	data := csv.NewReader(csvfile)
	rawdata, err := data.ReadAll()
	checkerror(err)

	var line Iptype
	var lines []Iptype
	layout := time.RFC3339

	for _, record := range rawdata {
		//Below columns assignment must be generic, as column names come from file and assignments should be generic
		line.Course_id = record[0]
		checkerror(err)
		line.Course_title = record[1]
		checkerror(err)
		line.Url = record[2]
		checkerror(err)
		line.Is_paid, err = strconv.ParseBool(record[3])
		checkerror(err)
		line.Price, err = strconv.ParseFloat(record[4], 64)
		checkerror(err)
		line.Num_subscribers, err = strconv.Atoi(record[5])
		checkerror(err)
		line.Num_reviews, err = strconv.Atoi(record[6])
		checkerror(err)
		line.Num_lectures, err = strconv.Atoi(record[7])
		checkerror(err)
		line.Level = record[8]
		checkerror(err)
		line.Content_duration, err = strconv.ParseFloat(record[9], 64)
		checkerror(err)
		line.Published_timestamp, err = time.Parse(layout, record[10])
		checkerror(err)
		line.Subject = record[11]
		checkerror(err)

		lines = append(lines, line)
	}
	return lines
}

func checkerror(err error){
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
}
