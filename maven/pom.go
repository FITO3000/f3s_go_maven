package maven

import (
	"fmt"
	"os"
	"path/filepath"
)

type Pom interface {
	Model
	Store() error
	Load() error
}

type Model interface {
}

func PomExists(directory string) (bool, error) {
	return fileExists(filepath.Join(directory, "pom.xml"))
}

func DeletePom(directory string) error {
	pomPath := filepath.Join(directory, "pom.xml")
	return os.Remove(pomPath)
}

func CreatePom(directory string) (Pom, error) {
	pomPath := filepath.Join(directory, "pom.xml")
	if exists, err := fileExists(pomPath); err != nil {
		return nil, err
	} else {
		if exists {
			return nil, fmt.Errorf("%s already exists", pomPath)
		} else {
			if err := createFile(pomPath); err != nil {
				return nil, err
			} else {
				pom := pom{
					path:  pomPath,
					model: NewProject(),
				}
				if err := pom.Store(); err != nil {
					return nil, err
				} else {
					return &pom, nil
				}
			}
		}
	}
}

func LoadPom(directory string) (Pom, error) {
	if exists, err := PomExists(directory); err != nil {
		return nil, err
	} else {
		if !exists {
			return nil, fmt.Errorf("there is no pom.xml file in directory: %s", directory)
		} else {
			pom := &pom{
				path: filepath.Join(directory, "pom.xml"),
			}
			if err := pom.Load(); err != nil {
				return nil, err
			} else {
				return pom, nil
			}
		}
	}
}

type pom struct {
	path  string
	model *Project
}

func SearchPom(path string) (Pom, error) {
	if exists, err := fileExists(path); err != nil {
		return nil, err
	} else if exists {
		if filepath.Base(path) == "pom.xml" {
			return LoadPom(filepath.Dir(path))
		} else {
			if exists, err := fileExists(filepath.Join(path, "pom.xml")); err != nil {
				return nil, err
			} else if exists {
				return LoadPom(path)
			} else {
				parent := filepath.Dir(path)
				if parent == path {
					return nil, nil
				} else {
					return SearchPom(parent)
				}
			}
		}
	} else {
		if path == "." {
			return nil, nil
		} else {
			parent := filepath.Dir(path)
			if parent == path {
				return nil, nil
			} else {
				return SearchPom(parent)
			}
		}
	}
}

func (p *pom) Store() error {
	if data, err := Marshal(p.model); err != nil {
		return err
	} else {
		return os.WriteFile(p.path, data, os.ModePerm)
	}
}

func (p *pom) Load() error {
	if data, err := os.ReadFile(p.path); err != nil {
		return err
	} else {
		if project, err := Unmarshal(data); err != nil {
			return err
		} else {
			p.model = project
			return nil
		}
	}
}
