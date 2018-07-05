// authors: wangoo
// created: 2018-07-04
// test gerror

package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/soundbus-technologies/gerror"
	"fmt"
)

func TestGerror(t *testing.T) {
	assert.Equal(t, "E0000", gerror.ErrInternalError.Code())

	err := gerror.NewCodeError("E1234", "test error")

	assert.Equal(t, "E1234", err.Code())
	assert.Equal(t, "test error", err.Error())

	assert.Equal(t, "test error", fmt.Sprint(err))
	assert.Equal(t, "{\"code\":\"E1234\",\"error\":\"test error\"}", err.Json())

	checkError(t, err)
}

func checkError(t *testing.T, e interface{}) {
	_, ok := e.(error)
	assert.True(t, ok)
}
