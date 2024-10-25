package kernel

type KernelBuilder struct {
	Version    string
	SourceURL  string
	SourceDir  string
	OutputDir  string
	NumCores   int
	ConfigFile string
}

func NewKernelBuilder(version, sourceurl, sourcedir, outputdir string, numcores int) *KernelBuilder {
	return &KernelBuilder{
		Version:   version,
		SourceURL: sourceurl,
		SourceDir: sourcedir,
		OutputDir: outputdir,
		NumCores:  numcores,
	}
}
