package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	input := `{
"name" : "Tolexo Online Pvt. Ltd",
"age_in_years" : 8.5,
"origin" : "Noida",
"head_office" : "Noida, Uttar Pradesh",
"address" : [
{
"street" : "91 Springboard",
"landmark" : "Axis Bank",
"city" : "Noida",
"pincode" : 201301, 
"state" : "Uttar Pradesh"
},
{
"street" : "91 Springboard",
"landmark" : "Axis Bank",
"city" : "Noida",
"pincode" : 201301,
"state" : "Uttar Pradesh"
}
],
"sponsers" : {
"name" : "One"
},
"revenue" : "19.8 million$",
"no_of_employee" : 630,
"str_text" : ["one","two"],
"int_text" : [1,3,4]
}`

	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	recursiveIterator(data)

}

func recursiveIterator(data interface{}) {
	val := reflect.ValueOf(data)

	if !val.IsValid() {
		fmt.Println("Type: nil, Value: nil")
		return
	}

	fmt.Printf("Type: %v, Value: %v\n", val.Kind(), val)

	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			fmt.Printf("Key: %v -> \n", key)
			recursiveIterator(val.MapIndex(key).Interface())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			fmt.Printf("Index: %d -> \n", i)
			recursiveIterator(val.Index(i).Interface())
		}
	}
}
