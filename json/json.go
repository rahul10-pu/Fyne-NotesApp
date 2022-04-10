package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2/data/binding"
)

func LoadJsonData() []string {
	fmt.Println("Loading data from JSON file")

	input, _ := ioutil.ReadFile("data.json")
	var data []string
	json.Unmarshal(input, &data)

	return data
}

func SaveJsonData(data binding.StringList) {
	fmt.Println("Saving data to JSON file")
	d, _ := data.Get()
	jsonData, _ := json.Marshal(d)
	ioutil.WriteFile("data.json", jsonData, 0644)

}
