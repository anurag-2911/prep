package filemanager

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// basic file operation commands like list files, create files etc
func FileManage() {
	fileoperations()
}

func fileoperations() {
	f := &FileOps{}

	fmt.Println("file operations")
	if len(os.Args) < 2 {
		fmt.Println("usage: command arguments")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "list":
		{
			dir := "c"

			if len(os.Args) > 2 {
				fmt.Println("length of arguments ",len(os.Args))
				fmt.Println(os.Args)
				dir = os.Args[2]
			}
			fi, err := f.ListFiles(dir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, v := range fi {
				fmt.Println(v.Name())
			}
			os.Exit(1)
		}
	case "createfile":
		{
			dirname := os.Args[2]
			filename := os.Args[3]
			fmt.Println(f.CreateFile(filename, dirname))
			os.Exit(1)
		}
	case "deletefile":
		{
			filename := os.Args[2]
			fmt.Println(f.Delete(filename))
			os.Exit(1)
		}
	default:
		{
			fmt.Println("incorrect usage")
			os.Exit(1)
		}
	}
}

type FileOps struct {
}

func (f *FileOps) ListFiles(dirpath string) ([]fs.FileInfo, error) {
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		infos = append(infos, info)
	}

	return infos, err
}
func (f *FileOps) CreateFile(filename, dir string) error {

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error is creating directory %v", err)
	}
	fullfilename := filepath.Join(dir, filename)
	fl, err := os.Create(fullfilename)

	if err != nil {
		fmt.Println("error in creating file ", err)
	}
	defer fl.Close()

	return nil
}
func (f *FileOps) Delete(filename string) error {
	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("error is deleting the file %v", err)
	}
	return nil
}
