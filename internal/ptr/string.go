package ptr

func String(p *string) string {
	if p == nil {
		return ""
	}

	return *p
}

func StringP(p string) *string {
	if p == "" {
		return nil
	}

	return &p
}
