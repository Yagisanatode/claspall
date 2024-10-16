package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Command line flags
type CmdFlags struct {
	Add     string
	Edit    string
	Del     int
	List    bool
	Version bool
}

var FlagUsage = struct {
	Add     string
	Del     string
	Edit    string
	List    string
	Version string
}{
	Add: `
Add a new Apps Script file location:
- [how] Between quotation marks add a name or title then 
    a colon (:) and the file id. Close quotation marks.
- [syntax] claspall -add "title:Apps-Script-File-Id"
- [example] 
    claspall -add "Prod:1_hg5Lj-lOXbZMm60FizXSEZBmYN27-ozK-JOX4fRmEWntroxQ"
- [more]
-- You might consider having a "Test" project and a "Production" project AppsScirpt file
-- You can find the project ID in the 'Project Settings' > 'IDs' section.
-- Don't include the current project you are working in. This will be your 'Dev' file.
  `,
	Del: `
Delete an Apps Script file location:
- [how] Select an Apps Script file reference to delete by number from the -list.
- [syntax] claspall -del Number
- [example] 
    claspall -del 1
- [more]
-- You can use the -list flag to get the selected file location to remove
-- This will not delete the file. It will stop the deleted file from being updated from the core file.
    `,
	Edit: `
Edit existing Apps Script file location information:
- [how] Select an Apps Script file reference from the list and update the title and/or file id.
- [syntax] claspall -edit id:title:Apps-Script-File-Id
- [example] 
    claspall -edit "1:Prod:1_hg5Lj-lOXbZMm60FizXSEZBmYN27-ozK-JOX4fRmEWntroxQ
- [more] 
-- to edit just the Title: "1:New Title:"
--- Leave out the id after the first colon.
-- To edit just the Apps Script File ID: 
    "1::1_hg5Lj-lOXbZMm60FizXSEZBmYN27-ozK-JOX4fRmEWntroxQ"
--- Leave out the name between the ID and the Apps Script File ID colon separators.
    `,
	List: `
Lists all connected file locations
    `,
	Version: version,
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", FlagUsage.Add)
	flag.IntVar(&cf.Del, "del", -1, FlagUsage.Del)
	flag.StringVar(&cf.Edit, "edit", "", FlagUsage.Edit)
	flag.BoolVar(&cf.List, "list", false, FlagUsage.List)
	flag.BoolVar(&cf.List, "ls", false, "See -list")
	flag.BoolVar(&cf.Version, "version", false, FlagUsage.Version)
	flag.BoolVar(&cf.Version, "v", false, "See -version")

	flag.Parse()

	return &cf
}

func (cf CmdFlags) Execute(files *Files) {
	switch {
	case flag.NFlag() == 0:
		ClaspPush(files)

	case cf.List:
		files.list()

	case cf.Add != "":
		parts := strings.SplitN(cf.Add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("💥 Error, invalid format for Add.")
			fmt.Print(FlagUsage.Add)
			os.Exit(1)
		}

		files.add(parts[0], parts[1])

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 3)
		if len(parts) != 3 {
			fmt.Println("💥 Error, invalid format for Edit.")
			fmt.Print(FlagUsage.Edit)
			os.Exit(1)
		}

		line, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("💥 Error, invalid line number to edit")
			fmt.Print(FlagUsage.Edit)
			os.Exit(1)
		}

		files.edit(line, parts[1], parts[2])

	case cf.Del != -1:
		files.delete(cf.Del)

	case cf.Version:
		fmt.Println(version)
	}
}
