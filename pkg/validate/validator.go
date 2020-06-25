package validate

type Validatable interface {
	Validate() error
}

func Validate(v interface{}) error {
	if va := v.(Validatable); va != nil {
		return va.Validate()
	}
	return nil
}
