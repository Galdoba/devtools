package v2

func WithRoot(root string) Option {
	return func(p *pathfinder) {
		p.root = root
	}
}

func WithAppName(appName string) Option {
	return func(p *pathfinder) {
		p.appName = appName
	}
}

func WithFileName(fileName string) Option {
	return func(p *pathfinder) {
		p.fileName = fileName
	}
}

func WithExtention(ext string) Option {
	return func(p *pathfinder) {
		p.ext = ext
		p.mustHaveExt = true
	}
}

func WithPrefixLayers(layers ...string) Option {
	return func(p *pathfinder) {
		p.prefixlayers = layers
	}
}

func WithSuffixLayers(layers ...string) Option {
	return func(p *pathfinder) {
		p.suffixlayers = layers
	}
}

var WithSignature = Option(
	func(p *pathfinder) {
		p.mustHaveSignature = true
	})

var NoSignature = Option(
	func(p *pathfinder) {
		p.mustHaveSignature = false
	})

var DirMustExist = Option(
	func(p *pathfinder) {
		p.mustExistDir = true
	})

var FileMustExist = Option(
	func(p *pathfinder) {
		p.mustExistFile = true
	})

var NoValidation = Option(
	func(p *pathfinder) {
		p.skipValidation = true
	})
