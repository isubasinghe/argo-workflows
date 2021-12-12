package transpiler

type CWLFormat interface{}

type CWLNull struct{}
type CWLBool struct{}
type CWLInt struct{}
type CWLLong struct{}
type CWLFloat struct{}
type CWLDouble struct{}
type CWLString struct{}
type CWLFile struct{}
type CWLDirectory struct{}

type CWLType interface{}

type CWLStdin struct{}

type CWLClass interface {
}

type CWLExpression struct {
	Pattern  CWLExpressionPattern
	Required bool
}
type CWLExpressionPattern interface {
}

type CWLDefinition struct {
	Class   CWLClass
	Version string
	Inputs  int
	Outputs int
}
