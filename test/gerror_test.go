// authors: wangoo
// created: 2018-07-04
// test gerror

package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/soundbus-technologies/gerror"
	"fmt"
	"net/http"
)

func TestGerror(t *testing.T) {
	assert.Equal(t, "E0000", gerror.ErrInternalError.Code())

	err := gerror.New("E1234", "test error")

	assert.Equal(t, "E1234", err.Code())
	assert.Equal(t, "test error", err.Msg())

	assert.Equal(t, "E1234:test error", fmt.Sprint(err))
	checkError(t, err)

	serviceErr := gerror.NewServiceCodeError(err, http.StatusBadRequest, "parameter required")
	assert.Equal(t, "E1234", serviceErr.ErrCode)
	assert.Equal(t, "test error", serviceErr.ErrMsg)

	serrString := fmt.Sprint(serviceErr)

	assert.Equal(t, "{\"status\":400,\"errCode\":\"E1234\",\"errMsg\":\"test error\",\"devMsg\":\"parameter required\"}", serrString)
	checkError(t, serviceErr)

}

func checkError(t *testing.T, e interface{}) {
	_, ok := e.(error)
	assert.True(t, ok)
}
