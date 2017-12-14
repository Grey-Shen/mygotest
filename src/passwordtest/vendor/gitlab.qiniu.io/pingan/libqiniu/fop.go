package libqiniu

import "gitlab.qiniu.io/pingan/libqiniu/fop"

func CreateFopCommand(commandValue interface{}) (FopCommand, error) {
	return fop.NewFopCommand(commandValue)
}

func NewFopCommand(commandValue interface{}) FopCommand {
	cmd, err := CreateFopCommand(commandValue)
	if err != nil {
		panic(err)
	}
	return cmd
}

func NewFopPipeline(commands ...FopCommand) FopPipeline {
	cmds := make([]fop.FopCommand, len(commands))
	for i, cmd := range commands {
		cmds[i] = cmd.ToCommand()
	}
	return fop.NewFopPipeline(cmds)
}

func NewMultiFopCommands(pipelineCommands ...FopPipeline) MultiFopCommands {
	pipelines := make([]fop.FopPipeline, len(pipelineCommands))
	for i, cmd := range pipelineCommands {
		pipelines[i] = cmd.ToPipeline()
	}
	return fop.NewMultiFopCommands(pipelines)
}

type FopCommand interface {
	ToCommand() fop.FopCommand
	FopPipeline
}

type FopPipeline interface {
	ToPipeline() fop.FopPipeline
	MultiFopCommands
}

type MultiFopCommands interface {
	ToMultiCommands() fop.MultiFopCommands
}
