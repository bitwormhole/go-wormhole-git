package worktrees

import (
	"github.com/bitwormhole/go-wormhole-git/git/repository"
)

// WorkingDirectory 表示git工作区的基础目录
type WorkingDirectory interface {
	repository.RepoDir
}
