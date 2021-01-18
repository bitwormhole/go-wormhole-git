package index

import "github.com/bitwormhole/go-wormhole-git/git/repository"

// GitIndex 表示git仓库里面的index文件
type GitIndex interface {
	repository.RepoFile
}
