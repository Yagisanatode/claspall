package main

var (
	version              = "v0.1.1"
	claspName            = "clasp"
	claspJsonFileName    = ".clasp.json"
	claspallJsonFileName = ".claspall.json"
)

func main() {
	files := Files{}
	storage := NewStorage[Files](claspallJsonFileName)
	storage.Load(&files)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&files)

	storage.Save(files)
}
