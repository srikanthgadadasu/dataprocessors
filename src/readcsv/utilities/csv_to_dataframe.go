package csvtodf

import(
	"encoding/csv"
	"fmt"
	"os"
)

func Csvtodataframe(Ipfile, Iptype){
	
	csvfile, err := os.Open(Ipfile)
	checkerror(err)
	defer csvfile.Close()

	data := csv.NewReader(csvfile)
	rawdata, err := data.ReadAll()
	checkerror(err)

	var line Iptype
	var lines []Iptype

	for _, record := range rawdata {
		line.course_id = record[0]
		line.course_title = record[1]
		line.url = record[2]
		line.is_paid = record[3]
		line.price = record[4]
		line.num_subscribers = record[5]
		line.num_reviews = record[6]
		line.num_lectures = record[7]
		line.level = record[8]
		line.content_duration = record[9]
		line.published_timestamp = record[10]
		line.subject = record[11]

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