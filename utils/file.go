package utils

import (
	"io/ioutil"
	"os"
)

// DeleteDirectory will attempt to delete a directory and all contents inside
//		DeleteDirectory("assets")
func DeleteDirectory(directory string) error {
	Log.Debugln("removing directory: " + directory)
	return os.RemoveAll(directory)
}

// CreateDirectory will attempt to create a directory
//		CreateDirectory("assets")
func CreateDirectory(directory string) error {
	Log.Debugln("creating directory: " + directory)
	if err := os.Mkdir(directory, os.ModePerm); err != os.ErrExist {
		return err
	}
	return nil
}

// FolderExists will return true if the folder exists
func FolderExists(folder string) bool {
	if _, err := os.Stat(folder); os.IsExist(err) {
		return true
	}
	return false
}

// FileExists returns true if a file exists
//		exists := FileExists("assets/css/base.css")
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		Log.Debugf("file exist: %v (%v)", name, !os.IsNotExist(err))
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// DeleteFile will attempt to delete a file
//		DeleteFile("newfile.json")
func DeleteFile(file string) error {
	Log.Debugln("deleting file: " + file)
	return os.Remove(file)
}

// RenameDirectory will attempt rename a directory to a new name
func RenameDirectory(fromDir string, toDir string) error {
	Log.Debugln("renaming directory: " + fromDir + "to: " + toDir)
	return os.Rename(fromDir, toDir)
}

// SaveFile will create a new file with data inside it
//		SaveFile("newfile.json", []byte('{"data": "success"}')
func SaveFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, os.ModePerm)
	return err
}

func OpenFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	return string(data), err
}
