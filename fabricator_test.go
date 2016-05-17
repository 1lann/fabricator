package fabricator

import (
	"math"
	"testing"
)

var unsorted = ResultSet{
	{1, 1},
	{3, 3},
	{5, 5},
	{2, 2},
	{4, 4},
	{6, 6},
}

func TestSorting(t *testing.T) {
	unsortedCopy := unsorted.Copy()
	unsortedCopy.Sort()

	for i := 0; i < 6; i++ {
		res := Result{float64(i + 1), float64(i + 1)}
		if unsortedCopy[i] != res {
			t.FailNow()
		}
	}
}

func TestAdding(t *testing.T) {
	set := NewResultSet()

	for i := 0; i < 6; i++ {
		set = set.Add(unsorted[i])
	}

	if len(set) != 6 {
		t.Log(set)
		t.Error("length of set is incorrect")
		t.FailNow()
	}

	for i := 0; i < 6; i++ {
		res := Result{float64(i + 1), float64(i + 1)}
		if set[i] != res {
			t.Error("unexpected value")
			t.FailNow()
		}
	}
}

func TestBasicXFabrication(t *testing.T) {
	set := ResultSet{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
	}

	unsortedCopy := unsorted.Copy()
	unsortedCopy.Sort()

	fb := unsortedCopy.Fabricate()

	if fb.Gradient() != 1 {
		t.Error("failed gradient")
		t.FailNow()
	}

	if fb.YIntercept() != 0 {
		t.Error("failed y-intercept")
		t.FailNow()
	}

	fb.Apply(false, set)

	for i := 0; i < 6; i++ {
		res := Result{float64(i + 1), float64(i + 1)}
		if set[i] != res {
			t.Error("failed value")
			t.FailNow()
		}
	}
}

func TestBasicYFabrication(t *testing.T) {
	set := ResultSet{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
		{0, 6},
	}

	unsortedCopy := unsorted.Copy()
	unsortedCopy.Sort()

	fb := unsortedCopy.Fabricate()

	if fb.Gradient() != 1 {
		t.Error("failed gradient")
		t.FailNow()
	}

	if fb.YIntercept() != 0 {
		t.Error("failed y-intercept")
		t.FailNow()
	}

	fb.Apply(true, set)

	for i := 0; i < 6; i++ {
		res := Result{float64(i + 1), float64(i + 1)}
		if set[i] != res {
			t.Error("failed value")
			t.FailNow()
		}
	}
}

func TestXFabrication(t *testing.T) {
	set := ResultSet{
		{0.5, 0},
		{1.2, 0},
		{1.6, 0},
		{1.8, 0},
		{3.4, 0},
		{5.9, 0},
	}

	unsortedCopy := unsorted.Copy()
	unsortedCopy.Sort()

	fb := unsortedCopy.Fabricate()

	if fb.Gradient() != 1 {
		t.Error("failed gradient")
		t.FailNow()
	}

	if fb.YIntercept() != 0 {
		t.Error("failed y-intercept")
		t.FailNow()
	}

	fb.Apply(false, set)

	for i := range set {
		if set[i].X != set[i].Y {
			t.Error("failed value")
			t.FailNow()
		}
	}
}

func TestYFabrication(t *testing.T) {
	set := ResultSet{
		{0, 0.5},
		{0, 1.2},
		{0, 1.6},
		{0, 1.8},
		{0, 3.4},
		{0, 5.9},
	}

	unsortedCopy := unsorted.Copy()
	unsortedCopy.Sort()

	fb := unsortedCopy.Fabricate()

	if fb.Gradient() != 1 {
		t.Error("failed gradient")
		t.FailNow()
	}

	if fb.YIntercept() != 0 {
		t.Error("failed y-intercept")
		t.FailNow()
	}

	fb.Apply(true, set)

	for i := range set {
		if set[i].X != set[i].Y {
			t.Error("failed value")
			t.FailNow()
		}
	}
}

func TestRandomFabrication(t *testing.T) {
	set := ResultSet{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
	}

	unsortedCopy := unsorted.Copy()
	unsortedCopy.Sort()

	fb := unsortedCopy.Fabricate()

	tries := 0
	for fb.Gradient() == 1 {
		tries++
		fb.GradientError(0.2)
		if tries > 10 {
			t.Error("failed to get random gradient (or extremely unlucky)")
			t.FailNow()
		}
	}
	if math.Abs(1-fb.Gradient()) > 0.2 {
		t.Error("failed gradient")
	}

	tries = 0
	for fb.YIntercept() == 0 {
		tries++
		fb.YInterceptError(0.2)
		if tries > 10 {
			t.Error("failed to get random y-intercept (or extremely unlucky)")
			t.FailNow()
		}
	}
	if math.Abs(fb.YIntercept()) > 0.2 {
		t.Error("failed y-intercept")
	}

	fb.Apply(false, set)

	for i := range set {
		t.Log("x:", set[i].X, "y:", set[i].Y)

		if set[i].X == set[i].Y {
			t.Error("same value for mixed errors (or just unlucky)")
			t.FailNow()
		}
	}
}

func TestErrorRangeFabrication(t *testing.T) {
	set := ResultSet{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
	}

	unsortedCopy := unsorted.Copy()
	unsortedCopy.Sort()

	fb := unsortedCopy.Fabricate()
	fb.SetErrorRange(0.3)
	fb.Apply(false, set)

	for i := range set {
		t.Log("x:", set[i].X, "y:", set[i].Y)

		if set[i].X != float64(i+1) {
			t.Error("X values have changed")
			t.FailNow()
		}

		if math.Abs(set[i].Y-float64(i+1)) > 0.3 {
			t.Error("fabricated values beyond error range")
			t.FailNow()
		}
	}
}

func TestRandomNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		r := randomRange(12.34, 89.34)
		if r < 12.34 || r > 89.34 {
			t.FailNow()
		}
	}
}

func TestSmallRandomNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		r := randomRange(15.5236, 15.5240)
		if r < 15.5236 || r > 15.5240 {
			t.FailNow()
		}
	}
}
