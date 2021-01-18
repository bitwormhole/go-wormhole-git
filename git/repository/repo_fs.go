package repository

import (
	"github.com/bitwormhole/go-wormhole-core/io/fs"
)

// RepoDir  表示git仓库里的一个目录
type RepoDir interface {
	GetPath() fs.Path
}

// RepoFile 表示git仓库里的一个文件
type RepoFile interface {
	GetPath() fs.Path
}
