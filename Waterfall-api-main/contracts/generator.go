// +build ignore

package main

import (
	// "bytes"
	"fmt"
	// "go/format"
	// "io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

var (
	folders = []string{
		filepath.Clean("../abis"),
	}

	targetFolder = "../box"
)

// Define vars for build template
var convs = map[string]interface{}{"conv": fmtByteSlice}
var tmpl = template.Must(
	template.New("box-blob-generator").Funcs(convs).Parse(
		`package box

// Code generated by go generate; DO NOT EDIT.

func init() {
    {{- range $name, $file := . }}
        box.Add("{{ $name }}", []byte{ {{ conv $file }} })
    {{- end }}
}`))

func fmtByteSlice(s []byte) string {
	builder := strings.Builder{}
	for _, v := range s {
		builder.WriteString(fmt.Sprintf("%d,", int(v)))
	}
	return builder.String()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(len(folders))

	for _, folder := range folders {
		// Checking directory with files
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			log.Fatal(folder + " directory does not exists!")
		}

		// Create map for filenames
		// configs := make(map[string][]byte)
		// folderName := filepath.Base(folder)

		// Walking through embed directory
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(info.Name(), ".json") {
				if err := abigen(filepath.Clean(path)); err != nil {
					log.Printf("Error: %s", err)
					return err
				}
			}

			return nil
		})
		if err != nil {
			log.Fatal("Error walking through embed directory:", err)
		}

		// // Create blob file
		// blobFile := fmt.Sprintf("%s/%s_blob.go", targetFolder, folderName)
		// f, err := os.Create(blobFile)
		// if err != nil {
		// 	log.Fatal("Error creating blob file:", err)
		// }
		// if err := f.Close(); err != nil {
		// 	log.Fatal("Error creating blob file:", err)
		// }

		// // Create buffer
		// builder := &bytes.Buffer{}

		// // Execute template
		// if err = tmpl.Execute(builder, configs); err != nil {
		// 	log.Fatal("Error executing template", err)
		// }

		// // Formatting generated code
		// data, err := format.Source(builder.Bytes())
		// if err != nil {
		// 	log.Fatal("Error formatting generated code", err)
		// }

		// // Writing blob file
		// if err = ioutil.WriteFile(blobFile, data, os.ModePerm); err != nil {
		// 	log.Fatal("Error writing blob file", err)
		// }

		wg.Done()
	}
	wg.Wait()
}

func abigen(p string) error {
	f := filepath.Base(p)
	ext := filepath.Ext(f)
	name := strings.TrimSuffix(f, ext)
	abigenStr := fmt.Sprintf("abigen --abi %s --type %s --pkg contracts --lang go --out ./%s.go", p, name, name)
	cmd := exec.Command("sh", "-c", abigenStr)
	fmt.Println("abigen:", p)
	return cmd.Run()
}
