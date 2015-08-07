package db

import (
	"mijstter/model"
)

// posts テーブルのデータをすべて削除します。
func (d *Database) DeleteAllPosts() error {
	_, err := d.db.Exec("delete from posts;")
	if err != nil {
		return err
	}

	return nil
}

// 指定された id のユーザーを検索します。
func (d *Database) ReadPost(id model.ID) (*model.Post, error) {
	stmt, err := d.db.Prepare("select id, user_id, user_name, message, url from posts where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	post := &model.Post{}
	err = stmt.QueryRow(id).Scan(&post.Id, &post.UserId, &post.UserName, &post.Message, &post.Url)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// ユーザーの登録を行います。
func (d *Database) WritePost(post *model.Post) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	if post.Id > 0 {
		// 更新
		stmt, err := tx.Prepare("update posts set user_id=?, user_name=?, message=?, url=? where id=?")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(post.UserId, post.UserName, post.Message, post.Url, post.Id)
		if err != nil {
			return err
		}
	} else {
		// 新規
		stmt, err := tx.Prepare("insert into posts (user_id, user_name, message, url) values (?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.Exec(post.UserId, post.UserName, post.Message, post.Url)
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		post.Id = model.ID(id)
	}
	tx.Commit()

	return nil
}

// 指定された件数のユーザーを検索します。
func (d *Database) ReadPosts(limit int) ([]*model.Post, error) {
	stmt, err := d.db.Prepare("select id, user_id, user_name, message, url from posts order by id limit ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]*model.Post, 0, 20)
	for rows.Next() {
		post := &model.Post{}
		rows.Scan(&post.Id, &post.UserId, &post.UserName, &post.Message, &post.Url)
		posts = append(posts, post)
	}

	return posts, nil
}
