package main

import (
	"fmt"
	"github/FITO3000/f3s_go_maven/maven"
)

func main() {
	p := maven.NewProject()

	p.SetCoordinates("tech.f3s.app", "test-app-1", "1.7.0")

	p.SetParent(&maven.Parent{
		GroupId:    "tech.f3s.parent",
		ArtifactId: "f3s-parent",
		Version:    "1",
	})

	p.AddProperties(map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})

	p.AddManagedDependency(maven.Dependency{
		GroupId:    "GG01",
		ArtifactId: "AA01",
		Version:    "${aa.version}",
	})

	p.AddDependency(maven.Dependency{
		GroupId:    "GG01",
		ArtifactId: "AA01",
	})

	p.AddDependency(maven.Dependency{
		GroupId:    "GG01",
		ArtifactId: "AA02",
		Version:    "1.2.3",
	})

	p.AddDependency(maven.Dependency{
		GroupId:    "GG01",
		ArtifactId: "AA01",
		Version:    "2.0",
	})

	p.AddModules([]string{
		"p1",
		"p2",
		"p7",
	})

	p.AddModules([]string{
		"p1",
		"p3",
		"p4",
	})

	pom, _ := maven.Marshal(p)
	fmt.Println(string(pom))
}
