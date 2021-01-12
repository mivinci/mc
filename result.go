package mc

func String(r *Result) string {
	if r.Error != "" {
		return string(r.Error)
	}
	return string(r.Value)
}
