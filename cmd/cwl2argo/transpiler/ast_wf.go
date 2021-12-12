package transpiler

type LoadListingEnum interface{}
type NoListing struct{}
type ShallowListing struct{}
type DeepListing struct{}

type InputEnumSchema struct {
	Symbols []string
	Label   string
	Doc     string
	Name    string
}
type InputArraySchema struct {
	Items int // TODO
	Label string
	Doc   string
	Name  string
}
type InputRecordField struct {
	Name           string
	Type           CWLType
	Doc            string
	Label          string
	SecondaryFiles int // TODO
	Streamable     bool
	Format         CWLFormat
	LoadContents   bool
	LoadListing    LoadListingEnum
}

type InputRecordSchema struct {
	Label string
	Doc   string
	Name  string
}

type WorkflowInputParameterType interface{}

type InputBinding struct {
	LoadContents bool
}

type WorkflowInputParameter struct {
	Type           WorkflowInputParameterType
	Label          string
	SecondaryFiles int //TODO
	Streamable     bool
	Doc            string
	Id             string
	Format         CWLFormat
	LoadContents   bool
	LoadListing    LoadListingEnum
	Default        interface{}
	InputBinding   InputBinding
}

type WorkflowOutputParameterType interface{}

type MergeNested struct{}
type MergeFlattened struct{}
type LinkMergeMethod interface{}

type FirstNonNull struct{}
type TheOnlyNonNull struct{}
type AllNonNull struct{}
type PickValueMethod interface{}

type WorkflowOutputParameter struct {
	Type           WorkflowOutputParameterType
	Label          string
	SecondaryFiles int //TODO
	Streamable     bool
	Doc            string
	Id             string
	Format         CWLFormat
	OutputSource   string
}

type WorkflowStepInput struct {
	Id           string
	Source       string
	LinkMerge    LinkMergeMethod
	PickValue    PickValueMethod
	LoadContents bool
	LoadListing  LoadListingEnum
	Label        string
	Default      interface{}
	ValueFrom    CWLExpressionString
}

type WorkflowStepOutput struct {
	Id string
}

type WorkflowStep struct {
	In  []WorkflowStepInput
	Out []WorkflowStepOutput
	Run int // TODO
}

type Workflow struct {
	Inputs       []WorkflowInputParameter
	Outputs      []WorkflowOutputParameter
	Steps        []WorkflowStep
	Id           string
	Label        string
	Doc          string
	Requirements []CWLRequirements
}
