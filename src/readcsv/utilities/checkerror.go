package csvtodf

import(
	"fmt"
	"os"
)

func Checkerror(err error){
	if err != nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
}