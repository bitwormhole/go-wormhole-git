package instruction

import "github.com/bitwormhole/go-wormhole-core/io/fs"

// GitInstruction 表示一条git指令
type GitInstruction struct {
	Path   fs.Path
	Runner GitInstructionRunner
}

// GitInstructionRunner 表示一条git指令
type GitInstructionRunner interface {
	Run() error
}
