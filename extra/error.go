package extra

const (
	ERROR_REGEX = "Invalid format: "
	KEY_ERROR   = "Key Error: "
	INDEX_ERROR = "Index Error: "
)

type errorGoFormat struct {
	ErrorType string
	Args      string
}

func (e *errorGoFormat) Error() string {
	return e.ErrorType + e.Args
}

func InvalidFormatError(args string) error {
	return &errorGoFormat{
		ErrorType: ERROR_REGEX,
		Args:      args,
	}
}

func KeyError(args string) error {
	return &errorGoFormat{
		ErrorType: KEY_ERROR,
		Args:      args,
	}
}

func IndexError(args string) error {
	return &errorGoFormat{
		ErrorType: INDEX_ERROR,
		Args:      args,
	}
}
