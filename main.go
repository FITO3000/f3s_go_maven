package main

import (
	"github.com/fit-o-matic/go-maven/maven"
)

func main() {

	dir := "/home/fit/temp/maven-test"

	maven.DeletePom(dir)

	if pom, err := maven.CreatePom(dir); err != nil {
		panic(err)
	} else {
		pom.SetParent(&maven.Parent{
			GroupId:    "tech.f3s",
			ArtifactId: "f3s-parent",
			Version:    "1",
		})
		pom.SetCoordinates(&maven.Coordinates{
			GroupId:    "tech.f3s",
			ArtifactId: "f3s-service",
			Version:    "1.0.0-SNAPSHOT",
		})

		pom.AddProperty("a.version", "1.9")
		pom.AddProperty("b.version", "1.1")
		pom.AddProperty("c.version", "2.0")

		pom.AddProperties(map[string]string{
			"d.version": "4.0.0",
			"e.version": "4.4.0",
			"f.version": "4.0.4",
		})

		pom.AddManagedDependency(&maven.Dependency{
			GroupId:    "tech.f3s",
			ArtifactId: "testing",
			Version:    "1.0.0",
		})

		pom.AddManagedDependencies([]*maven.Dependency{
			{
				GroupId:    "tech.f3s",
				ArtifactId: "lib-1",
				Version:    "1.0.0",
			},
			{
				GroupId:    "tech.f3s",
				ArtifactId: "lib-2",
				Version:    "1.0.0-SNAPSHOT",
			},
		})

		if err := pom.Store(); err != nil {
			panic(err)
		}
	}
}
