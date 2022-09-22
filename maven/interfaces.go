package maven

type ProjectFacade interface {
	SetParent(parent Parent)
	SetCoordinates(groupId, artifactId, version string)
	SetPackaging(packaging string)
	AddProperties(properties map[string]string)
	AddmanagedDependency(dependency *Dependency)
	AddDependency(dependency *Dependency)
}
