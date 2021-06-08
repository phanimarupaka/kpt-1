// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmddiff_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleContainerTools/kpt/internal/cmddiff"
	"github.com/GoogleContainerTools/kpt/internal/cmdget"
	"github.com/GoogleContainerTools/kpt/internal/printer/fake"
	"github.com/GoogleContainerTools/kpt/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(testutil.ConfigureTestKptCache(m))
}

func TestCmdInvalidDiffType(t *testing.T) {
	runner := cmddiff.NewRunner(fake.CtxWithFakePrinter(nil, nil), "")
	runner.C.SetArgs([]string{"--diff-type", "invalid"})
	err := runner.C.Execute()
	assert.EqualError(t,
		err,
		"invalid diff-type 'invalid'. Supported diff-types are: local, remote, combined, 3way")
}

func TestCmdInvalidDiffTool(t *testing.T) {
	runner := cmddiff.NewRunner(fake.CtxWithFakePrinter(nil, nil), "")
	runner.C.SetArgs([]string{"--diff-tool", "nodiff"})
	err := runner.C.Execute()
	assert.EqualError(t,
		err,
		"diff-tool 'nodiff' not found in the PATH.")
}

func TestCmdExecute(t *testing.T) {
	g, w, clean := testutil.SetupRepoAndWorkspace(t, testutil.Content{
		Data:   testutil.Dataset1,
		Branch: "master",
	})
	defer clean()

	defer testutil.Chdir(t, w.WorkspaceDirectory)()

	dest := filepath.Join(w.WorkspaceDirectory, g.RepoName)

	getRunner := cmdget.NewRunner(fake.CtxWithFakePrinter(nil, nil), "")
	getRunner.Command.SetArgs([]string{"file://" + g.RepoDirectory + ".git/", "./"})
	err := getRunner.Command.Execute()
	assert.NoError(t, err)

	runner := cmddiff.NewRunner(fake.CtxWithFakePrinter(nil, nil), "")
	runner.C.SetArgs([]string{dest, "--diff-type", "local"})
	err = runner.C.Execute()
	assert.NoError(t, err)
}
