package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	c "learning/golang/json-handling/constant"
	"learning/golang/json-handling/model"
	"log"
)

func main() {
	//librariesJSONPath := "test-data/libraries.json"
	libraryJSON := stringToJSONObject(c.LibrariesJSONPath)

	JSONToString(libraryJSON)

	jsonStringToStringObject(c.LibrariesJSONPath)
}

func readjsonFile(path string) (byteValue []byte, err error) {
	return ioutil.ReadFile(path)
}

func stringToJSONObject(path string) model.Libraries {

	byteValue, errReadFile := readjsonFile(path)
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}

	var libraries model.Libraries
	errJSONUnmarshall := json.Unmarshal(byteValue, &libraries)
	if errJSONUnmarshall != nil {
		log.Fatal(errJSONUnmarshall)
	}

	for i := 0; i < len(libraries.Libraries); i++ {
		fmt.Println("Id: ", libraries.Libraries[i].ID)
		fmt.Println("Name: ", libraries.Libraries[i].Name)
		fmt.Println("City: ", libraries.Libraries[i].City)
	}

	return libraries
}

// JSONToString ...
func JSONToString(libraries model.Libraries) {

	marshalledJSON, err := json.Marshal(libraries)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(marshalledJSON))

	type Libraries struct {
		Name string
	}

	library := &Libraries{Name: "The Library"}
	b, err := json.Marshal(library)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

func jsonStringToStringObject(path string) {

	byteValue, errReadFile := readjsonFile(path)
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}

	// Read unstructure data, will map all data
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["libraries"])

}
