package statistics

import "errors"

var (
	ErrWorktreeNotClean     = errors.New("worktree is not clean")
	ErrSubmoduleNotFound    = errors.New("submodule not found")
	ErrUnstagedChanges      = errors.New("worktree contains unstaged changes")
	ErrGitModulesSymlink    = errors.New(gitmodulesFile + " is a symlink")
	ErrNonFastForwardUpdate = errors.New("non-fast-forward update")
)

var (
	ErrMissingName    = errors.New("name field is required")
	ErrMissingTagger  = errors.New("tagger field is required")
	ErrMissingMessage = errors.New("message field is required")
)

var (
	NoErrAlreadyUpToDate     = errors.New("already up-to-date")
	ErrDeleteRefNotSupported = errors.New("server does not support delete-refs")
	ErrForceNeeded           = errors.New("some refs were not updated")
	ErrExactSHA1NotSupported = errors.New("server does not support exact SHA1 refspec")
	ErrEmptyUrls             = errors.New("URLs cannot be empty")
)

var (
	// ErrBranchExists an error stating the specified branch already exists
	ErrBranchExists = errors.New("branch already exists")
	// ErrBranchNotFound an error stating the specified branch does not exist
	ErrBranchNotFound = errors.New("branch not found")
	// ErrTagExists an error stating the specified tag already exists
	ErrTagExists = errors.New("tag already exists")
	// ErrTagNotFound an error stating the specified tag does not exist
	ErrTagNotFound = errors.New("tag not found")
	// ErrFetching is returned when the packfile could not be downloaded
	ErrFetching = errors.New("unable to fetch packfile")

	ErrInvalidReference          = errors.New("invalid reference, should be a tag or a branch")
	ErrRepositoryNotExists       = errors.New("repository does not exist")
	ErrRepositoryIncomplete      = errors.New("repository's commondir path does not exist")
	ErrRepositoryAlreadyExists   = errors.New("repository already exists")
	ErrRemoteNotFound            = errors.New("remote not found")
	ErrRemoteExists              = errors.New("remote already exists")
	ErrAnonymousRemoteName       = errors.New("anonymous remote name must be 'anonymous'")
	ErrWorktreeNotProvided       = errors.New("worktree should be provided")
	ErrIsBareRepository          = errors.New("worktree not available in a bare repository")
	ErrUnableToResolveCommit     = errors.New("unable to resolve commit")
	ErrPackedObjectsNotSupported = errors.New("packed objects not supported")
	ErrSHA256NotSupported        = errors.New("go-git was not compiled with SHA256 support")
)
