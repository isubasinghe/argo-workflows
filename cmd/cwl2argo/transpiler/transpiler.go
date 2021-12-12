package transpiler

import (
	"github.com/kr/pretty"
	log "github.com/sirupsen/logrus"
	"github.com/uc-cdis/mariner/wflib"
)

const (
	CWLVersion = "v1.0"
)

func validateWorkflow(wf *wflib.WorkflowJSON) {

}

type WorkflowErrors struct {
	GeneralErrors []string
	FileErrors    map[string][]string
}

type Validator struct {
	Workflow   *wflib.WorkflowJSON
	mainExists bool
}

func (vl *Validator) validate() {

}

func (vl *Validator) Validate() {
	werr := &WorkflowErrors{
		GeneralErrors: make([]string, 2),
		FileErrors:    make(map[string][]string),
	}
	if vl.Workflow.Graph == nil {
		log.Error("$graph was not found in Workflow")
	}
	if vl.Workflow.CWLVersion != CWLVersion {
		log.Errorf("CWL version not correct, %s accepted but %s received", CWLVersion, vl.Workflow.CWLVersion)
	}
	for _, entry := range *vl.Workflow.Graph {
		if entry["id"] == "#main" {
			pretty.Println(entry)
		}
	}
	_ = werr
}

func ToCWLDocument(wf wflib.WorkflowJSON) {}

func TranspileFile(inputFile string, outputFile string) error {
	data, err := wflib.PackWorkflow(inputFile)
	if err != nil {
		return err
	}
	_ = data
	return nil
}
