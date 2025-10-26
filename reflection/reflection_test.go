package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Tabriz"},
			[]string{"Tabriz"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Tabriz", "Baku"},
			[]string{"Tabriz", "Baku"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Tabriz", 33},
			[]string{"Tabriz"},
		},
		{
			"nested fields",
			Person{
				"Tabriz",
				Profile{33, "Baku"},
			},
			[]string{"Tabriz", "Baku"},
		},
		{
			"pointers to things",
			&Person{
				"Tabriz",
				Profile{33, "Baku"},
			},
			[]string{"Tabriz", "Baku"},
		},
		{
			"slices",
			[]Profile{
				{33, "Baku"},
				{34, "Gadabay"},
			},
			[]string{"Baku", "Gadabay"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Baku"},
				{34, "Gadabay"},
			},
			[]string{"Baku", "Gadabay"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Baku"}
			aChannel <- Profile{34, "Gadabay"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Baku", "Gadabay"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Baku"}, Profile{34, "Gadabay"}
		}

		var got []string
		want := []string{"Baku", "Gadabay"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
