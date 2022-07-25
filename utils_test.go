package helpers_test

import (
	"strings"
	"testing"

	helpers "github.com/jaigouk/utils"
)

func TestContains(t *testing.T) {
	t.Run("Contains checks char element", func(t *testing.T) {
		elems := []string{"a", "b", "c"}
		if !helpers.Contains(elems, "b") {
			t.Error("Expected true, got false")
		}
	})

	t.Run("Contains checks int element", func(t *testing.T) {
		elems := []int{1, 2, 3}
		if !helpers.Contains(elems, 2) {
			t.Error("Expected true, got false")
		}
	})

	t.Run("Contains returns false if there is no matching element", func(t *testing.T) {
		elems := []int{1, 2, 3}
		if helpers.Contains(elems, 10) {
			t.Error("Expected false, got true")
		}
	})
}

func TestKeys(t *testing.T) {
	t.Run("Keys returns keys of map", func(t *testing.T) {
		m := map[string]string{"a": "b", "c": "d"}
		keys := helpers.Keys(m)
		if len(keys) != 2 {
			t.Error("Expected 2, got", len(keys))
		}

		if !helpers.Contains(keys, "a") {
			t.Error("Expected true, got false")
		}

		if !helpers.Contains(keys, "c") {
			t.Error("Expected true, got false")
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("Filter returns filtered elements", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		f := func(s string) bool {
			return s == "b"
		}
		n := helpers.Filter(slice, f)
		if len(n) != 1 {
			t.Error("Expected 1, got", len(n))
		}

		if n[0] != "b" {
			t.Error("Expected b, got", n[0])
		}
	})

	t.Run("Filter can be used for string match", func(t *testing.T) {
		websites := []string{"http://foo.com", "https://bar.com", "https://gosamples.dev"}
		httpsWebsites := helpers.Filter(websites, func(v string) bool {
			return strings.HasPrefix(v, "https://")
		})
		if len(httpsWebsites) != 2 {
			t.Error("Expected 2, got", len(httpsWebsites))
		}
	})
}

func TestSortSlice(t *testing.T) {
	t.Run("SortSlice sorts slice", func(t *testing.T) {
		slice := []int{3, 2, 1}
		helpers.SortSlice(slice)
		if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
			t.Error("Expected 1, 2, 3, got", slice)
		}

		stringSlice := []string{"z", "a", "b"}
		helpers.SortSlice(stringSlice)
		if stringSlice[0] != "a" || stringSlice[1] != "b" || stringSlice[2] != "z" {
			t.Error("Expected a, b, z, got", stringSlice)
		}

		floatSlice := []float64{2.3, 1.2, 0.2, 51.2}
		helpers.SortSlice(floatSlice)
		if floatSlice[0] != 0.2 || floatSlice[1] != 1.2 || floatSlice[2] != 2.3 || floatSlice[3] != 51.2 {
			t.Error("Expected 0.2, 1.2, 2.3, 51.2, got", floatSlice)
		}
	})
}

func TestMapSlice(t *testing.T) {
	t.Run("MapSlice maps slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		mapped := helpers.MapSlice(slice, func(i int) int {
			return i * 2
		})

		if mapped[0] != 2 || mapped[1] != 4 || mapped[2] != 6 {
			t.Error("Expected 2, 4, 6, got", mapped)
		}

		words := []string{"a", "b", "c", "d"}
		quoted := helpers.MapSlice(words, func(s string) string {
			return "\"" + s + "\""
		})
		if quoted[0] != "\"a\"" || quoted[1] != "\"b\"" || quoted[2] != "\"c\"" || quoted[3] != "\"d\"" {
			t.Error("Expected \"a\", \"b\", \"c\", \"d\", got", quoted)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("Reduce reduces slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		sum := helpers.Reduce(slice, func(i int, j int) int {
			return i + j
		}, 0)
		if sum != 6 {
			t.Error("Expected 6, got", sum)
		}

		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		divided := helpers.Reduce(numbers, func(acc float64, current int) float64 {
			return acc + float64(current)/10.0
		}, 0)
		if divided != 5.5 {
			t.Error("Expected 3.1, got", divided)
		}
	})
}

type FruitRank struct {
	Fruit string
	Rank  uint64
}

func TestUniq(t *testing.T) {
	t.Run("Uniq returns unique elements", func(t *testing.T) {
		fruits := []FruitRank{
			{
				Fruit: "Strawberry",
				Rank:  1,
			},
			{
				Fruit: "Raspberry",
				Rank:  2,
			},
			{
				Fruit: "Blueberry",
				Rank:  3,
			},
			{
				Fruit: "Blueberry",
				Rank:  3,
			},
			{
				Fruit: "Strawberry",
				Rank:  1,
			},
		}
		uniq := helpers.Uniq(fruits)
		if len(uniq) != 3 {
			t.Error("Expected 3, got", len(uniq))
		}
	})
}
