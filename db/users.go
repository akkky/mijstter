package db

import (
	"mijstter/model"
)

// users テーブルのデータをすべて削除します。
func (d *Database) DeleteAllUsers() error {
	_, err := d.db.Exec("delete from users;")
	if err != nil {
		return err
	}

	return nil
}

// 指定された id のユーザーを検索します。
func (d *Database) ReadUser(id model.ID) (*model.User, error) {
	stmt, err := d.db.Prepare("select id, user_name, password_hash from users where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &model.User{}
	err = stmt.QueryRow(id).Scan(&user.Id, &user.UserName, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 指定された user_name のユーザーを検索します。
func (d *Database) IsUserNameExist(user_name string) (bool, error) {
	stmt, err := d.db.Prepare("select count(id) from users where user_name=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(user_name).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

// 指定された user_name のユーザーを検索します。
func (d *Database) ReadUserByUserName(user_name string) (*model.User, error) {
	stmt, err := d.db.Prepare("select id, user_name, password_hash from users where user_name=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &model.User{}
	err = stmt.QueryRow(user_name).Scan(&user.Id, &user.UserName, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ユーザーの登録を行います。
func (d *Database) WriteUser(user *model.User) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	if user.Id > 0 {
		// 更新
		stmt, err := tx.Prepare("update users set user_name=?, password_hash=? where id=?")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.UserName, user.PasswordHash, user.Id)
		if err != nil {
			return err
		}
	} else {
		// 新規
		stmt, err := tx.Prepare("insert into users (user_name, password_hash) values (?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.Exec(user.UserName, user.PasswordHash)
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		user.Id = model.ID(id)
	}
	tx.Commit()

	return nil
}

// 指定された件数のユーザーを検索します。
func (d *Database) ReadUsers(limit int) ([]*model.User, error) {
	stmt, err := d.db.Prepare("select id, user_name, password_hash from users order by id limit ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0, 20)
	for rows.Next() {
		user := &model.User{}
		rows.Scan(&user.Id, &user.UserName, &user.PasswordHash)
		users = append(users, user)
	}

	return users, nil
}
