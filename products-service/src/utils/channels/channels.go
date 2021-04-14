package channels

// OK takes in a channel that returns bool as arg and returns true once channel returns true, otherwise return false
func OK(done <-chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return true
		}
	}
	return false
}
