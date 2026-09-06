package git

import (
	"time"

	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
)

type BlameResult struct {
	Path string

	Rev plumbing.Hash

	Lines []*Line
}

func Blame(c *object.Commit, path string) (*BlameResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Line struct {
	Author string

	AuthorName string

	Text string

	Date time.Time

	Hash plumbing.Hash
}

func newLine(author, authorName, text string, date time.Time, hash plumbing.Hash) *Line {
	_ = "STUB: not implemented"
	return nil
}

func newLines(contents []string, commits []*object.Commit) []*Line {
	_ = "STUB: not implemented"
	return nil
}

type blame struct {
	path string

	fRev *object.Commit

	lineToCommit []*object.Commit

	q *priorityQueue
}

type lineMap struct {
	Orig, Cur    int
	Commit       *object.Commit
	FromParentNo int
}

func (b *blame) addBlames(curItems []*queueItem) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func finishNeeds(curItem *queueItem) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func applyNeeds(child *queueItem, needsMap []lineMap, identicalToChild bool, parentNo int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b BlameResult) String() string { _ = "STUB: not implemented"; return "" }

func (b BlameResult) maxAuthorLength() int { _ = "STUB: not implemented"; return 0 }

type childToNeedsMap struct {
	Child            *queueItem
	NeedsMap         []lineMap
	IdenticalToChild bool
	ParentNo         int
}

type queueItem struct {
	Child                   *queueItem
	MergedChildren          []childToNeedsMap
	Commit                  *object.Commit
	path                    string
	Contents                string
	NeedsMap                []lineMap
	numParentsNeedResolving int
	IdenticalToChild        bool
	ParentNo                int
}

type priorityQueueImp []*queueItem

func (pq *priorityQueueImp) Len() int           { _ = "STUB: not implemented"; return 0 }
func (pq *priorityQueueImp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq *priorityQueueImp) Swap(i, j int) { _ = "STUB: not implemented"; return }
func (pq *priorityQueueImp) Push(x any)    { _ = "STUB: not implemented"; return }
func (pq *priorityQueueImp) Pop() any      { _ = "STUB: not implemented"; return *new(any) }

func (pq *priorityQueueImp) Peek() *object.Commit { _ = "STUB: not implemented"; return nil }

type priorityQueue priorityQueueImp

func (pq *priorityQueue) Init()             { _ = "STUB: not implemented"; return }
func (pq *priorityQueue) Len() int          { _ = "STUB: not implemented"; return 0 }
func (pq *priorityQueue) Push(c *queueItem) { _ = "STUB: not implemented"; return }

func (pq *priorityQueue) Pop() *queueItem { _ = "STUB: not implemented"; return nil }

func (pq *priorityQueue) Peek() *object.Commit { _ = "STUB: not implemented"; return nil }

type parentCommit struct {
	Commit *object.Commit
	Path   string
}

func parentsContainingPath(path string, c *object.Commit) ([]parentCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func blobHash(path string, commit *object.Commit) (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}
