package pkg

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = map[int]*User{} // not thread safe for simplicity!
	seq   = 1
)
