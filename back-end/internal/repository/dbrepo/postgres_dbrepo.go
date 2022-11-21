package dbrepo

import (
	"back-end/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies(genre ...int) ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id=%d)", genre[0])
	}

	query := fmt.Sprintf(`select id, title, description, year, release_date,rating, runtime, mpaa_rating,
				created_at, updated_at from movies %s order by title`, where)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// get the genres
		genreQuery := `select
    		mg.id, mg.movie_id, mg.genre_id, g.genre_name
    		from
    		    movies_genres mg
    		    left join genres g on (g.id=mg.genre_id)
    		    where 
    		        mg.movie_id=$1`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.ID)

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg models.MovieGenre
			err := genreRows.Scan(
				&mg.ID,
				&mg.MovieID,
				&mg.GenreID,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[mg.ID] = mg.Genre.GenreName
		}
		genreRows.Close()
		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}
	return movies, nil
}
func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, email, first_name, last_name, password, 
			created_at, updated_at from users where email=$1`
	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LatName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (m *PostgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `select id, email, first_name, last_name, password, 
			created_at, updated_at from users where id=$1`
	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LatName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
