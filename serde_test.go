package serde

import "testing"

func TestSerde(t *testing.T) {

	t.Run("write strict to json file", func(t *testing.T) {
		type MyStruct struct {
			Name string
			Age  int
		}

		type MyStruct2 struct {
			Name string `json:"name,omitempty"`
			Age  int    `json:"age,omitempty"`
			Pet  string `json:"pet,omitempty"`
		}

		n := MyStruct{Name: "nikitous", Age: 39}
		p := MyStruct2{Name: "nikitoyulis", Age: 39, Pet: "miko"}

		err := WriteStructToFileAsJson("", n)
		if err != nil {
			t.Error(err)
		}

		err = WriteStructToFileAsJson("", p)
		if err != nil {
			t.Error(err)
		}

	})

}
