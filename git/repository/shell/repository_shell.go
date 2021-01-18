package shell

import (
	"github.com/bitwormhole/go-wormhole-git/git/repository"
	"github.com/bitwormhole/go-wormhole-git/git/repository/core"
	"github.com/bitwormhole/go-wormhole-git/git/repository/objects"
	"github.com/bitwormhole/go-wormhole-git/git/repository/refs"
	"github.com/bitwormhole/go-wormhole-git/git/repository/worktrees"
)

// RepositoryShell 接口表示一个仓库的对外面貌
type RepositoryShell interface {
	GetCore() core.RepositoryCore

	GetWorkingDirectory() worktrees.WorkingDirectory
	GetObjects() objects.ObjectStore
	GetConfig() repository.GitConfig
	GetHEAD() repository.GitHEAD
	GetIndex() repository.GitIndex
	GetRefs() refs.GitRefsManager
}
