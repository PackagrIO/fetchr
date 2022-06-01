package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestErrorsInterface(t *testing.T) {
	t.Parallel()

	//assert
	require.Implements(t, (*error)(nil), ConfigFileMissingError(fmt.Errorf("test")), "Should implement the error interface")
	require.Implements(t, (*error)(nil), ConfigValidationError(fmt.Errorf("test")), "Should implement the error interface")
	require.Implements(t, (*error)(nil), ProviderUnsupportedActionError(fmt.Errorf("test")), "Should implement the error interface")
	require.Implements(t, (*error)(nil), FileValidationError(fmt.Errorf("test")), "Should implement the error interface")
}

func Test_argsToError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		errName string
		args    []interface{}
		want    string
	}{
		{
			errName: "custom",
			args:    []interface{}{1, 2, 3},
			want:    "custom: 1, 2, 3",
		},
		{
			errName: "custom_with_string",
			args:    []interface{}{"hello", "world", 3},
			want:    "custom_with_string: hello, world, 3",
		},
		{
			errName: "custom_with_list",
			args:    []interface{}{[]string{"foo", "bar"}, "world", 3},
			want:    "custom_with_list: [foo bar], world, 3",
		},
		{
			errName: "custom_with_error",
			args:    []interface{}{[]string{"foo", "bar"}, "world", 3, errors.New("nested")},
			want:    "custom_with_error: [foo bar], world, 3, nested",
		},
	}

	//assert
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%v", tt.errName, tt.args)
		t.Run(testname, func(t *testing.T) {
			err := argsToError(tt.errName, tt.args...)
			assert.Equal(t, tt.want, err.Error())
		})
	}
}

func Test_Unwrap_Error(t *testing.T) {
	t.Parallel()
	var ErrorCritical = errors.New("critical error")

	err := argsToError("error", 1, "hello-world", ErrorCritical)
	assert.Equal(t, "error: 1, hello-world, critical error", err.Error())
	assert.True(t, errors.Is(err, ErrorCritical))
	//unwrap
	assert.Equal(t, "critical error", errors.Unwrap(err).Error())
}
