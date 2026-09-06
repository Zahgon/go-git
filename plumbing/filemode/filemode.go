package filemode

import (
	"os"
)

type FileMode uint32

const (
	Empty FileMode = 0

	Dir FileMode = 0o040000

	Regular FileMode = 0o100644

	Deprecated FileMode = 0o100664

	Executable FileMode = 0o100755

	Symlink FileMode = 0o120000

	Submodule FileMode = 0o160000
)

func New(s string) (FileMode, error) { _ = "STUB: not implemented"; return *new(FileMode), nil }

func FromBytes(b []byte) (FileMode, error) { _ = "STUB: not implemented"; return *new(FileMode), nil }

func NewFromOSFileMode(m os.FileMode) (FileMode, error) {
	_ = "STUB: not implemented"
	return *new(FileMode), nil
}

func isSetCharDevice(m os.FileMode) bool { _ = "STUB: not implemented"; return false }

func isSetTemporary(m os.FileMode) bool { _ = "STUB: not implemented"; return false }

func isSetUserExecutable(m os.FileMode) bool { _ = "STUB: not implemented"; return false }

func isSetSymLink(m os.FileMode) bool { _ = "STUB: not implemented"; return false }

func (m FileMode) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (m FileMode) IsMalformed() bool { _ = "STUB: not implemented"; return false }

func (m FileMode) String() string { _ = "STUB: not implemented"; return "" }

func (m FileMode) IsRegular() bool { _ = "STUB: not implemented"; return false }

func (m FileMode) IsFile() bool { _ = "STUB: not implemented"; return false }

func (m FileMode) ToOSFileMode() (os.FileMode, error) {
	_ = "STUB: not implemented"
	return *new(os.FileMode), nil
}
