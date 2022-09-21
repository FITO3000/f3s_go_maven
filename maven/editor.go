package maven

type Editor struct {
	project *Project
}

func NewEditor(project *Project) Editor {
	return Editor{
		project: project,
	}
}

func (e Editor) AddDependency(dependency Dependency) {
	if e.project.Dependencies == nil {
		e.project.Dependencies = &Dependencies{
			Values: []Dependency{
				dependency,
			},
		}
	} else {

	}
}
