package main

var (
	version              = "v1.0.0"
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
