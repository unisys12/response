package annotations

type Validator struct {
	Create string
	Update string
	Both   string
}

func (Validator) Name() string {
	return "Validator"
}
