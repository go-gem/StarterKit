package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	ProjectTypeBasic    = "basic"
	ProjectTypeAdvanced = "advanced"
)

var (
	goPath  = os.Getenv("GOPATH")
	srcPath = path.Join(goPath, "src")

	sourceCodePath = path.Join(srcPath, "github.com", "go-gem", "StarterKit", "src")

	projectName string
	projectType string
	projectRoot string

	completed bool
)

func main() {
	flag.StringVar(&projectName, "name", "", "project name which relative to $GOPATH/src")
	flag.StringVar(&projectType, "type", ProjectTypeBasic, `"basic" and "advanced" are valid.`)
	flag.Parse()

	projectName = strings.TrimSpace(projectName)
	if projectName == "" {
		log.Fatalln("the project name should not be blank")
		return
	}

	projectType = strings.TrimSpace(projectType)

	if projectType != ProjectTypeBasic && projectType != ProjectTypeAdvanced {
		log.Printf("invalid project type: %s, \"basic\" and \"advanced\" are valid", projectType)
		return
	}

	createProject()
}

func createProject() {
	log.Println("creating project...")
	projectRoot = path.Join(srcPath, projectName)

	if _, err := os.Stat(projectRoot); err == nil {
		log.Printf("the project already exists: %s\n", projectRoot)
		return
	}

	// rollback
	defer func() {
		if !completed {
			log.Println("something wrong, rollback.")
			rollback()
		}
	}()

	// copy source code
	if err := copyFolder(path.Join(sourceCodePath, projectType), projectRoot); err != nil {
		log.Printf("faild to copy source code: %s\n", err)
		return
	}

	// set process as completed
	complete()
}

func copyFolder(src string, dst string) (err error) {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	dir, _ := os.Open(src)
	objects, err := dir.Readdir(-1)
	for _, obj := range objects {
		newSrc := path.Join(src, obj.Name())
		newDst := path.Join(dst, obj.Name())
		if obj.IsDir() {
			err = copyFolder(newSrc, newDst)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copyFile(newSrc, newDst)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copyFile(src string, dst string) (err error) {
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	// fixed package import path
	if filepath.Base(src) == "app.go" {
		old := []byte(`github.com/go-gem/StarterKit/src/` + projectType)
		new := []byte(projectName)
		data = bytes.Replace(data, old, new, -1)
	}

	if err = ioutil.WriteFile(dst, data, fi.Mode()); err != nil {
		return err
	}

	return
}

func rollback() {
	if err := os.RemoveAll(projectRoot); err != nil {
		log.Printf("faild to remove project directory: %s\n", err)
	}
}

func complete() {
	completed = true
	log.Println("Congratulation! Your application has been set up.")
	log.Printf("Your project root: %s\n", projectRoot)
	log.Println("Please run the following command to launch your application:")
	printGuide()
	log.Println("and then visit http://localhost:8080")
}

func printGuide() {
	if projectType == ProjectTypeAdvanced {
		log.Printf("\tcd %s\n", path.Join(projectRoot, "frontend"))

		log.Println("1. Firstly, install static resource dependencies via yarn/npm:")
		log.Println("\tyarn install")
		log.Println("\tyarn build\n")

		log.Println("2. Next, import the database structure:")
		log.Println("\tmysql -uroot -p")
		log.Println("\t> CREATE SCHEMA `gem` DEFAULT CHARACTER SET utf8mb4;")
		log.Println("\t> use gem;")
		log.Printf("\t> source %s;\n", path.Join(projectRoot, "common", "data", "mysql.sql"))
		log.Println("\t> exit;\n")

		log.Println("3. And then, install go's package dependencies:")
		log.Println("\tgo get ./...\n")

		log.Println("4. Finally, build and run your application.")
		log.Println("\tgo build")
		log.Println("\tchmod +x ./frontend")
		log.Println("\t./frontend -c ./publish/app.json")
		return
	}

	log.Printf("\tcd %s", projectRoot)
	log.Println("\tgo build")
	scriptName := filepath.Base(projectName)
	log.Printf("\tchmod +x ./%s", scriptName)
	log.Printf("\t./%s -c ./publish/app.json", scriptName)
}
