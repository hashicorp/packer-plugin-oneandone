// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oneandone

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

func TestBuilderAcc_basic(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name:     "oneandone_basic_test",
		Template: testBuilderAccBasic,
		Setup: func() error {
			if v := os.Getenv("ONEANDONE_TOKEN"); v == "" {
				t.Fatal("ONEANDONE_TOKEN must be set for acceptance tests")
			}
			return nil
		},
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			if buildCommand.ProcessState != nil {
				if buildCommand.ProcessState.ExitCode() != 0 {
					return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
				}
			}
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}

const testBuilderAccBasic = `
{
      "builders": [{
	      "type": "oneandone",
	      "disk_size": "50",
	      "snapshot_name": "test5",
	      "image" : "ubuntu1604-64min"
    }]
}
`
