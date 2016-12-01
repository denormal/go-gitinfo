package gitinfo

type Commit interface {
	String() string
	Prefix(int) string
}

type commit struct {
	commit string
}

func newCommit(hash string) Commit {
	return &commit{commit: hash}
} // newCommit()

func (c *commit) String() string        { return c.commit }
func (c *commit) Prefix(len int) string { return c.commit[0:len] }

// ensure commit implements the Commit interface
var _ Commit = &commit{}
