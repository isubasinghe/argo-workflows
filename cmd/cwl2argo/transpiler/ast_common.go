package transpiler

type (
	String         string
	Bool           bool
	Int            int
	Float          float32
	Strings        []string
	SecondaryFiles []CWLSecondaryFileSchema
)

type CWLFormatKind int32

const (
	FormatStringKind CWLFormatKind = iota
	FormatStringsKind
	FormatExpressionKind
)

type CWLFormat struct {
	Kind       CWLFormatKind
	String     String
	Strings    Strings
	Expression CWLExpression
}

type CWLStdin struct{}

type CWLType interface {
	isCWLType()
}

type LoadListingKind int32

const (
	ShallowListingKind LoadListingKind = iota
	DeepListingKind
	NoListingKind
)

type LoadListingEnum struct {
	Kind LoadListingKind
}

type CWLExpressionKind int32

const (
	RawKind CWLExpressionKind = iota
	ExpressionKind
	BoolKind
	IntKind
	FloatKind
)

type CWLExpression struct {
	Kind       CWLExpressionKind
	Raw        string
	Expression string
	Bool       bool
	Int        int
	Float      float64
}

type CWLSecondaryFileSchema struct {
	Pattern  CWLExpression `yaml:"pattern"`
	Required CWLExpression `yaml:"required"`
}
