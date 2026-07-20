package noder

type Path []Noder

func (p Path) Skip() bool { _ = "STUB: not implemented"; return false }

func (p Path) String() string { _ = "STUB: not implemented"; return "" }

func (p Path) Last() Noder { _ = "STUB: not implemented"; return *new(Noder) }

func (p Path) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (p Path) Name() string { _ = "STUB: not implemented"; return "" }

func (p Path) IsDir() bool { _ = "STUB: not implemented"; return false }

func (p Path) Children() ([]Noder, error) { _ = "STUB: not implemented"; return nil, nil }

func (p Path) NumChildren() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p Path) Compare(other Path) int { _ = "STUB: not implemented"; return 0 }
