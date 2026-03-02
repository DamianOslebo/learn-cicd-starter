package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizationAPIKey(t *testing.T) {
	var header http.Header
	// Test empty input
	_, err := GetAPIKey(header)
	assert.ErrorIs(t, err, ErrNoAuthHeaderIncluded)

	req, err := http.NewRequest("", "", nil)
	header = req.Header

	// Test no apikey header
	header.Add("Authorization", "ApiKey")
	_, err = GetAPIKey(header)
	assert.Error(t, err)

	// Test valid header
	header.Set("Authorization", "ApiKey zkldjasdfonaldnfajhgv")
	key, err := GetAPIKey(header)
	assert.NoError(t, err)
	assert.Equal(t, key, "zkldjasdfonaldnfajhgv")
}
