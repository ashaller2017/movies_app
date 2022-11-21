package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}
	movie, err := app.models.DB.Get(id)

	err = app.writeJSON(w, http.StatusOK, movie)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, movies)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	//read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	//validate user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	//check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	//create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LatName,
	}
	//generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)
}
func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			//parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
			}
			// get the user ID from token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
			}
			user, err := app.DB.GetUserByID(userID)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
			}
			u := jwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LatName,
			}
			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating tokens"), http.StatusUnauthorized)
			}
			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)
		}
	}
}
func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}

//func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
//	genres, err := app.DB.GenresAll()
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//	err = app.writeJSON(w, http.StatusOK, genres)
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//}
//func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
//	params := httprouter.ParamsFromContext(r.Context())
//	genreID, err := strconv.Atoi(params.ByName("genre_id"))
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//	movies, err := app.DB.AllMovies(genreID)
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//	err = app.writeJSON(w, http.StatusOK, movies, "movies")
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//}
//func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
//
//}
//func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {
//
//}
//func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {
//
//}
//func (app *application) searchMovie(w http.ResponseWriter, r *http.Request) {
//
//}
