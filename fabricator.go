package fabricator

import (
	"sort"
)

// Result represents a single result, i.e. a data point.
type Result struct {
	X float64
	Y float64
}

// ResultSet represents a set of Result that is sorted in ascending order of X.
type ResultSet []Result

// Fabricator represents the fabricator generated from a set of results.
type Fabricator struct {
	yIntercept float64
	gradient   float64
	errorRange float64
}

// NewResultSet returns a new ResultSet.
func NewResultSet() ResultSet {
	return ResultSet{}
}

// Add adds a Result to a ResultSet in ascending order of X.
func (r ResultSet) Add(res Result) ResultSet {
	for i, val := range r {
		if val.X < res.X {
			continue
		}

		r = append(r[:i], append(ResultSet{res}, r[i:]...)...)
		return r
	}

	r = append(r, res)
	return r
}

// Copy returns a copy of the ResultSet.
func (r ResultSet) Copy() ResultSet {
	cpy := make(ResultSet, len(r))
	copy(cpy, r)
	return cpy
}

type byX ResultSet

func (a byX) Len() int           { return len(a) }
func (a byX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byX) Less(i, j int) bool { return a[i].X < a[j].X }

// Sort sorts the ResultSet in ascending order of X.
func (r ResultSet) Sort() {
	sort.Sort(byX(r))
}

// Fabricate returns a perfect (i.e. no errors) Fabricator based on the
// ResultSet.
func (r ResultSet) Fabricate() *Fabricator {
	fab := new(Fabricator)
	fab.gradient, fab.yIntercept = r.LinearRegression()
	return fab
}

// Apply applies the Fabricator on a given result set. Call this method
// multiple times to generate multiple trials. Set useY to true
// if you want to fill values of X using Y, instead of filling values of Y
// using X.
func (f *Fabricator) Apply(useY bool, res ResultSet) {
	if useY {
		gradient := 1.0 / f.gradient
		yIntercept := -f.yIntercept / f.gradient
		errorRange := f.errorRange / f.gradient

		for i := range res {
			val := res[i].Y*gradient + yIntercept
			res[i].X = randomRange(val-errorRange, val+errorRange)
		}

		return
	}

	for i := range res {
		val := res[i].X*f.gradient + f.yIntercept
		res[i].Y = randomRange(val-f.errorRange, val+f.errorRange)
	}
}

// YInterceptError sets the y-intercept of the
// Fabricator model to a random value in the given range from the original
// y-intercept.
func (f *Fabricator) YInterceptError(rangeVal float64) {
	f.yIntercept = randomRange(f.yIntercept-rangeVal, f.yIntercept+rangeVal)
}

// GradientError sets the slope of the Fabricator model to a random
// value in the given range from the original gradient.
func (f *Fabricator) GradientError(rangeVal float64) {
	f.gradient = randomRange(f.gradient-rangeVal, f.gradient+rangeVal)
}

// SetErrorRange sets the range of trials returned by the Fabricator using the
// given range.
func (f *Fabricator) SetErrorRange(rangeVal float64) {
	f.errorRange = rangeVal
}

// ErrorRange returns the error range of Y values on the Fabricator.
func (f *Fabricator) ErrorRange() float64 {
	return f.errorRange
}

// YIntercept returns the y-intercept of the Fabricator.
func (f *Fabricator) YIntercept() float64 {
	return f.yIntercept
}

// Gradient returns gradient of the Fabricator.
func (f *Fabricator) Gradient() float64 {
	return f.gradient
}
