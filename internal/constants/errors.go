package constants

type (
	RepoErr string
)

func (e RepoErr) Error() string {
	return string(e)
}

const (
	ErrUserAlreadyExists   = RepoErr("user already exists")
	ErrUserNotFound        = RepoErr("user not found")
	ErrUserNotExistingRole = RepoErr("role does not exist")
)
