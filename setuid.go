package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Path bool `short:"p" long:"path" description:"search directories in the $PATH variable first. Good for speed"`
	Find bool `short:"f" long:"find" description:"treat trailing args as paths to walk and search for setuid binaries. "`
}

func Walk(path string, fi fs.DirEntry, err error) error {
	if err != nil {
		log.Println(err)
	}

	if fi.IsDir() {
		return nil
	}

	f, _ := fi.Info()
	if f.Mode()&os.ModeSetuid != 0 {
		fmt.Printf("Found Setuid Binary: %v\n", f.Name())
	}

	return nil
}

func Find(root string) {
	filepath.WalkDir(root, Walk)
}

// check path variable first
func WookPath() ([]string, []string, error) {
	path := os.Getenv("PATH")
	if path == "" {
		return nil, nil, fmt.Errorf("Path not set")
	}

	var dirs []string
	var exes []string
	for _, dir := range filepath.SplitList(path) {
		if dir == "" {
			// Unix shell semantics: path element "" means "."
			dir = "."
		}

		dirs = append(dirs, dir)
		dirent, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
		}

		for x := range dirent {
			exe := dirent[x].Name()
			exe = filepath.Join(dir, exe)
			// fmt.Println(exe)
			exes = append(exes, exe)
		}

	}

	return dirs, exes, nil
}

func checkExecutable(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSetuid != 0
}

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}
	}

	// figure out how to force usage message lol
	// if !opts.Path && !opts.Find && len(args) == 0 {
	// 	fmt.Println("usage: findsuid --find / /home/")
	// 	fmt.Println("\tfindsuid --help for more")
	// 	os.Exit(0)
	// }

	if !opts.Path && !opts.Find || opts.Path {
		_, exes, err := WookPath()
		if err != nil {
			log.Println(err)
		}

		for _, exe := range exes {
			if checkExecutable(exe) {
				fmt.Printf("Found Setuid Binary: %v\n", exe)
			}
		}
	}

	if opts.Find {
		if len(args) < 1 {
			args = append(args, "/")
		}
		for _, dir := range args {
			Find(dir)
		}
	}
}
