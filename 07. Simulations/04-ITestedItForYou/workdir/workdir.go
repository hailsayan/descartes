package workdir

import (
	"errors"
	"github.com/otiai10/copy"
	"log"
	"os"
	"path/filepath"
)

type VCInterface interface {
	MarkFileAsModified(filePath string)
}

type WorkDir struct {
	RootPath string
	VC       VCInterface
}

func InitEmptyWorkDir() *WorkDir {
	tempDir, err := os.MkdirTemp("", "workdir")
	if err != nil {
		log.Fatalf("Failed to create a temporary directory: %v", err)
	}
	return &WorkDir{RootPath: tempDir}
}

func (wd *WorkDir) SetVC(vc VCInterface) {
	wd.VC = vc
}

func (wd *WorkDir) CreateFile(filePath string) error {
	fullPath := filepath.Join(wd.RootPath, filePath)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err.Error())
		}
	}(file)
	if wd.VC != nil {
		wd.VC.MarkFileAsModified(filePath)
	}
	return nil
}

func (wd *WorkDir) CreateDir(dirPath string) error {
	fullPath := filepath.Join(wd.RootPath, dirPath)
	err := os.MkdirAll(fullPath, 0755)
	if wd.VC != nil && err == nil {
		wd.VC.MarkFileAsModified(dirPath)
	}
	return err
}

func (wd *WorkDir) CatFile(filePath string) (string, error) {
	fullPath := filepath.Join(wd.RootPath, filePath)
	bytes, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (wd *WorkDir) WriteToFile(filePath, content string) error {
	fullPath := filepath.Join(wd.RootPath, filePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return errors.New("file does not exist")
	}

	dir := filepath.Dir(fullPath)

	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return errors.New("directory does not exist")
		}
		return err
	}

	err := os.WriteFile(fullPath, []byte(content), 0644)
	if wd.VC != nil && err == nil {
		wd.VC.MarkFileAsModified(filePath)
	}
	return err
}

func (wd *WorkDir) AppendToFile(filePath, content string) error {
	fullPath := filepath.Join(wd.RootPath, filePath)
	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err.Error())
		}
	}(f)
	_, err = f.WriteString(content)
	if wd.VC != nil && err == nil {
		wd.VC.MarkFileAsModified(filePath)
	}
	return err
}

func (wd *WorkDir) ListFilesRoot() []string {
	return wd.listFiles(wd.RootPath)
}

func (wd *WorkDir) ListFilesIn(dir string) ([]string, error) {
	path := filepath.Join(wd.RootPath, dir)
	return wd.listFiles(path), nil
}

func (wd *WorkDir) listFiles(path string) []string {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			relPath, err := filepath.Rel(wd.RootPath, path)
			if err == nil && !info.IsDir() {
				files = append(files, relPath)
			}
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return files
}

func (wd *WorkDir) Clone() *WorkDir {
	newPath := wd.RootPath + "_clone"
	if err := copy.Copy(wd.RootPath, newPath); err != nil {
		panic(err)
	}
	return &WorkDir{RootPath: newPath}
}
