package instruction

// GitInit 指令用来执行“git-init”命令
type GitInit struct {
	GitInstruction
	Bare bool
	Name string
}
