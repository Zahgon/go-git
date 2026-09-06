//go:build ignore

package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-git/go-billy/v6/osfs"

	"github.com/go-git/go-git/v6/backend"
	"github.com/go-git/go-git/v6/plumbing/transport"
)

func main() {
	log.SetFlags(0)

	addr := flag.String("addr", ":8080", "listen address")
	root := flag.String("root", ".", "root directory containing bare git repositories")
	allowAnon := flag.Bool("allow-anonymous-receive", true, "inject a fake Authorization header for receive-pack (useful for manual testing)")
	verbose := flag.Bool("verbose", true, "log incoming HTTP requests (very useful for debugging push/clone issues)")
	flag.Parse()

	if len(flag.Args()) > 0 {
		*root = flag.Arg(0)
	}

	loader := transport.NewFilesystemLoader(osfs.New(*root), false)
	b := backend.New(loader)

	b.ErrorLog = log.New(os.Stderr, "[backend] ", log.LstdFlags|log.Lmicroseconds)

	var handler http.Handler = b
	if *allowAnon {
		handler = allowAnonymousReceive(handler)
	}
	if *verbose {
		handler = requestLogger(handler)
	}

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen %s: %v", *addr, err)
	}

	srv := &http.Server{
		Handler: handler,
	}

	go func() {
		log.Printf("go-git http server listening on %s", ln.Addr())
		log.Printf("  root directory: %s", *root)
		log.Printf("")
		log.Printf("  Examples:")
		log.Printf("    git clone http://localhost%s/myrepo.git", *addr)
		log.Printf("    git push   http://localhost%s/myrepo.git main   # or master, depending on your local branch")
		log.Printf("")
		log.Printf("  Common issues:")
		log.Printf("    - 'src refspec master does not match any'  →  your local branch is probably called 'main', not 'master'.")
		log.Printf("      Try:  git branch -M main && git push -u origin main")
		log.Printf("    - 401 on push  →  use credentials in the URL or rely on -allow-anonymous-receive")
		log.Printf("")
		log.Printf("  Flags that help debugging:")
		log.Printf("    -verbose=true   (request logging - already on by default)")
		log.Printf("    -allow-anonymous-receive=true")
		log.Printf("")
		if *allowAnon {
			log.Printf("  -allow-anonymous-receive is ON (fake Authorization header injected for receive-pack).")
		}
		log.Printf("  Supports Git-Protocol: version=2 (and classic v0/v1) automatically.")
		log.Printf("  To verify v2 + debug fetch: look for 'v2 fetch:' lines in THIS server console (new diagnostic logs).")
		log.Printf("  Important: after 'ready' you should see 'packfile' immediately in the pkt stream (no 0000 flush between them).")
		log.Printf("")
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down...")
	_ = srv.Close()
	_ = ln.Close()
}

type allowAnonymous struct {
	http.Handler
}

func (a allowAnonymous) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func allowAnonymousReceive(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func requestLogger(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }
