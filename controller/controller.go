package controller

// Controller is an type alias for controller
type Controller int

const (
	// Triangle is the first button
	Triangle Controller = iota + 1

	// Circle is the second button
	Circle

	// Cross is the third button
	Cross

	// Square is the fourth button
	Square
)

func (c Controller) String() string {
	controllers := [...]string{"Unknown", "Triangle", "Circle", "Square"}
	return controllers[c]
}
