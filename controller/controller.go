package controller

// Controller is an type alias for controller
type Controller int

const (
	// Up is the first button
	Up Controller = iota + 1

	// Right is the second button
	Right

	// Left is the third button
	Left

	// Down is the fourth button
	Down
)

func (c Controller) String() string {
	controllers := [...]string{"Unknown", "Up", "Right", "Left", "Down"}
	return controllers[c]
}
