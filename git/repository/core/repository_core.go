package core

import (
	"github.com/bitwormhole/go-wormhole-core/io/fs"
	"github.com/bitwormhole/go-wormhole-git/git/repository/config"
	"github.com/bitwormhole/go-wormhole-git/git/repository/head"
	"github.com/bitwormhole/go-wormhole-git/git/repository/index"
	"github.com/bitwormhole/go-wormhole-git/git/repository/objects"
	"github.com/bitwormhole/go-wormhole-git/git/repository/refs"
)

// RepositoryCore 表示一个仓库的对内面貌
type RepositoryCore struct {
	Path fs.Path // the core repository directory

	createHandlers  []ElementCreateHandler
	destroyHandlers []ElementDestroyHandler

	Objects objects.ObjectStore
	Config  config.GitConfig
	HEAD    head.GitHEAD
	Index   index.GitIndex
	Refs    refs.GitRefs
}

// Close 方法用于关闭一个 RepositoryCore，以便释放该对象持有的所有资源
func (inst *RepositoryCore) Close() error {

	errlist := make([]error, 0)
	list := inst.destroyHandlers

	for index := range list {
		item := list[index]
		err := item.OnDestroy()
		if err != nil {
			errlist = append(errlist, err)
		}
	}

	if len(errlist) > 0 {
		return errlist[0]
	}

	return nil
}
