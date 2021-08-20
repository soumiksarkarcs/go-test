// go test -coverprofile cover.out
// go tool cover -func coverage.out
// go tool cover -html=cover.out opens cover.out in your default browser
// go tool cover -html=cover.out -o cover.html

package main

import "testing"

func TestAdd(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2.0, 4.0},
		{5.0, 2.0, 7},
	}

	for _, table := range tables {
		total := Add(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestSub(t *testing.T) {
	testData := []struct {
		par1     int
		par2     int
		expected int
	}{
		{3, 2, 1},
		{-3, 2, -5},
		{0, 2, -2},
	}

	for _, data := range testData {
		ans := Sub(data.par1, data.par2)
		if ans != data.expected {
			t.Errorf("Subtraction of (%d+%d) was incorrect, got: %d, want: %d.", data.par1, data.par2, ans, data.expected)
		}
	}

}

func TestMul(t *testing.T) {
	testData := []struct {
		par1     int
		par2     int
		expected int
	}{
		{3, 2, 6},
		{-3, 2, -6},
		{0, 2, 0},
	}

	for _, data := range testData {
		ans := Mul(data.par1, data.par2)
		if ans != data.expected {
			t.Errorf("Multiplication of (%d+%d) was incorrect, got: %d, want: %d.", data.par1, data.par2, ans, data.expected)
		}
	}

}

func TestDiv(t *testing.T) {
	testData := []struct {
		par1     int
		par2     int
		expected int
	}{
		{30, 2, 15},
		{-30, 5, -6},
		{0, 2, 0},
	}

	for _, data := range testData {
		ans := Div(data.par1, data.par2)
		if ans != data.expected {
			t.Errorf("Division of (%d+%d) was incorrect, got: %d, want: %d.", data.par1, data.par2, ans, data.expected)
		}
	}

}

func TestMain(t *testing.T) {
	main()
}
