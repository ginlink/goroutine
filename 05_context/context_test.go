package my_context

import "testing"

func TestContextCancel(*testing.T) {
	ContextCancel()
}

func TestContextChainedCancel(*testing.T) {
	ContextChainedCancel()
}
