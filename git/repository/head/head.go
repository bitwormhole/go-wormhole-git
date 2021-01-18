package head

import "github.com/bitwormhole/go-wormhole-git/git/repository"

// GitHEAD 表示 git 的 HEAD 文件
type GitHEAD interface {
	repository.RepoFile
}
