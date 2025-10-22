package action

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionFromReader(t *testing.T) {
	// Test: Good GET with argument
	a, err := ActionFromReader(strings.NewReader("GET key\r\n"))
	require.NoError(t, err)
	require.NotNil(t, a)
	assert.Equal(t, GET, a.Type)
	assert.Equal(t, []string{"key"}, a.Args)

	// Test: Good PUT with multiple arguments
	a, err = ActionFromReader(strings.NewReader("PUT key value\r\n"))
	require.NoError(t, err)
	require.NotNil(t, a)
	assert.Equal(t, PUT, a.Type)
	assert.Equal(t, []string{"key", "value"}, a.Args)

	// Test: Good PUT with multiple arguments and \n on value
	a, err = ActionFromReader(strings.NewReader("PUT key long\npoem\r\n"))
	require.NoError(t, err)
	require.NotNil(t, a)
	assert.Equal(t, PUT, a.Type)
	assert.Equal(t, []string{"key", "long\npoem"}, a.Args)

	// Test: Good DELETE
	a, err = ActionFromReader(strings.NewReader("DELETE key\r\n"))
	require.NoError(t, err)
	require.NotNil(t, a)
	assert.Equal(t, DELETE, a.Type)
	assert.Equal(t, []string{"key"}, a.Args)
}
