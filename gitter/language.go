package gitter

type Language struct {
	Name           string            `json:"name"`
	Repositories   map[string]string `json:"repositories"`
	Command        string            `json:"command"`
	ProjectsFolder string            `json:"projects_folder"`
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

	if len(l.ProjectsFolder) == 0 {
		l.ProjectsFolder = languagePath
	} else {
		l.ProjectsFolder = languagePath + "/" + l.ProjectsFolder
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

	// TODO
	// in the case of golang, the default "projects_folder" could be
	// /src/github.com/<username>
	if len(l.ProjectsFolder) == 0 {
		l.ProjectsFolder = dirs[0]
	} else {
		l.ProjectsFolder = dirs[0] + "/" + l.ProjectsFolder
	}
	return nil
}
