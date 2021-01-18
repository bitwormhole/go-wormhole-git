package instruction

type defaultGitInstructionSet struct {
}

func (inst *defaultGitInstructionSet) Add() *GitAdd {
	return nil
}

func (inst *defaultGitInstructionSet) Commit() *GitCommit {
	return nil
}

func (inst *defaultGitInstructionSet) Init() *GitInit {
	task := &GitInit{}
	task.Runner = &defaultGitInitRunner{task: task}
	return task
}

func (inst *defaultGitInstructionSet) Push() *GitPush {
	return nil
}

func (inst *defaultGitInstructionSet) Status() *GitStatus {
	return nil
}
