package config

import (
	"io"
	"strings"
)

type Encoder struct {
	w io.Writer
}

var (
	subsectionReplacer = strings.NewReplacer(`"`, `\"`, `\`, `\\`)
	valueReplacer      = strings.NewReplacer(`"`, `\"`, `\`, `\\`, "\n", `\n`, "\t", `\t`, "\b", `\b`)
)

func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

func (e *Encoder) Encode(cfg *Config) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeSection(s *Section) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) encodeSubsection(sectionName string, s *Subsection) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Encoder) encodeOptions(opts Options) error { _ = "STUB: not implemented"; return nil }

func (e *Encoder) printf(msg string, args ...any) error { _ = "STUB: not implemented"; return nil }
