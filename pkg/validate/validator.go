package validate

type Validatable interface {
	Validate() error
}

func Validate(v interface{}) error {
	if va, ok := v.(Validatable); ok {
		return va.Validate()
	}
	return nil
}
