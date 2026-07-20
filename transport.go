package git

import (
	_ "github.com/go-git/go-git/v6/plumbing/transport/file"
	_ "github.com/go-git/go-git/v6/plumbing/transport/git"
	_ "github.com/go-git/go-git/v6/plumbing/transport/http"
	_ "github.com/go-git/go-git/v6/plumbing/transport/ssh"
)
