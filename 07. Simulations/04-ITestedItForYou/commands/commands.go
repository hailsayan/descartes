package commands

import (
	"errors"
	"strconv"
	"strings"
	"vc/workdir"
)

type VC struct {
	WorkDir       *workdir.WorkDir
	Commits       []Commit
	StagedFiles   map[string]string
	ModifiedFiles map[string]bool
}

type Commit struct {
	Files   map[string]string
	Message string
}

type Status struct {
	ModifiedFiles []string
	StagedFiles   []string
}

func Init(wd *workdir.WorkDir) *VC {
	vc := &VC{
		WorkDir:       wd,
		Commits:       []Commit{},
		StagedFiles:   make(map[string]string),
		ModifiedFiles: make(map[string]bool),
	}
	wd.SetVC(vc) // Set the VC instance in WorkDir
	return vc
}

// MarkFileAsModified Implement the VCInterface
func (vc *VC) MarkFileAsModified(filePath string) {
	vc.ModifiedFiles[filePath] = true
}

func (vc *VC) GetWorkDir() *workdir.WorkDir {
	return vc.WorkDir
}

func (vc *VC) Add(files ...string) error {
	for _, file := range files {
		content, err := vc.WorkDir.CatFile(file)
		if err != nil {
			return err
		}
		vc.StagedFiles[file] = content
		delete(vc.ModifiedFiles, file)
	}
	return nil
}

func (vc *VC) AddAll() error {
	files := vc.WorkDir.ListFilesRoot()
	for _, file := range files {
		err := vc.Add(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vc *VC) Commit(message string) error {
	commit := Commit{
		Files:   make(map[string]string),
		Message: message,
	}
	for file, content := range vc.StagedFiles {
		commit.Files[file] = content
		delete(vc.ModifiedFiles, file)
	}
	vc.Commits = append(vc.Commits, commit)
	vc.StagedFiles = make(map[string]string)
	return nil
}

func (vc *VC) Status() Status {
	var modifiedFiles []string
	for file, modified := range vc.ModifiedFiles {
		if modified {
			modifiedFiles = append(modifiedFiles, file)
		}
	}
	return Status{
		ModifiedFiles: modifiedFiles,
		StagedFiles:   listKeys(vc.StagedFiles),
	}
}

func (vc *VC) Log() []string {
	logs := make([]string, 0)
	for i := len(vc.Commits) - 1; i >= 0; i-- {
		logs = append(logs, vc.Commits[i].Message)
	}
	return logs
}

func (vc *VC) Checkout(version string) (*workdir.WorkDir, error) {
	index, err := vc.parseVersion(version)
	if err != nil {
		return nil, err
	}
	if index < 0 || index >= len(vc.Commits) {
		return nil, errors.New("invalid commit index")
	}
	newWD := vc.WorkDir.Clone()
	for file, content := range vc.Commits[index].Files {
		err := newWD.WriteToFile(file, content)
		if err != nil {
			return nil, err
		}
	}
	return newWD, nil
}

func (vc *VC) parseVersion(version string) (int, error) {
	if strings.HasPrefix(version, "~") {
		offset, err := strconv.Atoi(version[1:])
		if err != nil {
			return 0, err
		}
		return len(vc.Commits) - 1 - offset, nil
	} else if strings.HasPrefix(version, "^") {
		count := len(version)
		if count > len(vc.Commits)-1 {
			return 0, errors.New("invalid commit index")
		}
		return len(vc.Commits) - 1 - count, nil
	}
	return strconv.Atoi(version)
}

func listKeys(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
