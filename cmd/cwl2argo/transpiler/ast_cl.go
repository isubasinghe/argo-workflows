package transpiler

type CWLRequirements interface {
	isCWLRequirement()
}

type DockerRequirement struct {
	DockerPull            string
	DockerLoad            string
	DockerFile            string
	DockerImport          string
	DockerImageId         string
	DockerOutputDirectory string
}

type SoftwarePackage struct {
	Package string
	Version []string
	Specs   []string
}
type SoftwareRequirement struct {
	Packages []SoftwarePackage
}

type LoadListingRequirement struct {
	LoadListing LoadListingEnum
}

type Dirent struct{}

type InitialWorkDirRequirement struct {
}

type EnvironmentDef struct {
	EnvName  string
	EnvValue CWLExpressionString
}

type EnvVarRequirement struct {
	EnvDef []EnvironmentDef
}

func (_ *DockerRequirement) isCWLRequirement() {}

type CommandlineInputParameterType interface{}

type CommandlineBindingPosition interface{}

type CommandlineBinding struct {
	LoadContents  bool
	Position      CommandlineBindingPosition
	Prefix        string
	Seperate      bool
	ItemSeperator string
	ValueFrom     CWLExpressionString
	ShellQuote    bool
}

type CommandlineInputParameter struct {
	Type           CommandlineInputParameterType
	Label          string
	SecondaryFiles int // TODO
	Streamable     bool
	Doc            string
	Id             string
	Format         CWLFormat
	LoadContents   bool
	LoadListing    LoadListingEnum
	Default        interface{}
	InputBinding   CommandlineBinding
}

type CommandlineOutputBindingGlob interface{}

type CommandlineOutputBinding struct {
	LoadContents bool
	LoadListing  LoadListingEnum
	Glob         CommandlineOutputBindingGlob
	OutputEval   CWLExpression
}

type CommandlineOutputParameter struct {
	Type           CommandlineInputParameterType
	Label          string
	SecondaryFiles int // TODO
	Streamable     bool
	Doc            string
	Id             string
	Format         CWLFormat
	OutputBinding  CommandlineOutputBinding
}

type CommandlineArguments interface{}

type CWLExpressionString interface{}

type CommandlineTool struct {
	Inputs       []CommandlineInputParameter
	Outputs      []CommandlineOutputParameter
	BaseCommand  []string
	Id           string
	Label        string
	Doc          []string
	Requirements []CWLRequirements
	Arguments    []CommandlineArguments
	Stdin        CWLExpressionString
	Stderr       CWLExpressionString
	Stdout       CWLExpressionString
}
