package core

type ElementFactory interface {
	Create(core RepositoryCore)
}

type ElementCreateHandler interface {
	OnCreate() error
}

type ElementDestroyHandler interface {
	OnDestroy() error
}
