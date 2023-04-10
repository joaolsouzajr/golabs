package main

import (
	"reflect"
	"testing"
)

func TestJsonParser(t *testing.T) {

	t.Run("Parser", func(t *testing.T) {
		data := `{"Name":"Joao", "Age":37 }`

		got := JsonParser(data)

		want := User{"Joao", 37}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("Parser List", func(t *testing.T) {
		data := `[{"Name":"Joao", "Age":37 }]`

		got := ListJsonParser(data)

		want := []User{{"Joao", 37}}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
