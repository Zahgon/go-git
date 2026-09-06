package git

type Status map[string]*FileStatus

func (s Status) File(path string) *FileStatus { _ = "STUB: not implemented"; return nil }

func (s Status) IsUntracked(path string) bool { _ = "STUB: not implemented"; return false }

func (s Status) IsClean() bool { _ = "STUB: not implemented"; return false }

func (s Status) String() string { _ = "STUB: not implemented"; return "" }

type FileStatus struct {
	Staging StatusCode

	Worktree StatusCode

	Extra string
}

type StatusCode byte

const (
	Unmodified         StatusCode = ' '
	Untracked          StatusCode = '?'
	Modified           StatusCode = 'M'
	Added              StatusCode = 'A'
	Deleted            StatusCode = 'D'
	Renamed            StatusCode = 'R'
	Copied             StatusCode = 'C'
	UpdatedButUnmerged StatusCode = 'U'
)

type StatusStrategy int

const (
	defaultStatusStrategy = Empty

	Empty StatusStrategy = 0

	Preload StatusStrategy = 1
)

func (s StatusStrategy) new(w *Worktree) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func preloadStatus(w *Worktree) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}
