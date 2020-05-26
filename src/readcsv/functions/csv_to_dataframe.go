package functions

import(
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"strconv"
)

type iptype struct{
	Course_id int
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
	
	ipfile := "/udemy_courses.csv"

	csvfile, err := os.Open(ipfile)
	checkerror(err)
	defer csvfile.Close()

	data := csv.NewReader(csvfile)
	rawdata, err := data.ReadAll()
	checkerror(err)

	var line iptype
	var lines []iptype
	layout := "2006-01-02T15:04:05.000Z"

	for _, record := range rawdata {
		line.Course_id, err = strconv.Atoi(record[0])
		checkerror(err)
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
	fmt.Println(lines)
}

func checkerror(err error){
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
}
