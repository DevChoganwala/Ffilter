/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var pf bool

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter file extensions",
	Long: `Filters file extensions
Default: Counts PDFs, images and text file on the working directory
Flags: --im=<image-format> -p`,
	Run: func(cmd *cobra.Command, args []string) {
		imgf, _ := cmd.Flags().GetString("im")

		if imgf != "" {
			getimages(checkargs(args), imgf)
		} else if pf {
			getpdfs(checkargs(args))
		} else {
			preview(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.PersistentFlags().String("im", "", "displays only images")
	filterCmd.PersistentFlags().BoolVarP(&pf, "pdf", "p", false, "checks for pdf")
}

func checkargs(args []string) (files []fs.FileInfo) {
	var path string
	var err error
	if len(args) == 0 {
		path, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory")
			return
		}
	} else if len(args) == 1 {
		path = args[0]
	} else {
		fmt.Println("Too many arguments (Try 0 or 1)")
		return
	}
	files, err = ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading current directory")
		return
	}
	return files
}

func preview(args []string) {
	images := []string{".png", ".jpeg", ".jpg"}
	text := []string{".txt", ".doc", ".docx"}
	pdf := []string{".pdf"}
	var files = checkargs(args)
	image, txt, pdfs := 0, 0, 0
	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if extinslice(ext, images) {
			image++
		} else if extinslice(ext, text) {
			txt++
		} else if extinslice(ext, pdf) {
			pdfs++
		}
	}
	fmt.Printf("Files Detected: \n 1. Images: %d\n 2. Text files: %d\n 3. PDFs: %d\n", image, txt, pdfs)
}

func extinslice(ext string, arr []string) (res bool) {
	for _, ex := range arr {
		if ex == ext {
			return true
		}
	}
	return false
}

func getimages(files []fs.FileInfo, arg string) {
	arg = "." + arg
	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if ext == arg {
			fmt.Println(f.Name())
		}
	}
}

func getpdfs(files []fs.FileInfo) {
	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if ext == ".pdf" {
			fmt.Println(f.Name())
		}
	}
}
