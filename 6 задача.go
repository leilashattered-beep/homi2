package main

import "fmt"

func tags(posts [][]string) {
	tags := make(map[string]bool)

	for i := 0; i < len(posts); i++ {

		for j := 0; j < len(posts[i]); j++ {

			tags[posts[i][j]] = true
		}
	}

	for tag := range tags {

		fmt.Println(tag)
	}
}

func main() {

	posts := [][]string{

		{"go", "backend"},
		{"git", "go", "tools"},
	}

	tags(posts)
}
