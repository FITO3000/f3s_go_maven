package maven

import "golang.org/x/exp/slices"

// Coordinates

func NewCoordinates(g, a, v string) Coordinates {
	return Coordinates{
		GroupId:    g,
		ArtifactId: a,
		Version:    v,
	}
}

func (c *Coordinates) RemoveVersion() {
	c.Version = ""
}

func (c *Coordinates) HasVersion() bool {
	return c.Version != ""
}

func (c *Coordinates) CreateDependency() *Dependency {
	return &Dependency{
		GroupId:    c.GroupId,
		ArtifactId: c.ArtifactId,
		Version:    c.Version,
	}
}

func (c *Coordinates) AssignTo(p *Project) {
	p.GroupId = c.GroupId
	p.ArtifactId = c.ArtifactId
	p.Version = c.Version
}

func (c *Coordinates) Similar(other *Coordinates) bool {
	return c.GroupId == other.GroupId && c.ArtifactId == other.ArtifactId
}

func (c *Coordinates) Equal(other *Coordinates) bool {
	return c.Similar(other) && c.Version == other.Version
}

// Project

func (p *Project) GetDependency(c *Coordinates) *Dependency {
	return p.provideDependencies().GetDependency(c)
}

func (p *Project) HasDependency(c *Coordinates) bool {
	return p.provideDependencies().HasDependency(c)
}

func (p *Project) GetManagedDependency(c *Coordinates) *Dependency {
	return p.provideDependencyManagement().Dependencies.GetDependency(c)
}

func (p *Project) HasManagedDependency(c *Coordinates) bool {
	return p.provideDependencyManagement().Dependencies.HasDependency(c)
}

func (p *Project) GetCoordinates() *Coordinates {
	return &Coordinates{
		GroupId:    p.GroupId,
		ArtifactId: p.ArtifactId,
		Version:    p.Version,
	}
}

func (p *Project) SetCoordinates(c *Coordinates) {
	p.GroupId = c.GroupId
	p.ArtifactId = c.ArtifactId
	p.Version = c.Version
}

func (p *Project) provideDependencies() *Dependencies {
	if p.Dependencies == nil {
		p.Dependencies = NewDependencies()
	}
	return p.Dependencies
}

func (p *Project) provideDependencyManagement() *DependencyManagement {
	if p.DependencyManagement == nil {
		p.DependencyManagement = NewDependencyManagement()
	}
	return p.DependencyManagement
}

func (p *Project) provideProperties() *Properties {
	if p.Properties == nil {
		p.Properties = NewProperties()
	}
	return p.Properties
}

func (p *Project) HasProperty(k string) bool {
	if properties := p.Properties; properties == nil {
		return false
	} else {
		return properties.HasProperty(k)
	}
}

func (p *Project) AddProperty(k, v string) {
	p.provideProperties().AddProperty(k, v)
}

func (p *Project) AddProperties(props map[string]string) {
	p.provideProperties().AddProperties(props)
}

func (p *Project) AddDependency(d *Dependency) {
	p.provideDependencies().AddDependency(d)
}

func (p *Project) AddDependencies(d []*Dependency) {
	p.provideDependencies().AddDependencies(d)
}

func (p *Project) RemoveDependency(c *Coordinates) *Dependency {
	return p.provideDependencies().RemoveDependency(c)
}

func (p *Project) AddManagedDependency(d *Dependency) {
	p.provideDependencyManagement().Dependencies.AddDependency(d)
}

func (p *Project) AddManagedDependencies(d []*Dependency) {
	p.provideDependencyManagement().Dependencies.AddDependencies(d)
}

func (p *Project) RemoveManagedDependency(c *Coordinates) *Dependency {
	return p.provideDependencyManagement().Dependencies.RemoveDependency(c)
}

// Dependency

func (d *Dependency) GetCoordinates() *Coordinates {
	return &Coordinates{
		GroupId:    d.GroupId,
		ArtifactId: d.ArtifactId,
		Version:    d.Version,
	}
}

// Dependencies

func NewDependencies() *Dependencies {
	return &Dependencies{
		Values: []Dependency{},
	}
}

func (d *Dependencies) search(c *Coordinates) int {
	return slices.IndexFunc(d.Values, func(d Dependency) bool {
		return d.GetCoordinates().Similar(c)
	})
}
func (d *Dependencies) HasDependency(c *Coordinates) bool {
	return d.search(c) >= 0
}

func (d *Dependencies) GetDependency(c *Coordinates) *Dependency {
	if idx := d.search(c); idx >= 0 {
		return &d.Values[idx]
	} else {
		return nil
	}
}

func (d *Dependencies) AddDependency(dependency *Dependency) {
	if idx := d.search(dependency.GetCoordinates()); idx < 0 {
		d.Values = append(d.Values, *dependency)
	} else {
		d.Values[idx] = *dependency
	}
}

func (d *Dependencies) AddDependencies(deps []*Dependency) {
	for _, dep := range deps {
		d.AddDependency(dep)
	}
}

func (d *Dependencies) RemoveDependency(c *Coordinates) *Dependency {
	if idx := d.search(c); idx >= 0 {
		res := d.Values[idx]
		copy(d.Values[idx:], d.Values[idx+1:])
		d.Values = d.Values[:len(d.Values)-1]
		return &res
	} else {
		return nil
	}
}

// DependencyManagement

func NewDependencyManagement() *DependencyManagement {
	return &DependencyManagement{
		Dependencies: NewDependencies(),
	}
}

// Properties

func NewProperties() *Properties {
	return &Properties{
		Entries: map[string]string{},
	}
}

func (p *Properties) HasProperty(k string) bool {
	_, exists := p.Entries[k]
	return exists
}

func (p *Properties) AddProperty(k, v string) {
	p.Entries[k] = v
}

func (p *Properties) AddProperties(props map[string]string) {
	for k, v := range props {
		p.AddProperty(k, v)
	}
}
