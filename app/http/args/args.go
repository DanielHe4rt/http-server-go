package args

import (
	"os"
)

type RunnerArgs struct {
	Directory string
}

func GetArgs() RunnerArgs {

	runnerArgs := RunnerArgs{
		Directory: "",
	}

	args := os.Args
	for key, arg := range args {

		if arg == "--directory" {

			if len(args) <= key+1 {
				continue
			}

			pathArg := args[key+1]
			runnerArgs.Directory = pathArg
		}
	}
	return runnerArgs
}
