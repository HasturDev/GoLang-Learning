package statistics

type code struct {
	file_something string
	code_something string
}

func (c code) Search() string {
	return c.file_something + "has this code" + c.code_something
}

const GitDirName = ".git"
