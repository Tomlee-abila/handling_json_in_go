package functions

import (
	"fmt"
	"encoding/json"
)

func UnmarshalJson(){
	/*
	Umarshaling json into map
	*/
	data := []byte(`{"foo":"bar"}`)

	var x interface{}
	_ = json.Unmarshal(data, &x)
	// spew.Dump(x)
	fmt.Println(x)

	/*
	Umarshaling json into struct
	*/
	type person struct{
		Name string `json:"name"`
		// Age int `json:"age"`
		// Description string `json:"description"`
		// Secret string  //Unexported fields are never (un)marshaled
	}

	data1 := []byte(`{"name":"Bob","age":32,"description":"hello","Secret":"Shh"}`)

	var x1 person
	_ = json.Unmarshal(data1, &x1)
	fmt.Println(x1)
}

func UnknownUnmarshal(){

}