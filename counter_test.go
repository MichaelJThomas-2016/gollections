package gollections

import (
	"reflect"
	"testing"
)

func TestCounter_MostCommon(t *testing.T) {
	c := NewCounter[string]()
	words := []string{"a", "b", "a", "c", "b", "a"}
	for _, w := range words {
		c.Add(w)
	}

	t.Run("is most common", func(t *testing.T) {
		expected := make([]CountItem[string], 0, 2)
		expected = append(expected, CountItem[string]{"a", 3})
		expected = append(expected, CountItem[string]{"b", 2})
		actual := c.MostCommon(2)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Got %v\n %v\n", actual, expected)
		}
	},
	)

}

func TestCounter_Total(t *testing.T) {
	c := NewCounter[string]()
	words := []string{"a", "b", "a", "c", "b", "a"}
	for _, w := range words {
		c.Add(w)
	}
	expected := 6
	actual := c.Total()
	t.Run("test totals", func(t *testing.T) {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Got %v\n %v\n", actual, expected)
		}
	})
}
