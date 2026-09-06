package plumbing

type TagMode int

const (
	InvalidTagMode TagMode = iota

	TagFollowing

	AllTags

	NoTags
)
