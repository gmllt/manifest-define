package main

import "fmt"

func main()  {
	var a Args
	a.ParseArgs()
	var myString = "b0.mts.sophia"
	var site Site
	site.Trigram = "sph"
	newString := site.Replace(myString)
	fmt.Println(newString)
}