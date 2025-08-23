package todoapp

type TodoLists struct {
	Id          int `json:"id"`
	Title       int `json:"title"`
	Description int `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int  `json:"id"`
	Title       int  `json:"title"`
	Description int  `json:"description"`
	Done        bool `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
