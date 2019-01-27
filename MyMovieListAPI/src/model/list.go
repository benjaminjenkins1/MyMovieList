package model

//import "errors"

type listItem struct {
	Id 				 int 		`json:"id"`
	Title 		 string `json:"title"`
	PosterPath string `json:"posterPath"`
	Status     string `json:"status"`
	Runtime    int		`json:"runtime"`
	Overview   string `json:"overview"`
}

type list struct {
	Id     int				`json:"id"`
	Owner  int				`json:"owner"`
	Public bool				`json:"public"`
	Items  []listItem `json:"items"`
}

/*
func NewList(id int) (list, error) {
	
}

func ReadList(id int) (list, error) {

}

func (l *list) WriteList() error {

}
*/