package gox_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uthng/jobflow"
	"github.com/uthng/jobflow/plugins/gox"
)

func TestCmdBuild(t *testing.T) {
	testCases := []struct {
		name   string
		params map[string]interface{}
		result *jobflow.CmdResult
	}{
		{
			"OSArchMissing",
			map[string]interface{}{
				"output": "output1",
			},
			&jobflow.CmdResult{
				Error:  fmt.Errorf("param osarch missing"),
				Result: map[string]interface{}{},
			},
		},
		{
			"OutputMissing",
			map[string]interface{}{
				"osarch": "osarch1",
			},
			&jobflow.CmdResult{
				Error:  fmt.Errorf("param output missing"),
				Result: map[string]interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := gox.CmdBuild(tc.params)
			assert.Equal(t, result, tc.result)
		})
	}

}
