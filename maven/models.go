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
	DefaultGoal           string            `xml:"defaultGoal,omitempty"`
	Directory             string            `xml:"directory,omitempty"`
	FinalName             string            `xml:"finalName,omitempty"`
	Filters               *Filters          `xml:"filters,omitempty"`
	Resources             *Resources        `xml:"resources,omitempty"`
	TestResources         *Resources        `xml:"testResources,omitempty"`
	Plugins               *Plugins          `xml:"plugins,omitempty"`
	PluginManagement      *PluginManagement `xml:"pluginManagement,omitempty"`
	SourceDirectory       string            `xml:"sourceDirectory,omitempty"`
	ScriptSourceDirectory string            `xml:"scriptSourceDirectory,omitempty"`
	TestSourceDirectory   string            `xml:"testSourceDirectory,omitempty"`
	OutputDirectory       string            `xml:"outputDirectory,omitempty"`
	TestOutputDirectory   string            `xml:"testOutputDirectory,omitempty"`
	Extensions            *Extensions       `xml:"extensions,omitempty"`
}

type Filters []string

type Resources []Resource

type Resource struct {
	TargetPath string    `xml:"targetPath,omitempty"`
	Filtering  bool      `xml:"filtering,omitempty"`
	Directory  string    `xml:"directory,omitempty"`
	Includes   *Includes `xml:"includes,omitempty"`
	Excludes   *Excludes `xml:"excludes,omitempty"`
}

type Includes []string

type Excludes []string

type Plugins []Plugin

type Plugin struct {
	GroupId       string        `xml:"groupId,omitempty"`
	ArtifactId    string        `xml:"artifactId,omitempty"`
	Version       string        `xml:"version,omitempty"`
	Extensions    bool          `xml:"extensions,omitempty"`
	Inherited     bool          `xml:"inherited,omitempty"`
	Configuration *Properties   `xml:"configuration,omitempty"`
	Dependencies  *Dependencies `xml:"dependencies,omitempty"`
	Executions    *Executions   `xml:"executions,omitempty"`
}

type Executions []Execution

type Execution struct {
	Id            string         `xml:"id,omitempty"`
	Goals         *Goals         `xml:"goals,omitempty"`
	Phase         string         `xml:"phase,omitempty"`
	Inherited     bool           `xml:"inherited,omitempty"`
	Configuration *Configuration `xml:"configuration,omitempty"`
}

type Goals []string

type Configuration []interface{}

type PluginManagement []Plugin

type Extensions []Extension

type Extension struct {
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
}

type Reporting struct {
	OutputDirectory string   `xml:"outputDirectory,omitempty"`
	Plugins         *Plugins `xml:"plugins,omitempty"`
}

type Licenses []License

type License struct {
	Name         string `xml:"name,omitempty"`
	Url          string `xml:"url,omitempty"`
	Distribution string `xml:"distribution,omitempty"`
	Comments     string `xml:"comments,omitempty"`
}

type Organization struct {
	Name string `xml:"name,omitempty"`
	Url  string `xml:"url,omitempty"`
}

type Developers []Person

type Person struct {
	Id              string      `xml:"id,omitempty"`
	Name            string      `xml:"name,omitempty"`
	Email           string      `xml:"email,omitempty"`
	Url             string      `xml:"url,omitempty"`
	Organization    string      `xml:"organization,omitempty"`
	OrganizationUrl string      `xml:"organizationUrl,omitempty"`
	Roles           *Roles      `xml:"roles,omitempty"`
	Timezone        string      `xml:"timezone,omitempty"`
	Properties      *Properties `xml:"properties,omitempty"`
}

type Roles []string

type Contributors []Person

type IssueManagement struct {
	System string `xml:"system,omitempty"`
	Url    string `xml:"url,omitempty"`
}

type CiManagement struct {
	System    string     `xml:"system,omitempty"`
	Url       string     `xml:"url,omitempty"`
	Notifiers *Notifiers `xml:"notifiers,omitempty"`
}

type Notifiers []Notifier

type Notifier struct {
	Type          string         `xml:"type,omitempty"`
	SendOnError   string         `xml:"sendOnError,omitempty"`
	SendOnFailure string         `xml:"sendOnFailure,omitempty"`
	SendOnSuccess string         `xml:"sendOnSuccess,omitempty"`
	SendOnWarning string         `xml:"sendOnWarning,omitempty"`
	Configuration *Configuration `xml:"configuration,omitempty"`
}

type MailingLists []MailingList

type MailingList struct {
	Name          string         `xml:"name,omitempty"`
	Subscribe     string         `xml:"subscribe,omitempty"`
	Unsubscribe   string         `xml:"unsubscribe,omitempty"`
	Post          string         `xml:"post,omitempty"`
	Archive       string         `xml:"archive,omitempty"`
	OtherArchives *OtherArchives `xml:"otherArchives,omitempty"`
}

type OtherArchives []string

type Scm struct {
	Connection          string `xml:"connection,omitempty"`
	DeveloperConnection string `xml:"developerConnection,omitempty"`
	Tag                 string `xml:"tag,omitempty"`
	Url                 string `xml:"url,omitempty"`
}

type Prerequisites struct {
	Maven string `xml:"maven,omitempty"`
}

type Repositories []Repository

type Repository struct {
	Releases  *RepositoryPolicies `xml:"releases,omitempty"`
	Snapshots *RepositoryPolicies `xml:"snapshots,omitempty"`
	Name      string              `xml:"name,omitempty"`
	Id        string              `xml:"id,omitempty"`
	Url       string              `xml:"url,omitempty"`
	Layout    string              `xml:"layout,omitempty"`
}

type RepositoryPolicies struct {
	Enabled        bool   `xml:"enabled,omitempty"`
	UpdatePolicy   string `xml:"updatePolicy,omitempty"`
	ChecksumPolicy string `xml:"checksumPolicy,omitempty"`
}

type DistributionManagement struct {
	Repository         *DistributionRepository `xml:"repository,omitempty"`
	SnapshotRepository *DistributionRepository `xml:"snapshotRepository,omitempty"`
	Site               *Site                   `xml:"site,omitempty"`
	Relocation         *Relocation             `xml:"relocation,omitempty"`
	DownloadUrl        string                  `xml:"downloadUrl,omitempty"`
	Status             string                  `xml:"status,omitempty"`
}

type DistributionRepository struct {
	UniqueVersion bool   `xml:"uniqueVersion,omitempty"`
	Id            string `xml:"id,omitempty"`
	Name          string `xml:"name,omitempty"`
	Url           string `xml:"url,omitempty"`
	Layout        string `xml:"layout,omitempty"`
}
type Site struct {
	Id   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	Url  string `xml:"url,omitempty"`
}

type Relocation struct {
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Message    string `xml:"message,omitempty"`
}

type Profiles []Profile

type Profile struct {
	Id                     string                  `xml:"id,omitempty"`
	Activation             *Activation             `xml:"activation,omitempty"`
	Build                  *Build                  `xml:"build,omitempty"`
	Modules                *Modules                `xml:"modules,omitempty"`
	Repositories           *Repositories           `xml:"repositories,omitempty"`
	PluginRepositories     *Repositories           `xml:"pluginRepositories,omitempty"`
	Dependencies           *Dependencies           `xml:"dependencies,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
}

type Activation struct {
	ActiveByDefault bool      `xml:"activeByDefault,omitempty"`
	Jdk             string    `xml:"jdk,omitempty"`
	Os              *Os       `xml:"os,omitempty"`
	Property        *Property `xml:"property,omitempty"`
	File            *File     `xml:"file,omitempty"`
}

type Os struct {
	Name    string `xml:"name,omitempty"`
	Family  string `xml:"family,omitempty"`
	Arch    string `xml:"arch,omitempty"`
	Version string `xml:"version,omitempty"`
}

type Property struct {
	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type File struct {
	Exists  string `xml:"exists,omitempty"`
	Missing string `xml:"missing,omitempty"`
}
