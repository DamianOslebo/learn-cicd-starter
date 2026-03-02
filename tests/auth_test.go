package tests

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizationAPIKey(t *testing.T) {
	var header http.Header
	// Test empty input
	_, err := auth.GetAPIKey(header)
	assert.ErrorIs(t, err, auth.ErrNoAuthHeaderIncluded)

	req, err := http.NewRequest("", "", nil)
	header = req.Header

	// Test no apikey header
	header.Add("Authorization", "ApiKey")
	_, err = auth.GetAPIKey(header)
	assert.Error(t, err)

	// Test valid header
	header.Set("Authorization", "ApiKey zkldjasdfonaldnfajhgv")
	key, err := auth.GetAPIKey(header)
	assert.NoError(t, err)
	assert.Equal(t, key, "zkldjasdfonaldnfajhgv")
}
