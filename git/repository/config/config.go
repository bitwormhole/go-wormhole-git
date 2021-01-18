package config

import "github.com/bitwormhole/go-wormhole-git/git/repository"

// GitConfig 接口表示仓库里的 config 文件
type GitConfig interface {
	repository.RepoFile
}
