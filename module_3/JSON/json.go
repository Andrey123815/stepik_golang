package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonInputFormat struct {
	Students []struct {
		Rating []float64
	}
}

type jsonOutputFormat struct {
	Average float64
}

func main() {
	var s jsonInputFormat
	data, _ := ioutil.ReadAll(os.Stdin)
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}
	average := 0.0
	for i := 0; i < len(s.Students); i++ {
		average += float64(len(s.Students[i].Rating))
	}
	s_out := jsonOutputFormat{Average: average / float64(len(s.Students))}
	data_output, _ := json.MarshalIndent(s_out, "", "    ")
	fmt.Printf("%s", data_output)
}
