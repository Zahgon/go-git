package transport

const (
	UploadPackService    = "git-upload-pack"
	UploadArchiveService = "git-upload-archive"
	ReceivePackService   = "git-receive-pack"
)

func ServiceName(service string) string { _ = "STUB: not implemented"; return "" }
