package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

// Runs the main package.
type File struct {
	Title string
	Id    string
}

type Files []File

func (files *Files) add(name string, id string) {
	file := File{
		Title: name,
		Id:    id,
	}
	*files = append(*files, file)
}

func (files *Files) validateIndex(index int) error {
	fmt.Println("validate index")
	if index < 0 || index > len(*files) {
		err := errors.New("Error: Invalid index")
		log.Fatalln(err)
	}
	return nil
}

func (files *Files) delete(num int) error {
	idx := num - 1
	f := *files

	if err := f.validateIndex(idx); err != nil {
		return err
	}
	// TODO There must be a better way to update files.
	*files = append(f[:idx], f[:idx+1]...)

	return nil
}

func (files *Files) edit(num int, title string, id string) error {
	idx := num - 1
	f := *files

	if err := f.validateIndex(idx); err != nil {
		return err
	}

	if len(title) > 0 {
		f[idx].Title = title
	}

	if len(id) > 0 {
		f[idx].Id = id
	}

	return nil
}

func (files *Files) list() {
	fmt.Print("\n\n")

	w := tabwriter.NewWriter(os.Stdout, 1, 5, 2, ' ', 0)

	fmt.Fprintln(w, "#\tTITLE\tID")

	for index, file := range *files {
		num := index + 1
		fmt.Fprintf(w, "%d\t%s\t%s\n", num, file.Title, file.Id)
	}
	w.Flush()

	fmt.Print("\n\n")
}
