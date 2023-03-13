package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string // 后面这部分统称为结构体成员Tag
	Year   int    `json:"released"`
	Color  bool   `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
}

func main() {
	unMar()
}

func unMar() {
	// 通过一致的结构。来选择性的解码 json中有兴趣的成员
	var titles []struct{ Title string }
	data, _ := json.MarshalIndent(movies, "", "  ") // 缩进编组json
	fmt.Printf("%v \n", string(data))
	json.Unmarshal(data, &titles)
	fmt.Println(titles)
}

// 比json编码
func mar() {
	fmt.Println(movies[0])
	fmt.Println(movies[1])
	fmt.Println(movies[2])
	//data, _ := json.Marshal(movies[0]) // 正常编组json
	data, _ := json.MarshalIndent(movies[0], "", "  ") // 缩进编组json
	fmt.Printf("%v", string(data))
}
