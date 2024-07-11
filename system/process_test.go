package system

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIsAlive(t *testing.T) {
	assert.Equal(t, true, IsAlive("/opt/visual-studio-code/code"))
	assert.Equal(t, false, IsAlive("/opt/visual-studio-code/code1"))
	assert.Equal(t, true, IsAlive(""))
	assert.Equal(t, true, IsAlive("sway"))
	assert.Equal(t, false, IsAlive("waybar"))
}
