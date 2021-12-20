package transpiler

import (
	"errors"
	"fmt"
)

func errorNilRequirements(id *string) error {
	if id != nil {
		return fmt.Errorf("Requirements cannot be nil in %s", *id)
	} else {
		return errors.New("Requiremnets cannot be nil")
	}
}

func errorDockerRequirement(id *string) error {
	if id != nil {
		return fmt.Errorf("DockerRequirement must be present in all Argo CWL definitions, %s does not satisfy this", *id)
	} else {
		return errors.New("DockerRequirement must be present in all Argo CWL definitions")
	}
}

func TypeCheckCommandlineInputs(clins []CommandlineInputParameter) error {
	for _, clin := range clins {
		// type check secondary files
		if clin.SecondaryFiles != nil {
		}
	}
	return nil
}

func TypeCheckCommandlineOutputs(clouts []CommandlineOutputParameter) error {

	return nil
}

func TypeCheckCommandlineClass(id *string, class string) error {
	if class == "CommandlineTool" {
		return nil
	}
	if id != nil {
		return fmt.Errorf("\"CommandlineTool\" required but %s was provided in %s", class, *id)
	} else {
		return fmt.Errorf("\"CommandlineTool\" required but %s provided", class)
	}
}

func TypeCheckCommandlineRequirements(id *string, clrs []CWLRequirements) error {
	if clrs == nil {
		return errorNilRequirements(id)
	}

	foundDocker := false

	for _, requirement := range clrs {
		if _, ok := requirement.(DockerRequirement); ok == true {
			foundDocker = true
			break
		}
	}

	if foundDocker == false {
		return errorDockerRequirement(id)
	}
	return nil
}

func TypeCheckCommandlineHints(id *string, hints []interface{}) error {

	return nil
}

func TypeCheckCommandlineTool(cl CommandlineTool) error {
	var err error

	err = TypeCheckCommandlineInputs(cl.Inputs)
	if err != nil {
		return err
	}

	err = TypeCheckCommandlineOutputs(cl.Outputs)
	if err != nil {
		return err
	}

	err = TypeCheckCommandlineClass(cl.Id, cl.Class)
	if err != nil {
		return err
	}

	err = TypeCheckCommandlineRequirements(cl.Id, cl.Requirements)
	if err != nil {
		return err
	}

	return nil
}
