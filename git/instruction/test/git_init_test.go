package test

import (
	"testing"

	"github.com/bitwormhole/go-wormhole-git/git/instruction"

	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

func prepareTestingDir(t *testing.T) fs.Path {
	dir := t.TempDir()
	fsys := fs.Default()
	return fsys.GetPath(dir + "/todo")
}

func TestGitInit(t *testing.T) {
	dir := prepareTestingDir(t)

	task := instruction.Default().Init()
	task.Path = dir
	task.Name = "repo_with_works_test"

	err := task.Runner.Run()
	if err != nil {
		t.Error(err)
	}
}

func TestGitInitBare(t *testing.T) {
	dir := prepareTestingDir(t)

	task := instruction.Default().Init()
	task.Path = dir
	task.Name = "bare_repo_test"
	task.Bare = true

	err := task.Runner.Run()
	if err != nil {
		t.Error(err)
	}
}
