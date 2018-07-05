package senml

var (
	True  = Bool(true)
	False = Bool(false)
)

// Bool returns a pointer to a bool value, to be used for Record.BoolValue.
func Bool(b bool) *bool {
	return &b
}

// Float returns a pointer to a float64 value, to be used for Record.Value and Record.Sum.
func Float(f float64) *float64 {
	return &f
}
