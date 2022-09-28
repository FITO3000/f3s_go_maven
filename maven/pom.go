package maven

import (
	"fmt"
	"os"
	"path/filepath"
)

type Pom interface {
	Store() error
	Load() error
	GetPath() string
	GetDirectory() string

	ModelReadWriter
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

// LoadPom loads a pom.xml file from directory and returns a Pom structure
// If pom.xml does not exist, LoadPom returns an error
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

func (p *pom) GetPath() string {
	return p.path
}

func (p *pom) GetDirectory() string {
	return filepath.Dir(p.path)
}

func (p *pom) GetParent() *Parent {
	return p.model.Parent
}

func (p *pom) HasParent() bool {
	return p.model.Parent != nil
}

func (p *pom) GetCoordinates() *Coordinates {
	return p.model.GetCoordinates()
}

func (p *pom) GetDependency(c *Coordinates) *Dependency {
	return p.model.GetDependency(c)
}

func (p *pom) HasDependency(c *Coordinates) bool {
	return p.model.HasDependency(c)
}

func (p *pom) GetManagedDependency(c *Coordinates) *Dependency {
	return p.model.GetManagedDependency(c)
}

func (p *pom) HasManagedDependency(c *Coordinates) bool {
	return p.model.HasManagedDependency(c)
}

func (p *pom) GetProperties() *map[string]string {
	panic("not implemented") // TODO: Implement
}

func (p *pom) GetProperty() string {
	panic("not implemented") // TODO: Implement
}

func (p *pom) HasProperty(k string) bool {
	return p.model.HasProperty(k)
}

func (p *pom) AddProperty(k, v string) {
	p.model.provideProperties().AddProperty(k, v)
}

func (p *pom) AddProperties(props map[string]string) {
	p.model.provideProperties().AddProperties(props)
}

func (p *pom) SetParent(parent *Parent) {
	p.model.SetParent(parent)
}

func (p *pom) RemoveParent() *Parent {
	return p.model.RemoveParent()
}

func (p *pom) SetCoordinates(c *Coordinates) {
	p.model.SetCoordinates(c)
}

func (p *pom) AddDependency(d *Dependency) {
	p.model.AddDependency(d)
}

func (p *pom) RemoveDependency(c *Coordinates) *Dependency {
	return p.model.RemoveDependency(c)
}

func (p *pom) AddManagedDependency(d *Dependency) {
	p.model.AddManagedDependency(d)
}

func (p *pom) AddManagedDependencies(d []*Dependency) {
	p.model.AddManagedDependencies(d)
}

func (p *pom) RemoveManagedDependency(c *Coordinates) *Dependency {
	return p.model.RemoveManagedDependency(c)
}
