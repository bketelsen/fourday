package main

import (
	"fmt"
)

func main() {

	things := map[string]interface{}{
		"string": "stringval",
		"bool":   true,
		"int":    5,
		"struct": struct{}{}, // inline declaration of a struct with no fields
	}

	for k, v := range things {
		thingType(k, v)
	}

}

func thingType(k string, v interface{}) {

	switch t := v.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case string:
		fmt.Printf("string %s\n", t) // t has type string
	}
}
