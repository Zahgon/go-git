package index

func match(pattern, name string) (matched bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func scanChunk(pattern string) (star bool, chunk, rest string) {
	_ = "STUB: not implemented"
	return false, "", ""
}

func matchChunk(chunk, s string) (rest string, ok bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func getEsc(chunk string) (r rune, nchunk string, err error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}
