package main

import "fmt"
import "regexp"

func main() {

	match, _ := regexp.MatchString("h([a-z]+)ch", "hitchhike")
	fmt.Println(match)

	r, _ := regexp.Compile("h([a-z]+)ch")

	fmt.Println(r.MatchString("hitchhiker"))
	fmt.Println(r.FindString("hitchhiker's guide"))
	fmt.Println(r.FindStringIndex("hitchhiker's guide"))

}
