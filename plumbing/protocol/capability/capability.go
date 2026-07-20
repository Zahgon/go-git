package capability

import (
	"errors"
)

var (
	ErrArguments = errors.New("capability does not accept arguments")

	ErrArgumentsRequired = errors.New("capability requires an argument")

	ErrMultipleArguments = errors.New("capability accepts only one argument")

	ErrEmptyArgument = errors.New("capability argument cannot be empty")
)

type Capability = string

const (
	MultiACK Capability = "multi_ack"

	MultiACKDetailed Capability = "multi_ack_detailed"

	NoDone Capability = "no-done"

	ThinPack Capability = "thin-pack"

	NoThin Capability = "no-thin"

	Sideband Capability = "side-band"

	Sideband64k Capability = "side-band-64k"

	OFSDelta Capability = "ofs-delta"

	Agent Capability = "agent"

	Shallow Capability = "shallow"

	DeepenSince Capability = "deepen-since"

	DeepenNot Capability = "deepen-not"

	DeepenRelative Capability = "deepen-relative"

	NoProgress Capability = "no-progress"

	IncludeTag Capability = "include-tag"

	ReportStatus Capability = "report-status"

	ReportStatusV2 Capability = "report-status-v2"

	DeleteRefs Capability = "delete-refs"

	Quiet Capability = "quiet"

	Atomic Capability = "atomic"

	PushOptions Capability = "push-options"

	AllowTipSHA1InWant Capability = "allow-tip-sha1-in-want"

	AllowReachableSHA1InWant Capability = "allow-reachable-sha1-in-want"

	PushCert Capability = "push-cert"

	SymRef Capability = "symref"

	ObjectFormat Capability = "object-format"

	Filter Capability = "filter"
)

const (
	LsRefs Capability = "ls-refs"

	FetchCmd Capability = "fetch"

	ObjectInfo Capability = "object-info"

	BundleURI Capability = "bundle-uri"

	ServerOption Capability = "server-option"

	SessionID Capability = "session-id"

	WaitForDone Capability = "wait-for-done"
)

const userAgent = "go-git/6.x"

func DefaultAgent() string { _ = "STUB: not implemented"; return "" }

func Validate(l *List) error { _ = "STUB: not implemented"; return nil }

func validateCapability(c Capability, values []string) error { _ = "STUB: not implemented"; return nil }

func isKnown(c Capability) bool { _ = "STUB: not implemented"; return false }

func requiresArgument(c Capability) bool { _ = "STUB: not implemented"; return false }

func allowsMultipleArguments(c Capability) bool { _ = "STUB: not implemented"; return false }

func validateNoEmptyArgs(values []string) error { _ = "STUB: not implemented"; return nil }

func validateSessionID(sessionID string) error { _ = "STUB: not implemented"; return nil }
