package git

import (
	"errors"
	"regexp"
	"time"

	"github.com/go-git/go-billy/v6"

	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/client"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp"
	"github.com/go-git/go-git/v6/plumbing/protocol/packp/sideband"
)

type SubmoduleRecursivity uint

const (
	DefaultRemoteName = "origin"

	NoRecurseSubmodules SubmoduleRecursivity = 0

	DefaultSubmoduleRecursionDepth SubmoduleRecursivity = 10
)

var ErrMissingURL = errors.New("URL field is required")

type CloneOptions struct {
	URL string

	ClientOptions []client.Option

	RemoteName string

	ReferenceName plumbing.ReferenceName

	SingleBranch bool

	Mirror bool

	NoCheckout bool

	Depth int

	RecurseSubmodules SubmoduleRecursivity

	ShallowSubmodules bool

	Progress sideband.Progress

	Tags plumbing.TagMode

	Shared bool

	Filter packp.Filter

	Bare bool

	AllowEmptyRepo bool

	worktree billy.Filesystem
}

type MergeOptions struct {
	Strategy MergeStrategy
}

type MergeStrategy int8

const (
	FastForwardMerge MergeStrategy = iota
)

type OrtMergeStrategyOption int8

const (
	TheirsMergeStrategy OrtMergeStrategyOption = iota

	OursMergeStrategy
)

func (o *CloneOptions) Validate() error { _ = "STUB: not implemented"; return nil }

type PullOptions struct {
	RemoteName string

	RemoteURL string

	ReferenceName plumbing.ReferenceName

	SingleBranch bool

	Depth int

	ClientOptions []client.Option

	RecurseSubmodules SubmoduleRecursivity

	Progress sideband.Progress

	Force bool
}

func (o *PullOptions) Validate() error { _ = "STUB: not implemented"; return nil }

const (
	InvalidTagMode = plumbing.InvalidTagMode

	TagFollowing = plumbing.TagFollowing

	AllTags = plumbing.AllTags

	NoTags = plumbing.NoTags
)

type FetchOptions struct {
	RemoteName string

	RemoteURL string
	RefSpecs  []config.RefSpec

	Depth int

	ClientOptions []client.Option

	Progress sideband.Progress

	Tags plumbing.TagMode

	Force bool

	Prune bool

	Filter packp.Filter
}

func (o *FetchOptions) Validate() error { _ = "STUB: not implemented"; return nil }

type PushOptions struct {
	RemoteName string

	RemoteURL string

	RefSpecs []config.RefSpec

	ClientOptions []client.Option

	Progress sideband.Progress

	Prune bool

	Force bool

	RequireRemoteRefs []config.RefSpec

	FollowTags bool

	ForceWithLease *ForceWithLease

	Options []string

	Atomic bool

	Quiet bool
}

type ForceWithLease struct {
	RefName plumbing.ReferenceName

	Hash plumbing.Hash
}

func (o *PushOptions) Validate() error { _ = "STUB: not implemented"; return nil }

type SubmoduleUpdateOptions struct {
	Init bool

	NoFetch bool

	RecurseSubmodules SubmoduleRecursivity

	ClientOptions []client.Option

	Depth int
}

var (
	ErrBranchHashExclusive  = errors.New("Branch and Hash are mutually exclusive")
	ErrCreateRequiresBranch = errors.New("Branch is mandatory when Create is used")
)

type CheckoutOptions struct {
	Hash plumbing.Hash

	Branch plumbing.ReferenceName

	Create bool

	Force bool

	Keep bool

	SparseCheckoutDirectories []string
}

func (o *CheckoutOptions) Validate() error { _ = "STUB: not implemented"; return nil }

type ResetMode int8

const (
	MixedReset ResetMode = iota

	HardReset

	MergeReset

	SoftReset

	KeepReset
)

type ResetOptions struct {
	Commit plumbing.Hash

	Mode ResetMode

	Files []string

	SparseDirs []string

	SkipSparseDirValidation bool

	fromTree *object.Tree
}

func (o *ResetOptions) Validate(r *Repository) error { _ = "STUB: not implemented"; return nil }

type LogOrder int8

const (
	LogOrderDefault LogOrder = iota
	LogOrderDFS
	LogOrderDFSPost
	LogOrderBSF
	LogOrderCommitterTime
	LogOrderDFSPostFirstParent
)

type LogOptions struct {
	From plumbing.Hash

	To plumbing.Hash

	Order LogOrder

	FileName *string

	PathFilter func(string) bool

	All bool

	Since *time.Time

	Until *time.Time
}

var ErrMissingAuthor = errors.New("author field is required")

type AddOptions struct {
	All bool

	Path string

	Glob string

	SkipStatus bool
}

func (o *AddOptions) Validate(_ *Repository) error { _ = "STUB: not implemented"; return nil }

type CommitOptions struct {
	All bool

	AllowEmptyCommits bool

	Author *object.Signature

	Committer *object.Signature

	Parents []plumbing.Hash

	Signer Signer

	Amend bool
}

func (o *CommitOptions) Validate(r *Repository) error { _ = "STUB: not implemented"; return nil }

func (o *CommitOptions) loadConfigAuthorAndCommitter(r *Repository) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	ErrMissingName    = errors.New("name field is required")
	ErrMissingTagger  = errors.New("tagger field is required")
	ErrMissingMessage = errors.New("message field is required")
)

type CreateTagOptions struct {
	Tagger *object.Signature

	Message string

	Signer Signer
}

func (o *CreateTagOptions) Validate(r *Repository, _ plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *CreateTagOptions) loadConfigTagger(r *Repository) error {
	_ = "STUB: not implemented"
	return nil
}

type ListOptions struct {
	ClientOptions []client.Option

	PeelingOption PeelingOption

	Timeout int
}

type PeelingOption uint8

const (
	IgnorePeeled PeelingOption = 0

	OnlyPeeled PeelingOption = 1

	AppendPeeled PeelingOption = 2
)

type CleanOptions struct {
	Dir bool
}

type GrepOptions struct {
	Patterns []*regexp.Regexp

	InvertMatch bool

	CommitHash plumbing.Hash

	ReferenceName plumbing.ReferenceName

	PathSpecs []*regexp.Regexp
}

var ErrHashOrReference = errors.New("ambiguous options, only one of CommitHash or ReferenceName can be passed")

func (o *GrepOptions) Validate(w *Worktree) error { _ = "STUB: not implemented"; return nil }

func (o *GrepOptions) validate(r *Repository) error { _ = "STUB: not implemented"; return nil }

type PlainOpenOptions struct {
	DetectDotGit bool

	AlternatesFS billy.Filesystem
}

func (o *PlainOpenOptions) Validate() error { _ = "STUB: not implemented"; return nil }

var ErrNoRestorePaths = errors.New("you must specify path(s) to restore")

type RestoreOptions struct {
	Staged bool

	Worktree bool

	Files []string
}

func (o *RestoreOptions) Validate() error { _ = "STUB: not implemented"; return nil }
