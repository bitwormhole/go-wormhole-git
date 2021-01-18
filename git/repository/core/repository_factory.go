package core

import "github.com/bitwormhole/go-wormhole-core/io/fs"

// RepositoryCoreFactory 是用于创建 RepositoryCore 的工厂
type RepositoryCoreFactory struct {
	elements []ElementFactory
}

// Open 方法打开指定位置的仓库
func (inst *RepositoryCoreFactory) Open(path fs.Path) (*RepositoryCore, error) {
	core := &RepositoryCore{}
	core.Path = path
	core.CreateHandlers = make([]ElementCreateHandler, 0)
	core.DestroyHandlers = make([]ElementDestroyHandler, 0)

	err := inst.init1(core)
	if err != nil {
		return nil, err
	}

	err = inst.init2(core)
	if err != nil {
		return nil, err
	}

	return core, nil
}

func (inst *RepositoryCoreFactory) init1(core *RepositoryCore) error {
	list := inst.elements
	for index := range list {
		el := list[index]
		el.Create(core)
	}
	return nil
}

func (inst *RepositoryCoreFactory) init2(core *RepositoryCore) error {
	list := core.CreateHandlers
	for index := range list {
		handler := list[index]
		err := handler.OnCreate()
		if err == nil {
			continue
		} else {
			return err
		}
	}
	return nil
}
