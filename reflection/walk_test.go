package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("without maps", func(t *testing.T) {
		cases := []struct {
			CaseName string
			Input    interface{}
			Expected []string
		}{
			{
				"struct with one field",
				struct {
					Name string
				}{"John Smith"},
				[]string{"John Smith"},
			},
			{
				"struct with two fields",
				struct {
					Name string
					City string
				}{"John Smith", "New York"},
				[]string{"John Smith", "New York"},
			},
			{
				"struct with no string field",
				struct {
					Name string
					Age  int
				}{"John Smith", 33},
				[]string{"John Smith"},
			},
			{
				"struct with nested struct",
				Person{"John Smith", Profile{33, "New York"}},
				[]string{"John Smith", "New York"},
			},
			{
				"struct with nested struct",
				&Person{"John Smith", Profile{33, "New York"}},
				[]string{"John Smith", "New York"},
			},
			{
				"slices of structs",
				[]Profile{
					{53, "Paris"},
					{33, "Berlin"},
				},
				[]string{"Paris", "Berlin"},
			},
			{
				"arrays of structs",
				[2]Profile{
					{53, "Paris"},
					{33, "Berlin"},
				},
				[]string{"Paris", "Berlin"},
			},
		}

		for _, testCase := range cases {
			var result []string

			t.Run(testCase.CaseName, func(t *testing.T) {
				Walk(testCase.Input, func(val string) {
					result = append(result, val)
				})
			})

			if !reflect.DeepEqual(testCase.Expected, result) {
				t.Errorf("Expected %v but received %v", testCase.Expected, result)
			}
		}
	})
	t.Run("with maps", func(t *testing.T) {
		input := map[string]string{
			"France":  "Paris",
			"Germany": "Berlin",
		}

		var result []string

		Walk(input, func(input string) {
			result = append(result, input)
		})

		AssertContains(t, result, "Paris")
		AssertContains(t, result, "Berlin")
	})
	t.Run("channel", func(t *testing.T) {
		someChannel := make(chan Profile)

		go func() {
			someChannel <- Profile{53, "Paris"}
			someChannel <- Profile{33, "Berlin"}
			close(someChannel)
		}()

		var result []string
		expected := []string{"Paris", "Berlin"}

		Walk(someChannel, func(input string) {
			result = append(result, input)
		})

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v but received %v", expected, result)
		}
	})
	t.Run("function", func(t *testing.T) {
		someFunc := func() (Profile, Profile) {
			return Profile{33, "Paris"}, Profile{43, "Berlin"}
		}

		var result []string
		expected := []string{"Paris", "Berlin"}

		Walk(someFunc, func(input string) {
			result = append(result, input)
		})

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v but received %v", expected, result)
		}
	})
}

func AssertContains(t testing.TB, result []string, expected string) {
	t.Helper()

	contains := false

	for _, val := range result {
		if expected == val {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("Expected to contain %q but it doesn't", expected)
	}
}
