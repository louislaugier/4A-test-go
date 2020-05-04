package content

// Video export
type Video struct {
	Platform    string
	URL         string
	Title       string
	DurationMs  int
	Description string
	Category    string
	Tags        []string
	Author      *User
}
