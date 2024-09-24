package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type claspConfig struct {
	ScriptId string `json:"scriptId"`
	RootDir  string `json:"rootDir"`
}

func ClaspPush(files *Files) {
	fmt.Print("\n\n *** CLASP ALL ***\n\n")
	claspExists()

	// Read the exiting JSON file.
	claspFile := getClaspJSON()

	var config claspConfig
	err := json.Unmarshal(claspFile, &config)
	if err != nil {
		fmt.Printf("💥 ERROR parsing '%s':'%s'\n", claspJsonFileName, err)
	}

	// Store the oringal scriptId
	original := config.ScriptId

	// Iterate over the clasp location arrays.
	if len(*files) == 0 {
		fmt.Printf("💥 ERROR no alternate Apps Script Files Listed.\n use the '-add' flag.\n")
		fmt.Print(FlagUsage.Add)
		os.Exit(1)
	}

	for _, f := range *files {

		fmt.Printf("🚀 Updating '%s'(%s)", f.Title, f.Id)

		// Update the script ID in the json file.
		config.ScriptId = f.Id
		writeToClaspJson(config)
		runCmdClaspPush()
	}

	// Return the original file id back and update.

	fmt.Print("🚀 Updating the main file")
	config.ScriptId = original
	writeToClaspJson(config)
	runCmdClaspPush()

	fmt.Print("\n*** CLASPALL complete ***\n")
}

func claspExists() {
	_, err := exec.LookPath(claspName)
	if err != nil {

		fmt.Println("Could not find the 'clasp' command.")
		fmt.Println("Go to: https://github.com/google/clasp to learn more.")
		os.Exit(1)
	}

	fmt.Print("✅ 'Clasp' is installed\n")
}

func getClaspJSON() []byte {
	file, err := os.ReadFile(claspJsonFileName)
	if err != nil {
		fmt.Printf("💥 ERROR parsing '%s': %s\n", claspJsonFileName, err)
		os.Exit(1)
	}
	fmt.Printf("✅  '%s' file exists in directory\n", claspJsonFileName)
	return file
}

func writeToClaspJson(config claspConfig) {
	// Marshal the data back to JSON
	updatedData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Printf("💥 ERROR marshalling JSON: '%s'\n", err)
		return
	}

	// Write the updated JSON data to the file
	err = os.WriteFile(claspJsonFileName, updatedData, 0644)
	if err != nil {
		fmt.Printf("💥 ERROR Could not update '%s' JSON with alternate IDs: %s", claspJsonFileName, err)
		return
	}
}

func runCmdClaspPush() {
	cmd := exec.Command(claspName, "push")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("💥 ERROR accessing termial: %s", err)
	}
}
