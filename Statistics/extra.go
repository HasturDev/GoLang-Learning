package statistics

type Code struct {
	FileSomething string
	CodeSomething string
}

func Search(c Code) string {
	return c.FileSomething + "has this code" + c.CodeSomething
}

const GitDirName = ".git"
