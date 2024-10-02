package models

import "database/sql"

type Like struct {
	Id        int64         `db:"id"`
	UserId    int64         `db:"user_id"`
	PostId    sql.NullInt64 `db:"post_id"`
	CommentId sql.NullInt64 `db:"comment_id"`
}
