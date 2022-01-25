package transpiler

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

type CWLFile struct {
	Class          string     `yaml:"class"` // constant value File
	Location       *string    `yaml:"location"`
	Path           *string    `yaml:"path"`
	Dirname        *string    `yaml:"dirname"`
	Nameroot       *string    `yaml:"nameroot"`
	Nameext        *string    `yaml:"nameext"`
	Checksum       *string    `yaml:"checksum"`
	Size           *int64     `yaml:"size"`
	SecondaryFiles []CWLFile  `yaml:"secondaryFiles"`
	Format         *CWLFormat `yaml:"format"`
	Contents       *string    `yaml:"contents"`
}

type CWLInputEntry struct {
	Kind       Type
	FileData   *CWLFile
	BoolData   *bool
	StringData *string
	IntData    *int
}

func (cwlInputEntry *CWLInputEntry) UnmarshalYAML(value *yaml.Node) error {

	var b bool
	err := value.Decode(&b)
	if err == nil {
		cwlInputEntry.Kind = CWLBoolKind
		cwlInputEntry.BoolData = &b
		return nil
	}

	var i int
	err = value.Decode(&i)
	if err == nil {
		cwlInputEntry.Kind = CWLIntKind
		cwlInputEntry.IntData = &i
		return nil
	}

	var s string
	err = value.Decode(&s)
	if err == nil {
		cwlInputEntry.Kind = CWLStringKind
		cwlInputEntry.StringData = &s
		return nil
	}

	var file CWLFile
	err = value.Decode(&file)
	if err == nil {
		if file.Class != "File" {
			return fmt.Errorf("%s was received instead of ", file.Class)
		}
		cwlInputEntry.Kind = CWLFileKind
		cwlInputEntry.FileData = &file
		return nil
	}

	return errors.New("unable to convert into CWLInputEntry")
}
