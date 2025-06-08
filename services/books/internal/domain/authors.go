package domain

type (
	Author struct {
		AuthorID  string `json:"author_id"`
		Name      string `json:"name"`
		Biography string `json:"biography"`
	}

	AuthorRequest struct {
		Name      string `json:"name" valid:"required,minstringlength(3)"`
		Biography string `json:"biography"`
	}
)
