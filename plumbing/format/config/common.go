package config

func New() *Config { _ = "STUB: not implemented"; return nil }

type Config struct {
	Comment  *Comment
	Sections Sections
	Includes Includes
}

type Includes []*Include

type Include struct {
	Path   string
	Config *Config
}

type Comment string

const (
	NoSubsection = ""
)

func (c *Config) Section(name string) *Section { _ = "STUB: not implemented"; return nil }

func (c *Config) HasSection(name string) bool { _ = "STUB: not implemented"; return false }

func (c *Config) RemoveSection(name string) *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) RemoveSubsection(section, subsection string) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) AddOption(section, subsection, key, value string) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) SetOption(section, subsection, key, value string) *Config {
	_ = "STUB: not implemented"
	return nil
}
