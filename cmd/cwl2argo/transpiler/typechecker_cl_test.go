package transpiler

import "testing"

var (
	exampleCLI1Id          = "exampleCLI1"
	exampleCLIRequirements = []CWLRequirements{DockerRequirement{}}
)

var exampleCLI1 = CommandlineTool{
	Inputs:       make([]CommandlineInputParameter, 0),
	Outputs:      make([]CommandlineOutputParameter, 0),
	Class:        "CommandlineTool",
	Id:           &exampleCLI1Id,
	Doc:          make([]string, 0),
	Requirements: make([]CWLRequirements, 0),
	Hints:        make([]interface{}, 0),
	CWLVersion:   nil,
	Intent:       make([]string, 0),
	BaseCommand:  make([]string, 0),
	Arguments:    make([]CommandlineArgument, 0),
	Stdin:        nil,
	Stderr:       nil,
	Stdout:       nil,
}

func TestCLIRequirementTypeChecking(t *testing.T) {
	err := TypeCheckCommandlineTool(exampleCLI1)
	if err == nil {
		t.Errorf("Failed to type check: %s", err)
	}
	exampleCLI1.Requirements = exampleCLIRequirements

	err = TypeCheckCommandlineTool(exampleCLI1)
	if err != nil {
		t.Errorf("Failed to type check: %s", err)
	}
}
