package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robertding/The-Go-Programming-Language/ch4/4-10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Println("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
