// vim: ts=4 et sts=4 sw=4
package config

import (
	//"bytes"
	//"fmt"
	//"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uthng/jobflow/job"
	// import all GoJobs's builtin modules
	_ "github.com/uthng/jobflow/plugins/all"
	//log "github.com/uthng/golog"
)

func TestReadFlowFile(t *testing.T) {
	var yamlFlowFile = []byte(`
variables:
  var1: $VAR1
  var2: ${VAR2}

build:
  tasks:
    - shell:
        cmd: exec
        params:
          cmd: echo 10
    - shell:
        cmd: exec
        params:
          cmd: echo 20

release:
  tasks:
    - name: "github release"
      github:
        cmd: release
        params:
          target: hello
`)

	//cmdFuncShellExec, _ := job.GetCmdByName("shell.exec")
	//cmdFuncGithubRelease, _ := job.GetCmdByName("github.release")

	flowOK := &JobFlow{
		Variables: map[string]interface{}{
			"var1": "$VAR1",
			"var2": "${VAR2}",
		},
		Jobs: []*job.Job{
			{
				Name: "build",
				Tasks: []*job.Task{
					{
						Name: "task-1",
						//Func: cmdFuncShellExec.Func,
						Params: map[string]interface{}{
							"cmd": "echo 10",
						},
					},
					{
						Name: "task-2",
						//Func: cmdFuncShellExec.Func,
						Params: map[string]interface{}{
							"cmd": "echo 20",
						},
					},
				},
				ValueRegistry: &job.ValueRegistry{
					ValueList: map[string]interface{}{
						"context": map[string]interface{}{
							"var1": "$VAR1",
							"var2": "${VAR2}",
						},
					},
				},
			},
			{
				Name: "release",
				Tasks: []*job.Task{
					{
						Name: "github release",
						//Func: cmdFuncGithubRelease.Func,
						Params: map[string]interface{}{
							"target": "hello",
						},
					},
				},
				ValueRegistry: &job.ValueRegistry{
					ValueList: map[string]interface{}{
						"context": map[string]interface{}{
							"var1": "$VAR1",
							"var2": "${VAR2}",
						},
					},
				},
			},
		},
	}

	jf := ReadFlowFile(yamlFlowFile)

	assert.Equal(t, flowOK.Variables, jf.Variables)

	for index, job := range jf.Jobs {
		assert.Equal(t, flowOK.Jobs[index].Name, job.Name)
		for idx, task := range job.Tasks {
			assert.Equal(t, flowOK.Jobs[index].Tasks[idx].Name, task.Name)

			// As we cannot compare 2 funcs in go, so we zap it
			assert.Equal(t, flowOK.Jobs[index].Tasks[idx].OnSuccess, task.OnSuccess)
			assert.Equal(t, flowOK.Jobs[index].Tasks[idx].OnFailure, task.OnFailure)
			assert.Equal(t, flowOK.Jobs[index].Tasks[idx].Result, task.Result)
		}
		assert.Equal(t, flowOK.Jobs[index].ValueRegistry, job.ValueRegistry)
	}
}