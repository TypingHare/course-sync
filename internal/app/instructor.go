package app

func GetInstructorFiles(projectDir string) []string {
	dataDir := GetDataDir(projectDir)
	srcDir := GetSrcDir(projectDir)

	return []string{
		GetDocsDir(projectDir),
		GetPrototypeWorkspaceDir(srcDir),
		GetDocDataFile(dataDir),
		GetAssignmentDataFile(dataDir),
		GetGradeDataFile(dataDir),
		GetInstructorPrivateKeyFile(dataDir),
	}
}
