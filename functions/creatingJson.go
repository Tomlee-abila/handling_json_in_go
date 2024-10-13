package functions

import (
	"fmt"
	"encoding/json"
)

func MarshalJson(){
	/*
	creating json using map
	*/
	x := map[string]string{
		"foo": "bar",
	}

	data, _ := json.Marshal(x)
	fmt.Println(string(data))


	/*
	creating json using struct
	*/
	type person struct{
		Name string `json:"name"`
		Age int `json:"age"`
		Description string `json:"description"`
		Secret string  //Unexported fields are never (un)marshaled
	}

	x1 := person{
		Name: "Bob",
		Age: 32,
		Secret: "Shh",
	}

	data1, _ := json.Marshal(x1)

	fmt.Println(string(data1))
}