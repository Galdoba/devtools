package pathfinder

import (
	"fmt"
)

/*
app := "my_app"
pathfinder.NewPath(pathfinder.IsConfig(),pathfinder.WithProgram(app), pathfinder.WithFileName(app+".config"),)

*/

func NewPath(opts ...StdPathOption) (string, error) {
	path := ""
	settings := stdPathOpt{}
	settings.root = homeDir()
	for _, modify := range opts {
		modify(&settings)
	}
	if settings.isConfig && settings.isLog {
		return "", fmt.Errorf("both config and log options sellected")
	}
	path = settings.root
	switch {
	case settings.isConfig:
		path += ".config" + sep
	case settings.isLog:
		path += ".log" + sep
	default:
		path += "Programs" + sep
	}
	path += "galdoba" + sep
	if settings.system != "" {
		path += settings.system + sep
	}
	if settings.program != "" {
		path += settings.program + sep
	}
	for _, layer := range settings.layers {
		path += layer + sep
	}
	if settings.fileName != "" {
		path += settings.fileName
	}
	if err := pathValidation(path); err != nil {
		return "", err
	}

	return path, nil
}

type StdPathOption func(*stdPathOpt)

type stdPathOpt struct {
	root             string
	system           string
	program          string
	fileName         string
	layers           []string
	isConfig         bool
	isLog            bool
	envVarKey        string
	ensureExistiance bool
}

func WithRoot(root string) StdPathOption {
	return func(spo *stdPathOpt) {
		spo.root = root
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
