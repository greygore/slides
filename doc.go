package main

func asdf() {
	var stopCh chan int
	var t struct {
		C chan int
	}
	var sawTimeout bool

	// START COMPLEX LOGIC OMIT
	// NOTE: b/c there is no priority selection in golang
	// it is possible for this to race, meaning we could
	// trigger t.C and stopCh, and t.C select falls through.
	// In order to mitigate we re-check stopCh at the beginning
	// of every loop to prevent extra executions of f().
	select {
	case <-stopCh:
		return
	case <-t.C:
		sawTimeout = true
	}
	// END COMPLEX LOGIC OMIT

	_ = sawTimeout
}
