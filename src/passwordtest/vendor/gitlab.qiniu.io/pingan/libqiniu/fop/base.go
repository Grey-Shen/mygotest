package fop

import "gitlab.qiniu.io/pingan/libqiniu/op"

type saveasCmd struct {
	Entry op.Entry `fop:"saveas"`
}

type FopCommand string

func NewFopCommand(commandValue interface{}) (FopCommand, error) {
	if fop, ok := commandValue.(string); ok {
		return FopCommand(fop), nil
	}
	converter, err := convertToFopCommand(commandValue)
	if err != nil {
		return FopCommand(""), err
	} else {
		return FopCommand(converter.String()), nil
	}
}

func (cmd FopCommand) ToCommand() FopCommand {
	return cmd
}

func (cmd FopCommand) ToPipeline() FopPipeline {
	return NewFopPipeline([]FopCommand{cmd})
}

func (cmd FopCommand) ToMultiCommands() MultiFopCommands {
	return MultiFopCommands(MultiFopCommands{cmd.ToPipeline()})
}

func (cmd FopCommand) String() string {
	return string(cmd)
}

type FopPipeline []FopCommand

func NewFopPipeline(commands []FopCommand) FopPipeline {
	return FopPipeline(commands)
}

func (pipeline FopPipeline) ToPipeline() FopPipeline {
	return pipeline
}

func (pipeline FopPipeline) SaveAs(bucket, key string) FopPipeline {
	fopCmd, err := NewFopCommand(saveasCmd{Entry: op.Entry{Bucket: bucket, Key: key}})
	if err != nil {
		panic(err)
	}
	pipeline = append(pipeline, fopCmd)
	return pipeline
}

func (pipeline FopPipeline) ToMultiCommands() MultiFopCommands {
	return MultiFopCommands(MultiFopCommands{pipeline})
}

func (pipeline FopPipeline) String() string {
	fop := ""

	for i, cmd := range pipeline {
		if i > 0 {
			fop += "|"
		}
		fop += cmd.String()
	}
	return fop
}

type MultiFopCommands []FopPipeline

func NewMultiFopCommands(pipelineCommands []FopPipeline) MultiFopCommands {
	return MultiFopCommands(pipelineCommands)
}

func (multiFopCommands MultiFopCommands) ToMultiCommands() MultiFopCommands {
	return multiFopCommands
}

func (multiFopCommands MultiFopCommands) String() string {
	fop := ""

	for i, pipeline := range multiFopCommands {
		if i > 0 {
			fop += ";"
		}
		fop += pipeline.String()
	}

	return fop
}
