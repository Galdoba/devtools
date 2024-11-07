package pathfinder

import (
	"fmt"
)

func NewPath(opts ...StdPathOption) (string, error) {
	path := ""
	settings := stdPathOpt{}
	for _, modify := range opts {
		modify(&settings)
	}

	if settings.isConfig && settings.isLog {
		return "", fmt.Errorf("both config and log options sellected")
	}
	return path, nil
}

type StdPathOption func(*stdPathOpt)

type stdPathOpt struct {
	home             string
	system           string
	program          string
	fileName         string
	layers           []string
	isConfig         bool
	isLog            bool
	envVarKey        string
	ensureExistiance bool
}

func WithHome(home string) StdPathOption {
	return func(spo *stdPathOpt) {
		spo.home = home
	}
}

func WithSystem(system string) StdPathOption {
	return func(spo *stdPathOpt) {
		spo.system = system
	}
}

func WithProgram(program string) StdPathOption {
	return func(spo *stdPathOpt) {
		spo.program = program
	}
}

func WithFileName(fileName string) StdPathOption {
	return func(spo *stdPathOpt) {
		spo.fileName = fileName
	}
}

func WithLayers(layers []string) StdPathOption {
	return func(spo *stdPathOpt) {
		spo.layers = layers
	}
}

func IsConfig() StdPathOption {
	return func(spo *stdPathOpt) {
		spo.isConfig = true
	}
}

func IsLog() StdPathOption {
	return func(spo *stdPathOpt) {
		spo.isLog = true
	}
}

func WithEnviromentVariable(key string) StdPathOption {
	return func(spo *stdPathOpt) {
		spo.envVarKey = key
	}
}

func EnsureExistiance() StdPathOption {
	return func(spo *stdPathOpt) {
		spo.ensureExistiance = true
	}
}
