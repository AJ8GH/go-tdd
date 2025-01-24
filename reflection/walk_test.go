package reflection

import (
	"reflect"
	"testing"
)

type (
	simpleThing struct {
		Thing string
		Num   int
	}

	nestedThing struct {
		ThingString string
		InnerThing  innerThing
	}

	innerThing struct {
		InnerThingString string
		Num              int
	}
)

func TestWalk(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  []string
	}{
		{
			"One string field",
			struct {
				Thing string
			}{"potato"},
			[]string{"potato"},
		},
		{
			"Two string field",
			struct {
				Thing      string
				OtherThing string
			}{"potato", "cat"},
			[]string{"potato", "cat"},
		},
		{
			"One string, one int",
			struct {
				Thing string
				Num   int
			}{"potato", 32},
			[]string{"potato"},
		},
		{
			"One string, one int",
			nestedThing{"potato", innerThing{"cat", 23}},
			[]string{"potato", "cat"},
		},
		{
			"Pointers",
			&struct {
				Thing string
				Num   int
			}{"potato", 32},
			[]string{"potato"},
		},
		{
			"Slices",
			[]struct {
				Thing string
				Num   int
			}{
				{"potato", 32},
				{"cat", 73},
			},
			[]string{"potato", "cat"},
		},
		{
			"Arrays",
			[2]simpleThing{
				{"potato", 32},
				{"cat", 73},
			},
			[]string{"potato", "cat"},
		},
		{
			"Maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var got []string

			Walk(test.input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Got %q, want %q", got, test.want)
			}
		})
	}
}
