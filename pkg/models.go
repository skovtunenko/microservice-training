package pkg

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//type Users struct {
//	sync.RWMutex
//	data map[int]*User
//	seq  int
//}

var (
	users = map[int]*User{}
	seq   = 1
)
