package main

import (
	"fmt"
	"github/FITO3000/f3s_go_maven/maven"
)

func main() {

	dir := "/home/fit/temp/maven-test"

	maven.DeletePom(dir)

	pom1, _ := maven.CreatePom(dir)
	pom2, _ := maven.LoadPom(dir)
	pom3, err := maven.SearchPom(dir)
	if err != nil {
		panic(err)
	}

	if pom1 != nil && pom2 != nil && pom3 != nil {
		fmt.Println("OK")
	} else {
		fmt.Println("Shit!")
	}

	pom4, err := maven.SearchPom("/home/fit/temp")
	if pom4 == nil && err == nil {
		fmt.Println("OK 2")
	} else {
		fmt.Println("Shit 2!")
	}

	pom5, err := maven.SearchPom("/home/fit/temp/maven-test/u1/u2/u3")
	if pom5 != nil && err == nil {
		fmt.Println("OK 3")
	} else {
		fmt.Println("Shit 3!")
	}
}
