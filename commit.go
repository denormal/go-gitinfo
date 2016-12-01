package gitinfo

// Commit represents a git commit hash.
type Commit interface {
	// String returns the full commit hash.
	String() string

	// Prefix returns the first n characters of the commit hash.
	Prefix(n int) string
}

type commit struct {
	commit string
}

func newCommit(hash string) Commit {
	return &commit{commit: hash}
} // newCommit()

func (c *commit) String() string { return c.commit }

func (c *commit) Prefix(n int) string {
	// ensure len is sane
	if n < 0 {
		return ""
	} else if n > len(c.commit) {
		return c.commit
	} else {
		return c.commit[0:n]
	}
} // Prefix()

// ensure commit implements the Commit interface
var _ Commit = &commit{}
