package objects

import (
	"github.com/bitwormhole/go-wormhole-git/git/repository"
)

// ObjectStore 接口表示本地对象存储
type ObjectStore interface {
	repository.RepoDir
}
