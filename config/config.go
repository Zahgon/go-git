package config

import (
	"errors"
	"io"

	format "github.com/go-git/go-git/v6/plumbing/format/config"
	"github.com/go-git/go-git/v6/plumbing/protocol"
)

const (
	DefaultFetchRefSpec = "+refs/heads/*:refs/remotes/%s/*"

	DefaultPushRefSpec = "refs/heads/*:refs/heads/*"

	DefaultProtocolVersion = protocol.V2
)

type ConfigStorer interface { //nolint:revive // stutters but is a well-established name
	Config() (*Config, error)
	SetConfig(*Config) error
}

var (
	ErrInvalid = errors.New("config invalid key in remote or branch")

	ErrRemoteConfigNotFound = errors.New("remote config not found")

	ErrRemoteConfigEmptyURL = errors.New("remote config: empty URL")

	ErrRemoteConfigEmptyName = errors.New("remote config: empty name")
)

type Scope int

const (
	LocalScope Scope = iota
	GlobalScope
	SystemScope
)

type Config struct {
	Core struct {
		IsBare bool

		Worktree string

		CommentChar string

		RepositoryFormatVersion format.RepositoryFormatVersion

		AutoCRLF string

		FileMode bool

		HooksPath string

		ProtectNTFS OptBool

		ProtectHFS OptBool
	}

	User user

	Author struct {
		Name string

		Email string
	}

	Committer struct {
		Name string

		Email string
	}

	Tag struct {
		GpgSign OptBool
	}

	Commit struct {
		GpgSign OptBool
	}

	GPG struct {
		Format string

		SSH struct {
			AllowedSignersFile string
		}
	}

	Pack struct {
		Window uint

		ReadReverseIndex bool

		WriteReverseIndex bool
	}

	Index struct {
		SkipHash OptBool
	}

	Init struct {
		DefaultBranch string
	}

	UploadArchive struct {
		AllowUnreachable OptBool
	}

	Extensions struct {
		ObjectFormat format.ObjectFormat

		WorktreeConfig bool
	}

	Protocol struct {
		Version protocol.Version
	}

	Remotes map[string]*RemoteConfig

	Submodules map[string]*Submodule

	Branches map[string]*Branch

	URLs []*URL

	Raw *format.Config
}

type user struct {
	Name string

	Email string

	SigningKey string
}

func Merge(src ...*Config) Config { _ = "STUB: not implemented"; return *new(Config) }

func merge(dst, src any) { _ = "STUB: not implemented"; return }

func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

func ReadConfig(r io.Reader) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadConfig(scope Scope) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func Paths(scope Scope) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

const (
	remoteSection              = "remote"
	submoduleSection           = "submodule"
	branchSection              = "branch"
	coreSection                = "core"
	packSection                = "pack"
	userSection                = "user"
	tagSection                 = "tag"
	commitSection              = "commit"
	authorSection              = "author"
	committerSection           = "committer"
	gpgSection                 = "gpg"
	initSection                = "init"
	urlSection                 = "url"
	extensionsSection          = "extensions"
	protocolSection            = "protocol"
	fetchKey                   = "fetch"
	urlKey                     = "url"
	pushurlKey                 = "pushurl"
	bareKey                    = "bare"
	worktreeKey                = "worktree"
	commentCharKey             = "commentChar"
	windowKey                  = "window"
	readReverseIndexKey        = "readReverseIndex"
	writeReverseIndexKey       = "writeReverseIndex"
	mergeKey                   = "merge"
	rebaseKey                  = "rebase"
	nameKey                    = "name"
	emailKey                   = "email"
	signingKey                 = "signingKey"
	descriptionKey             = "description"
	defaultBranchKey           = "defaultBranch"
	repositoryFormatVersionKey = "repositoryformatversion"
	objectFormatKey            = "objectformat"
	worktreeConfigKey          = "worktreeConfig"
	mirrorKey                  = "mirror"
	promisorKey                = "promisor"
	partialCloneFilterKey      = "partialclonefilter"
	versionKey                 = "version"
	autoCRLFKey                = "autocrlf"
	fileModeKey                = "filemode"
	hooksPathKey               = "hooksPath"
	protectNTFSKey             = "protectNTFS"
	protectHFSKey              = "protectHFS"
	indexSection               = "index"
	skipHashKey                = "skipHash"
	formatKey                  = "format"
	allowedSignersFileKey      = "allowedSignersFile"
	gpgSignKey                 = "gpgSign"
	uploadArchiveSection       = "uploadArchive"
	allowUnreachableKey        = "allowUnreachable"

	DefaultPackWindow = uint(10)

	DefaultFileMode = true
)

func (c *Config) Unmarshal(b []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Config) unmarshalCore() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalExtensions() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalTag() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalCommit() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalUser() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalGPG() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalPack() error { _ = "STUB: not implemented"; return nil }

func (c *Config) unmarshalRemotes() error { _ = "STUB: not implemented"; return nil }

func (c *Config) unmarshalURLs() error { _ = "STUB: not implemented"; return nil }

func unmarshalSubmodules(fc *format.Config, submodules map[string]*Submodule) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) unmarshalBranches() error { _ = "STUB: not implemented"; return nil }

func (c *Config) unmarshalProtocol() error { _ = "STUB: not implemented"; return nil }

func (c *Config) unmarshalIndex() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalInit() { _ = "STUB: not implemented"; return }

func (c *Config) unmarshalUploadArchive() { _ = "STUB: not implemented"; return }

func (c *Config) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Config) marshalCore() { _ = "STUB: not implemented"; return }

func (c *Config) marshalExtensions() { _ = "STUB: not implemented"; return }

func (c *Config) marshalTag() { _ = "STUB: not implemented"; return }

func (c *Config) marshalCommit() { _ = "STUB: not implemented"; return }

func (c *Config) marshalUser() { _ = "STUB: not implemented"; return }

func (c *Config) marshalGPG() { _ = "STUB: not implemented"; return }

func (c *Config) marshalPack() { _ = "STUB: not implemented"; return }

func (c *Config) marshalRemotes() { _ = "STUB: not implemented"; return }

func (c *Config) marshalSubmodules() { _ = "STUB: not implemented"; return }

func (c *Config) marshalBranches() { _ = "STUB: not implemented"; return }

func (c *Config) marshalURLs() { _ = "STUB: not implemented"; return }

func (c *Config) marshalProtocol() { _ = "STUB: not implemented"; return }

func (c *Config) marshalIndex() { _ = "STUB: not implemented"; return }

func (c *Config) marshalInit() { _ = "STUB: not implemented"; return }

func (c *Config) marshalUploadArchive() { _ = "STUB: not implemented"; return }

type RemoteConfig struct {
	Name string

	URLs []string

	Mirror bool

	insteadOfRulesApplied bool

	originalURLs []string

	Fetch []RefSpec

	Promisor bool

	PartialCloneFilter string

	raw *format.Subsection
}

func (c *RemoteConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *RemoteConfig) unmarshal(s *format.Subsection) error { _ = "STUB: not implemented"; return nil }

func (c *RemoteConfig) marshal() *format.Subsection { _ = "STUB: not implemented"; return nil }

func (c *RemoteConfig) IsFirstURLLocal() bool { _ = "STUB: not implemented"; return false }

func (c *RemoteConfig) applyURLRules(urlRules []*URL) { _ = "STUB: not implemented"; return }
