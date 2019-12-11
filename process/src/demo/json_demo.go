package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	names := []string{"xiaotian", "age", "test"}
	nameJSON, _ := json.Marshal(names)
	fmt.Println(string(nameJSON))

	person := map[string]int{"age": 24, "children": 2}
	personJSON, _ := json.Marshal(person)
	fmt.Println(string(personJSON))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	res := &Response2{}
	resbyte := []byte(`{"page":1,"fruits":["apple","peach","pear"]}`)
	if err := json.Unmarshal(resbyte, &res); err != nil {
		panic(err)
	}
	fmt.Println("res: ", res)
	fmt.Println("res.Page: ", res.Page)
	fmt.Println("res.Fruits: ", res.Fruits)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
