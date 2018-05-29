package gitter

type Language struct {
	Name         string            `json:"name"`
	Repositories map[string]string `json:"repositories"`
	Command      string            `json:"command"`
}
