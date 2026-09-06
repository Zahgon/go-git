package packfile

import (
	"bytes"

	"github.com/go-git/go-git/v6/plumbing"
)

type Version uint32

const (
	V2 Version = 2
)

func (v Version) Supported() bool { _ = "STUB: not implemented"; return false }

type ObjectHeader struct {
	Type            plumbing.ObjectType
	Offset          int64
	ContentOffset   int64
	Size            int64
	Reference       plumbing.Hash
	OffsetReference int64
	Crc32           uint32
	Hash            plumbing.Hash

	content     *bytes.Buffer
	parent      *ObjectHeader
	diskType    plumbing.ObjectType
	externalRef bool

	chainDepth int
}

func (oh *ObjectHeader) ID() plumbing.Hash { _ = "STUB: not implemented"; return *new(plumbing.Hash) }

type SectionType int

const (
	HeaderSection SectionType = iota
	ObjectSection
	FooterSection
)

type Header struct {
	Version    Version
	ObjectsQty uint32
}

type PackData struct {
	Section      SectionType
	header       Header
	objectHeader ObjectHeader
	checksum     plumbing.Hash
}

func (p PackData) Value() any { _ = "STUB: not implemented"; return *new(any) }
