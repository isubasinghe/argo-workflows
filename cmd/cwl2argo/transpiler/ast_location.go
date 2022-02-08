package transpiler

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

type FileLocations struct {
	Inputs  map[string]FileLocationData `json:"inputs"`
	Outputs map[string]FileLocationData `json:"outputs"`
}

type FileLocationKind string

const (
	HTTPKind FileLocationKind = "http"
	S3Kind   FileLocationKind = "s3"
	GITKind  FileLocationKind = "git"
)

type FileLocationData struct {
	Name string                 `json:"name"`
	Type FileLocationKind       `json:"type"`
	HTTP *v1alpha1.HTTPArtifact `json:"http"`
	S3   *v1alpha1.S3Artifact   `json:"s3"`
	HDFS *v1alpha1.HDFSArtifact `json:"hdfs"`
}

func (f *FileLocationData) UnmarshalJSON(data []byte) error {
	type rawFileLocationData FileLocationData

	tmp := rawFileLocationData{}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	switch tmp.Type {
	case HTTPKind:
		if tmp.HTTP == nil {
			return errors.New("http data not provided")
		}
	case S3Kind:
		if tmp.S3 == nil {
			return errors.New("s3 data not provided")
		}
	default:
		return fmt.Errorf("%s is not a valid type", tmp.Type)
	}
	f.Name = tmp.Name
	f.Type = tmp.Type
	f.HTTP = tmp.HTTP
	f.S3 = tmp.S3
	return nil
}
