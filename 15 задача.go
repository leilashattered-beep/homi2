package main

import "fmt"

type Movie struct {
	Title string
	Year int
	Rating float64
	Genres []string
}

func max(m []Movie) Movie {
	x := m[0]
	for i := 1; i < len(m); i++ {
		if m[i].Rating > x.Rating {
			x = m[i]
		}
	}
	return x
}

func main () {
	m := []Movie {
		{"Годзилла", 2014 ,7.0, []string{"Фантастика", "Боевик"}},
		{"Оно", 2017, 7.3, []string{"Ужасы" , "Драма"}},
		{"Чужой", 1979, 8.1, []string{"Ужасы", "Фантастика"}},
		{"Сияние", 1980, 7.8, []string{"Ужасы", "Драма"}},
		{"Матрица", 1999, 8.7,[]string{"Фантастика", "Боевик"}},

	}
	m[0].Genres = append(m[0].Genres, "Ужасы")

	x := max(m)
	fmt.Println("Лучший:" , x.Title, x.Rating) 

	fmt.Print("Жанр: ")
	var g string
	fmt.Scan(&g)

	for i := 0; i < len(m); i++ {
		for j :=0; j< len(m[i].Genres); j++ {
			if m[i].Genres[j] == g {
				fmt.Println(m[i].Title)
			}
		}
	}
}