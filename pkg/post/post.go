package post

type Post struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Created  string `json:"created"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Text     string `json:"text"`
	URL      string `json:"url"`

	Vote             []*Vote `json:"votes"`
	Score            int     `json:"score"`
	UpvotePercentage int     `json:"upvotePercentage"`

	Comment       []*Comment `json:"comments"`
	lastCommentID int        `json:"-"`

	Views int `json:"views"`

	Author Author `json:"author"`
}

type Author struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}
