package instruction

// GitInstructionSet 表示一个 git 指令集
type GitInstructionSet interface {
	Init() *GitInit
	Add() *GitAdd
	Commit() *GitCommit
	Push() *GitPush
	Status() *GitStatus
}

func Default() GitInstructionSet {
	return &defaultGitInstructionSet{}
}
