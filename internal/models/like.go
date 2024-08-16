package models

type Like struct {
	Id 	   int	  `db:"id"`
	UserId string `db:"user_id"`
	PostId string `db:"post_id"`
}