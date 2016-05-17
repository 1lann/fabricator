package fabricator

// LinearRegression performs linear regression on the ResultSet
// using least squares, and returns the gradient and y-intercept.
func (r ResultSet) LinearRegression() (float64, float64) {
	var sumX float64
	var sumY float64
	var sumXX float64
	var sumXY float64
	var count float64

	for _, res := range r {
		sumX += res.X
		sumY += res.Y
		sumXX += res.X * res.X
		sumXY += res.X * res.Y
		count++
	}

	gradient := (count*sumXY - sumX*sumY) / (count*sumXX - sumX*sumX)
	yIntercept := (sumY / count) - (gradient*sumX)/count
	return gradient, yIntercept
}
