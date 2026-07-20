package capability

type List struct {
	m    map[string]*entry
	sort []string
}

type entry struct {
	Name   string
	Values []string
}

func (l *List) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func DecodeList(raw []byte, l *List) { _ = "STUB: not implemented"; return }

func EncodeList(l *List) []byte { _ = "STUB: not implemented"; return nil }

func (l *List) Get(capability string) []string { _ = "STUB: not implemented"; return nil }

func (l *List) Set(capability string, values ...string) { _ = "STUB: not implemented"; return }

func (l *List) init() {
	if l.m == nil {
		l.m = make(map[string]*entry)
	}
}

func (l *List) Add(c string, values ...string) { _ = "STUB: not implemented"; return }

func (l *List) Supports(capability string) bool { _ = "STUB: not implemented"; return false }

func (l *List) Delete(capability string) { _ = "STUB: not implemented"; return }

func (l *List) All() []string { _ = "STUB: not implemented"; return nil }

func (l *List) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *List) AppendText(b []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *List) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l *List) String() string { _ = "STUB: not implemented"; return "" }
