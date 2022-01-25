package transpiler

import (
	"errors"
	"fmt"
	"sort"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	log "github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
)

const (
	ArgoType    = "Workflow"
	ArgoVersion = "argoproj.io/v1alpha1"
)

// de-sum typed "CommandlineInputParameter
// contains all necessary data to output argo yaml
type flatCommandlineInputParameter struct {
	Type             Type
	Label            *string
	Value            *string
	Emit             bool
	File             *CWLFile
	FileLocationData *FileLocationData
	SecondaryFiles   *SecondaryFiles
	Streamable       *bool
	Doc              Strings
	Id               *string
	Format           *CWLFormat
	LoadContents     *bool
	LoadListing      *LoadListingEnum
	InputBinding     *CommandlineBinding
}

type ParamTranslater interface {
	TranslateToParam(*CommandlineBinding) ([]v1alpha1.Parameter, error)
}

func emitDockerRequirement(container *apiv1.Container, d *DockerRequirement) error {
	tmpContainer := container.DeepCopy()

	if d.DockerPull == nil {
		return errors.New("dockerPull is a required field")
	}

	tmpContainer.Image = *d.DockerPull

	if d.DockerFile != nil {
		return errors.New("")
	}

	if d.DockerImageId != nil {
		return errors.New("")
	}

	if d.DockerImport != nil {
		return errors.New("")
	}

	*container = *tmpContainer
	return nil
}

func emitInputParam(input flatCommandlineInputParameter) (*v1alpha1.Parameter, error) {
	name := *input.Id
	param := v1alpha1.Parameter{Name: name}
	return &param, nil
}

func dockerNotPresent() error { return errors.New("DockerRequirement was not found") }

func findDockerRequirement(clTool *CommandlineTool) (*DockerRequirement, error) {
	var docker *DockerRequirement
	docker = nil
	for _, req := range clTool.Requirements {
		d, ok := req.(DockerRequirement)
		if ok {
			log.Info("Found DockerRequirement")
			docker = &d
		}
	}

	if docker != nil {
		return docker, nil
	} else {
		return nil, dockerNotPresent()
	}
}

func emitInputParams(template *v1alpha1.Template, inputs []flatCommandlineInputParameter) error {
	params := make([]v1alpha1.Parameter, 0)
	for _, input := range inputs {
		newInput, err := emitInputParam(input)
		if err != nil {
			return err
		}
		params = append(params, *newInput)
	}
	template.Inputs.Parameters = params
	return nil
}

// dummy function to evaluate CommandlineTool
// until proper eval functionality is added
func evalArgument(arg CommandlineArgument) (*string, error) {
	switch arg.Kind {
	case ArgumentStringKind:
		return (*string)(&arg.String), nil
	default:
		return nil, errors.New("only string is accepted at the moment")
	}
}

func canFindType(input CWLInputEntry, tys CommandlineTypes) error {
	for _, currTy := range tys {
		if currTy.Kind == input.Kind {
			return nil
		}
	}
	return nil
}

func (inputParameter CommandlineInputParameter) getInputBindings(inputs map[string]CWLInputEntry) (*flatCommandlineInputParameter, error) {
	if inputParameter.Id == nil {
		return nil, errors.New("input parameter is nil")
	}

	input, ok := inputs[*inputParameter.Id]
	if !ok {
		return nil, fmt.Errorf("%s was not present in input", *inputParameter.Id)
	}

	binding := flatCommandlineInputParameter{
		SecondaryFiles: &inputParameter.SecondaryFiles,
		Streamable:     inputParameter.Streamable,
		Doc:            inputParameter.Doc,
		Id:             inputParameter.Id,
		Format:         inputParameter.Format,
		InputBinding:   inputParameter.InputBinding,
		Emit:           true,
	}

	err := canFindType(input, inputParameter.Type)
	if err != nil {
		return nil, err
	}

	binding.Type = input.Kind
	switch input.Kind {
	case CWLStringKind:
		binding.Value = input.StringData
	case CWLIntKind:
		strValue := fmt.Sprintf("%d", *input.IntData)
		binding.Value = &strValue
	case CWLFileKind:
		binding.File = input.FileData
	default:
		return nil, fmt.Errorf("%T unknown type", input.Kind)
	}
	return &binding, nil
}

func sortBindingsByPosition(bindings []flatCommandlineInputParameter) {
	sort.Slice(bindings[:], func(i, j int) bool {
		leftPost := 0
		rightPost := 0
		if bindings[i].InputBinding.Position != nil {
			leftPost = *bindings[i].InputBinding.Position
		}
		if bindings[i].InputBinding.Position != nil {
			rightPost = *bindings[j].InputBinding.Position
		}
		return leftPost < rightPost
	})
}

func emitArgumentParams(container *apiv1.Container,
	baseCommand Strings,
	arguments Arguments,
	bindings []flatCommandlineInputParameter,
	inputs map[string]CWLInputEntry) error {
	cmds := make([]string, 0)
	skip := false

	if len(baseCommand) == 0 {
		if len(arguments) == 0 {
			return errors.New("len(baseCommand)==0 && len(arguments)==0")
		}
		cmd, err := evalArgument(arguments[0])
		if err != nil {
			return err
		}
		cmds = append(cmds, *cmd)
		skip = false
	}

	for _, cmd := range baseCommand {
		cmds = append(cmds, cmd)
	}

	for i, arg := range arguments {
		if i == 0 && skip {
			continue
		}
		cmd, err := evalArgument(arg)
		if err != nil {
			return err
		}
		cmds = append(cmds, *cmd)
	}

	sortBindingsByPosition(bindings)

	args := make([]string, 0)
	for _, binding := range bindings {

		prefix := ""
		if binding.InputBinding != nil && binding.InputBinding.Prefix != nil {
			sep := true
			if binding.InputBinding.Separate != nil {
				sep = *binding.InputBinding.Separate
			}

			if sep {
				sepArg := *binding.InputBinding.Prefix
				args = append(args, sepArg)
			} else {
				prefix = *binding.InputBinding.Prefix
			}
		}
		var arg string
		arg = fmt.Sprintf("%s{{inputs.parameters.%s}}", prefix, *binding.Id)

		if binding.Type == CWLFileKind {
			if binding.InputBinding == nil {
				continue
			}
			if binding.File == nil || binding.File.Path == nil {
				return errors.New("file information was not available")
			}
			arg = *binding.File.Path
		}
		args = append(args, arg)
	}

	container.Command = cmds
	container.Args = args

	return nil
}

func emitArguments(spec *v1alpha1.WorkflowSpec, bindings []flatCommandlineInputParameter) error {
	params := make([]v1alpha1.Parameter, 0)
	arts := make([]v1alpha1.Artifact, 0)

	for _, binding := range bindings {
		switch binding.Type {
		case CWLStringKind:
			params = append(params, v1alpha1.Parameter{Name: *binding.Id, Value: (*v1alpha1.AnyString)(binding.Value)})
		default:
			return fmt.Errorf("%T is not supported", binding.Type)
		}
	}
	args := v1alpha1.Arguments{Parameters: params, Artifacts: arts}
	spec.Arguments = args
	return nil
}

func flatten(inputs Inputs, input map[string]CWLInputEntry) ([]flatCommandlineInputParameter, error) {
	flatInputs := make([]flatCommandlineInputParameter, 0)
	for _, inputBinding := range inputs {
		newBindings, err := inputBinding.getInputBindings(input)
		if err != nil {
			return nil, err
		}
		flatInputs = append(flatInputs, *newBindings)
	}
	return flatInputs, nil
}

func filterParams(inputs []flatCommandlineInputParameter) []flatCommandlineInputParameter {
	newInputs := make([]flatCommandlineInputParameter, 0)
	for _, input := range inputs {
		switch input.Type {
		case CWLFileKind:
			continue
		case CWLRecordFieldKind:
			continue
		case CWLArrayKind:
			continue
		case CWLEnumKind:
			continue
		default:
			newInputs = append(newInputs, input)
		}
	}
	return newInputs
}

func emitInputArtifacts(template *v1alpha1.Template, inputs map[string]CWLInputEntry, locations FileLocations) error {
	arts := make([]v1alpha1.Artifact, 0)

	for key, inputEntry := range inputs {
		if inputEntry.Kind != CWLFileKind {
			continue
		}
		location, ok := locations.Inputs[key]
		if !ok {
			return fmt.Errorf("location data not present for %s", key)
		}

		art := v1alpha1.Artifact{}
		art.Name = location.Name
		art.Path = *inputEntry.FileData.Path
		art.HTTP = location.HTTP
		arts = append(arts, art)
	}

	template.Inputs.Artifacts = arts
	return nil
}

func EmitCommandlineTool(clTool *CommandlineTool, inputs map[string]CWLInputEntry, locations FileLocations) (*v1alpha1.Workflow, error) {
	var wf v1alpha1.Workflow
	var err error

	wf.Name = *clTool.Id
	spec := v1alpha1.WorkflowSpec{}
	wf.APIVersion = ArgoVersion
	wf.Kind = ArgoType

	container := apiv1.Container{}

	dockerRequirement, err := findDockerRequirement(clTool)
	if err != nil {
		return nil, err
	}

	err = emitDockerRequirement(&container, dockerRequirement)
	if err != nil {
		return nil, err
	}

	template := v1alpha1.Template{}
	template.Container = &container
	template.Name = *clTool.Id

	bindings, err := flatten(clTool.Inputs, inputs)
	if err != nil {
		return nil, err
	}

	paramBindings := filterParams(bindings)

	err = emitInputParams(&template, paramBindings)
	if err != nil {
		return nil, err
	}

	err = emitArgumentParams(&container, clTool.BaseCommand, clTool.Arguments, bindings, inputs)
	if err != nil {
		return nil, err
	}

	emitArguments(&spec, bindings)

	emitInputArtifacts(&template, inputs, locations)

	spec.Templates = []v1alpha1.Template{template}
	spec.Entrypoint = template.Name

	wf.Spec = spec
	return &wf, nil
}
