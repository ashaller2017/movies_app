import React, { Fragment} from "react";
import {Link, useLocation, useParams} from "react-router-dom";
import {useEffect, useState} from "react";

export default function OneGenre(){
    const {id} = useParams();
    const [movies, setMovies]=useState([]);
    const [error, setError]=useState(null);
    const [isLoaded, setLoaded]=useState(false);

    const {genreName}=useLocation().state;

    useEffect(() => {
        fetch("http://localhost:4000/v1/movies/" + id)
            .then((response) => {
                if (response.status !== 200) {
                    setError("invalid response code: " + response.status);
                }else {
                    setError(null)
                }
                return response.json();
            })
            .then((json) => {
                setMovies(json.movies);
                setLoaded(true);
                });
            }, [id]);

    //figure out to make empty array
        if (error !== null) {
            return <div>error: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>Genre: {genreName}</h2>
                    <div className="list-group">
                        {movies.map((m)=>(
                            <Link to={`/movies/${m.id}`} key={m.id} className="list-group-item list-group-item-action">{m.title}</Link>
                        ))}
                    </div>

                </Fragment>
            );
        }
}