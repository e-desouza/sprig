package sprig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFail(t *testing.T) {
	const msg = "This is an error!"
	tpl := fmt.Sprintf(`{{fail "%s"}}`, msg)
	_, err := runRaw(tpl, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), msg)
}

func TestPass(t *testing.T) {
	const msg = "This is a useful debug message!"
	tpl := fmt.Sprintf(`{{pass "%s"}}`, msg)
	output, err := runRaw(tpl, nil)
	assert.NoError(t, err)
	assert.Contains(t, output, msg)
}
