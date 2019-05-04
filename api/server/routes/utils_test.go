package routes

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	router.GET("/version", Version)

	code, body, err := MockRequest(http.MethodGet, "/version", nil)
	if err != nil {
		t.FailNow()
	}

	b, err := ioutil.ReadAll(body)
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "{\"data\":{\"version\":\"MarktHub APIs v0.0.1\"}}", string(b))
}
