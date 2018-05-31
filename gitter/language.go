package gitter

type Language struct {
	Name         string            `json:"name"`
	Repositories map[string]string `json:"repositories"`
	Command      string            `json:"command"`
}

func (l *Language) BuildFolderStructure(base string) error {
	switch l.Name {
	case "java":
		return l.buildJavaFolderStructure(base)
	case "go":
		return l.buildGoFolderStructure(base)
	default:
		return nil
	}
}

func (l *Language) buildJavaFolderStructure(base string) error {
	var languagePath = base + "/" + l.Name
	if _, err := createFolderIfNotExists(languagePath); err != nil {
		return err
	}
	return nil
}

func (l *Language) buildGoFolderStructure(base string) error {
	dirs := []string{
		base + "/" + l.Name,
		base + "/" + l.Name + "/bin",
		base + "/" + l.Name + "/pkg",
		base + "/" + l.Name + "/src/github.com/AndreaGhizzoni",
	}

	for _, dir := range dirs {
		if _, err := createFolderIfNotExists(dir); err != nil {
			return err
		}
	}
	return nil
}
