package main
import (
	"fmt"
	"io/ioutil"
)


func main() {
	bytes,err := ioutil.ReadFile("input.txt")
	if err == nil {
		fmt.Println(bytes)
	} else {
		fmt.Println(err)
	}
}
