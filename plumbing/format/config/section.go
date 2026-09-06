package config

type Section struct {
	Name        string
	Options     Options
	Subsections Subsections
}

type Subsection struct {
	Name    string
	Options Options
}

type Sections []*Section

func (s Sections) GoString() string { _ = "STUB: not implemented"; return "" }

type Subsections []*Subsection

func (s Subsections) GoString() string { _ = "STUB: not implemented"; return "" }

func (s *Section) IsName(name string) bool { _ = "STUB: not implemented"; return false }

func (s *Section) Subsection(name string) *Subsection { _ = "STUB: not implemented"; return nil }

func (s *Section) HasSubsection(name string) bool { _ = "STUB: not implemented"; return false }

func (s *Section) RemoveSubsection(name string) *Section { _ = "STUB: not implemented"; return nil }

func (s *Section) Option(key string) string { _ = "STUB: not implemented"; return "" }

func (s *Section) OptionAll(key string) []string { _ = "STUB: not implemented"; return nil }

func (s *Section) HasOption(key string) bool { _ = "STUB: not implemented"; return false }

func (s *Section) AddOption(key, value string) *Section { _ = "STUB: not implemented"; return nil }

func (s *Section) SetOption(key, value string) *Section { _ = "STUB: not implemented"; return nil }

func (s *Section) RemoveOption(key string) *Section { _ = "STUB: not implemented"; return nil }

func (s *Subsection) IsName(name string) bool { _ = "STUB: not implemented"; return false }

func (s *Subsection) Option(key string) string { _ = "STUB: not implemented"; return "" }

func (s *Subsection) OptionAll(key string) []string { _ = "STUB: not implemented"; return nil }

func (s *Subsection) HasOption(key string) bool { _ = "STUB: not implemented"; return false }

func (s *Subsection) AddOption(key, value string) *Subsection {
	_ = "STUB: not implemented"
	return nil
}

func (s *Subsection) SetOption(key string, value ...string) *Subsection {
	_ = "STUB: not implemented"
	return nil
}

func (s *Subsection) RemoveOption(key string) *Subsection { _ = "STUB: not implemented"; return nil }
