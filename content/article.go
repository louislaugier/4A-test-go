package content

// Article export
type Article struct {
	Platform string
	URL      string
	Title    string
	Content  string
	Category string
	Tags     []string
	Author   *User
}
