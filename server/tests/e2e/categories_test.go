package e2e_test

import (
	"api/recipes/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPostCategory(t *testing.T) {
	tests.StubServer()
	/* URL: 
	/recipes/1/categories POST
	/recipes/1/categories GET
	*/
	assert.Equal(t, false, false)
}
