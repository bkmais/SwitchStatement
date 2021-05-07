package main

func shortswitch() string {

	//Short switch statement declares a variable and check "true"
	// there is a 'true' after the ; but we dont need to
	// show it in Go as it is true by default
	switch i := 10; {
	case i > 0:
		return "Positive"
	case i < 0:
		return "Negative"

	default:
		return "Nothing."
	}
}
