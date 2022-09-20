package maven

import (
	"encoding/xml"
)

// https://maven.apache.org/pom.html

type Project struct {
	XMLName                xml.Name                `xml:"project"`
	ModelVersion           string                  `xml:"modelVersion,omitempty"`
	Parent                 *Parent                 `xml:"parent,omitempty"`
	GroupId                string                  `xml:"groupId,omitempty"`
	ArtifactId             string                  `xml:"artifactId,omitempty"`
	Version                string                  `xml:"version,omitempty"`
	Packaging              string                  `xml:"packaging,omitempty"`
	Properties             *Properties             `xml:"properties,omitempty"`
	Dependencies           *Dependencies           `xml:"dependencies,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Modules                *Modules                `xml:"modules>module,omitempty"`
	Build                  *Build                  `xml:"build,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
	Name                   string                  `xml:"name,omitempty"`
	Description            string                  `xml:"description,omitempty"`
	Url                    string                  `xml:"url,omitempty"`
	InceptionYear          string                  `xml:"inceptionYear,omitempty"`
	Licenses               *Licenses               `xml:"licenses>license,omitempty"`
	Organization           *Organization           `xml:"organization,omitempty"`
	Developers             *Developers             `xml:"developers>developer,omitempty"`
	Contributors           *Contributors           `xml:"contributors>contributor,omitempty"`
	IssueManagement        *IssueManagement        `xml:"issueManagement,omitempty"`
	CiManagement           *CiManagement           `xml:"ciManagement,omitempty"`
	MailingLists           *MailingLists           `xml:"mailingLists>mailingList,omitempty"`
	Scm                    *Scm                    `xml:"scm,omitempty"`
	Prerequisites          *Prerequisites          `xml:"prerequisites,omitempty"`
	Repositories           *Repositories           `xml:"repositories>repository,omitempty"`
	PluginRepositories     *Repositories           `xml:"pluginRepositories>repository,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	Profiles               *Profiles               `xml:"profiles>profile,omitempty"`
}

type Dependencies struct {
	Dependency []Dependency `xml:"dependency,omitempty"`
}

type Dependency struct {
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Type       string `xml:"type,omitempty"`
	Scope      string `xml:"scope,omitempty"`
	Optional   bool   `xml:"optional,omitempty"`
}

type Parent struct {
	GroupId      string `xml:"groupId,omitempty"`
	ArtifactId   string `xml:"artifactId,omitempty"`
	Version      string `xml:"version,omitempty"`
	RelativePath string `xml:"relativePath,omitempty"`
}

type DependencyManagement struct {
	Dependencies *Dependencies `xml:"dependencies,omitempty"`
}

type Modules []string

type Properties struct {
	Entries map[string]string
}

type Build struct {
	DefaultGoal           string
	Directory             string
	FinalName             string
	Filters               *Filters
	Resources             *Resources
	TestResources         *Resources
	Plugins               *Plugins
	PluginManagement      *PluginManagement
	SourceDirectory       string
	ScriptSourceDirectory string
	TestSourceDirectory   string
	OutputDirectory       string
	TestOutputDirectory   string
	Extensions            *Extensions
}

type Filters []string

type Resources []Resource

type Resource struct {
	TargetPath string
	Filtering  bool
	Directory  string
	Includes   *Includes
	Excludes   *Excludes
}

type Includes []string

type Excludes []string

type Plugins []Plugin

type Plugin struct {
	GroupId       string
	ArtifactId    string
	Version       string
	Extensions    bool
	Inherited     bool
	Configuration *Properties
	Dependencies  *Dependencies
	Executions    *Executions
}

type Executions []Execution

type Execution struct {
	Id            string
	Goals         *Goals
	Phase         string
	Inherited     bool
	Configuration *Configuration
}

type Goals []string

type Configuration []interface{}

type PluginManagement []Plugin

type Extensions []Extension

type Extension struct {
	GroupId    string
	ArtifactId string
	Version    string
}

type Reporting struct {
	OutputDirectory string
	Plugins         *Plugins
}

type Licenses []License

type License struct {
	Name         string
	Url          string
	Distribution string
	Comments     string
}

type Organization struct {
	Name string
	Url  string
}

type Developers []Person

type Person struct {
	Id              string
	Name            string
	Email           string
	Url             string
	Organization    string
	OrganizationUrl string
	Roles           *Roles
	Timezone        string
	Properties      *Properties
}

type Roles []string

type Contributors []Person

type IssueManagement struct {
	System string
	Url    string
}

type CiManagement struct {
	System    string
	Url       string
	Notifiers *Notifiers
}

type Notifiers []Notifier

type Notifier struct {
	Type          string
	SendOnError   string
	SendOnFailure string
	SendOnSuccess string
	SendOnWarning string
	Configuration *Configuration
}

type MailingLists []MailingList

type MailingList struct {
	Name          string
	Subscribe     string
	Unsubscribe   string
	Post          string
	Archive       string
	OtherArchives *OtherArchives
}

type OtherArchives []string

type Scm struct {
	Connection          string
	DeveloperConnection string
	Tag                 string
	Url                 string
}

type Prerequisites struct {
	Maven string
}

type Repositories []Repository

type Repository struct {
	Releases  *RepositoryPolicies
	Snapshots *RepositoryPolicies
	Name      string
	Id        string
	Url       string
	Layout    string
}

type RepositoryPolicies struct {
	Enabled        bool
	UpdatePolicy   string
	ChecksumPolicy string
}

type DistributionManagement struct {
	Repository         *DistributionRepository
	SnapshotRepository *DistributionRepository
	Site               *Site
	Relocation         *Relocation
	DownloadUrl        string
	Status             string
}

type DistributionRepository struct {
	UniqueVersion bool
	Id            string
	Name          string
	Url           string
	Layout        string
}
type Site struct {
	Id   string
	Name string
	Url  string
}

type Relocation struct {
	GroupId    string
	ArtifactId string
	Version    string
	Message    string
}

type Profiles []Profile

type Profile struct {
	Id                     string
	Activation             *Activation
	Build                  *Build
	Modules                *Modules
	Repositories           *Repositories
	PluginRepositories     *Repositories
	Dependencies           *Dependencies
	Reporting              *Reporting
	DependencyManagement   *DependencyManagement
	DistributionManagement *DistributionManagement
}

type Activation struct {
	ActiveByDefault bool
	Jdk             string
	Os              *Os
	Property        *Property
	File            *File
}

type Os struct {
	Name    string
	Family  string
	Arch    string
	Version string
}

type Property struct {
	Name  string
	Value string
}

type File struct {
	Exists  string
	Missing string
}
