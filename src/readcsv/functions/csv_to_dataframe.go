// Copyright 2020 The Nadi Authors. All rights reserved.
// Use of this source code is not permitted

package functions

import(
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"time"
	"strconv"
	"bufio"
)

type iptype struct{
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
func Csvtodataframe(){
	
	ipfile := "/Users/sripri/Documents/nadi/udemy_courses.csv"

	csvfile, err := os.Open(ipfile)
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

	var line iptype
	var lines []iptype
	layout := time.RFC3339

	for _, record := range rawdata {
		line.Course_id = record[0]
		line.Course_title = record[1]
		line.Url = record[2]
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
		line.Content_duration, err = strconv.ParseFloat(record[9], 64)
		checkerror(err)
		line.Published_timestamp, err = time.Parse(layout, record[10])
		checkerror(err)
		line.Subject = record[11]

		lines = append(lines, line)
	}
	fmt.Println(lines[0])

	//{1070968 Ultimate Investment Banking Course https://www.udemy.com/ultimate-investment-banking-course/ true 200 2147 23 51 All Levels 1.5 2017-01-18 20:58:58 +0000 UTC Business Finance}
}

func checkerror(err error){
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
}
