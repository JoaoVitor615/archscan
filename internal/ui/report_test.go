package ui

import (
	"testing"
)

func TestRenderDashboard(t *testing.T) {
	// Must execute without panic
	RenderDashboard(".", 3, 6, 12)
}
