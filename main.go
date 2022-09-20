package main

import (
	"fmt"
	"github/FITO3000/f3s_go_maven/maven"
)

func main() {
	p := &maven.Project{
		Parent: &maven.Parent{
			GAV: maven.GAV{
				GroupId:    "tech.f3s.parent",
				ArtifactId: "f3s-parent",
				Version:    "1",
			},
		},
		GAV: maven.GAV{
			GroupId:    "tech.f3s.test",
			ArtifactId: "test-pom",
			Version:    "1.0.0-SNAPSHOT",
		},
		Packaging: "pom",
		DependencyManagement: &maven.DependencyManagement{
			Dependencies: &maven.Dependencies{
				Dependency: []maven.Dependency{
					{
						GAV: maven.GAV{
							GroupId:    "org.oss-a",
							ArtifactId: "oss-a-super-lib",
							Version:    "7.0.1",
						},
					},
					{
						GAV: maven.GAV{
							GroupId:    "org.oss-b",
							ArtifactId: "oss-a-super-lib",
							Version:    "1.9.4",
						},
					},
				},
			},
		},
		Dependencies: &maven.Dependencies{
			Dependency: []maven.Dependency{
				{
					GAV: maven.GAV{
						GroupId:    "G1",
						ArtifactId: "A1",
						Version:    "1",
					},
				},
			},
		},
		Properties: &maven.Properties{
			Entries: map[string]string{
				"A": "1",
				"B": "2",
			},
		},
		Modules: &maven.Modules{
			"p1",
			"p2",
		},
		Build: &maven.Build{
			FinalName: "final-name",
		},
	}

	pom, _ := maven.Marshal(p)
	fmt.Println(string(pom))
}
