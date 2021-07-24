package tests

import "testing"

func Test_Transact(t *testing.T) {
	testTable := []struct {
		Name string
	}{
		{
			Name: "",
		},
	}
	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {

		})
	}
}
