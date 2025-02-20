package repo

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil

}

func (repository Posts) SearchID(postID uint64) (models.Post, error) {
	line, err := repository.db.Query(`
		select p.*, u.nick from
		posts p inner join users u
		on u.id = p.author_id where p.id = ?
	`, postID)

	if err != nil {
		return models.Post{}, err
	}

	var post models.Post
	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil
}

func (repository Posts) Search(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select distinct p.*, u.nick from posts p
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc
	`, userID, userID)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare(
		"update posts set title = ?, content = ? where id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from posts where id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) SearchByUser (userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select p.*, u.nick from posts p
		join users u on u.id = p.author_id
		where p.author_id = ?
	`, userID)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}

func (repository Posts) Like(postID uint64) error {
	statement, err := repository.db.Prepare(
		"update posts set likes = likes + 1 where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Unlike(postID uint64) error {
	statement, err := repository.db.Prepare(
		`update posts set likes = 
		CASE WHEN likes > 0 THEN likes - 1
		ELSE likes END 
		where id = ?`,
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
