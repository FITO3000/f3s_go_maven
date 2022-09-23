package maven

type ModelReader interface {
	GetParent() *Parent
	HasParent() bool
	GetCoordinates() *Coordinates
	GetDependency(c *Coordinates) *Dependency
	HasDependency(c *Coordinates) bool
	GetManagedDependency(c *Coordinates) *Dependency
	HasManagedDependency(c *Coordinates) bool
	GetProperties() *map[string]string
	GetProperty() string
	HasProperty(k string) bool
}

type ModelWriter interface {
	SetParent(p *Parent)
	RemoveParent() *Parent
	SetCoordinates(c *Coordinates)
	AddDependency(d *Dependency)
	RemoveDependency(c *Coordinates) *Dependency
	AddManagedDependency(d *Dependency)
	AddManagedDependencies(d []*Dependency)
	RemoveManagedDependency(c *Coordinates) *Dependency
	AddProperty(k, v string)
	AddProperties(p map[string]string)
}

type ModelReadWriter interface {
	ModelReader
	ModelWriter
}
