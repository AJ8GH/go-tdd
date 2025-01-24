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

	t.Run("Maps", func(t *testing.T) {
		var got []string
		input := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		Walk(input, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("Channels", func(t *testing.T) {
		channel := make(chan simpleThing)

		go func() {
			channel <- simpleThing{"potato", 32}
			channel <- simpleThing{"cat", 32}
			close(channel)
		}()

		var got []string
		want := []string{"potato", "cat"}

		Walk(channel, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, want %q", got, want)
		}
	})

	t.Run("Funcs", func(t *testing.T) {
		input := func() (simpleThing, simpleThing) {
			return simpleThing{"potato", 32}, simpleThing{"cat", 69}
		}

		var got []string
		want := []string{"potato", "cat"}

		Walk(input, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, want %q", got, want)
		}
	})
}

func assertContains(t *testing.T, got []string, s string) {
	t.Helper()
	for _, v := range got {
		if v == s {
			return
		}
	}
	t.Errorf("Expected %v to contain %q but it didn't", got, s)
}
