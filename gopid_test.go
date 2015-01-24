package gopid

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPid(t *testing.T) {
	assert.Equal(t, 0, getPid("somepid.pid"))
}
