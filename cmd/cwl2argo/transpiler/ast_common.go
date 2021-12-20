package transpiler

type String string
type Bool bool
type Int int
type Float float32
type Strings []string

type CWLFormat interface {
	isCWLFormat()
}

func (_ String) isCWLFormat()        {}
func (_ Strings) isCWLFormat()       {}
func (_ CWLExpression) isCWLFormat() {}

type CWLNull struct{}

func (_ CWLNull) isCWLType() {}

type CWLBool struct{}

func (_ CWLBool) isCWLType() {}

type CWLInt struct{}

func (_ CWLInt) isCWLType() {}

type CWLLong struct{}

func (_ CWLLong) isCWLType() {}

type CWLFloat struct{}

func (_ CWLFloat) isCWLType() {}

type CWLDouble struct{}

func (_ CWLDouble) isCWLType() {}

type CWLString struct{}

func (_ CWLString) isCWLType() {}

type CWLFile struct{}

func (_ CWLFile) isCWLType() {}

type CWLDirectory struct{}

func (_ CWLDirectory) isCWLType() {}

type CWLStdin struct{}

type CWLType interface {
	isCWLType()
}

type LoadListingEnum interface {
	isLoadListingEnum()
}
type NoListing struct{}

func (_ NoListing) isLoadListingEnum() {}

type ShallowListing struct{}

func (_ ShallowListing) isLoadListingEnum() {}

type DeepListing struct{}

func (_ DeepListing) isLoadListingEnum() {}

type CWLClass interface {
}

type CWLExpression struct {
	Pattern  CWLExpressionString
	Required CWLExpressionBool
}

type CWLExpressionString interface {
	isCWLExpressionString()
}

func (_ String) isCWLExpressionString()        {}
func (_ CWLExpression) isCWLExpressionString() {}

type CWLExpressionBool interface {
	isCWLExpressionBool()
}

func (_ Bool) isCWLExpressionBool()          {}
func (_ CWLExpression) isCWLExpressionBool() {}

type CWLExpressionInt interface {
	isCWLExpressionInt()
}

func (_ Int) isCWLExpressionInt()           {}
func (_ CWLExpression) isCWLExpressionInt() {}

type CWLExpressionNum interface {
	isCWLExpressionNum()
}

func (_ Int) isCWLExpressionNum()           {}
func (_ Float) isCWLExpressionNum()         {}
func (_ CWLExpression) isCWLExpressionNum() {}

type CWLSecondaryFileSchema struct {
	Pattern  CWLExpressionString
	Required bool
}

type CWLDefinition struct {
	Class   CWLClass
	Version string
	Inputs  int
	Outputs int
}
