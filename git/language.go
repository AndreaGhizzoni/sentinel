package git

type Language struct {
	Name         string            `json:"name"`
	Repositories map[string]string `json:"repositories"`
}
