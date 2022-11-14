import React, {useEffect, useState, Fragment} from "react";
import {useParams} from "react-router-dom";

export default function OneMovie() {
    const [movie, setMovie]=useState([])
    const [error, setError]=useState(null)
    const {id} = useParams();

    useEffect(() => {
        fetch("http://localhost:4000/v1/movie/" + id)
            .then((response) => {
                if (response.status !== 200) {
                    setError("invalid response code: " + response.status);
                } else {
                    setError(null)
                }
                return response.json();
            })
            .then((json) => {
                setMovie(json.movie)
            })
    },[id]);
    if (movie.genres){
        movie.genres=Object.values(movie.genres)
    }else{
        movie.genres=[]
    }
    if (error!==null){
        return <div>error: {error.message}</div>
    }else {
        return (
            <Fragment>
                <h2>
                    Movie: {movie.title} - ({movie.year})
                </h2>
                <div className="float-start">
                    <small>rating: {movie.mpaa_rating}</small>
                </div>
                <div className="float-end">
                    {movie.genres.map((m, index) =>(
                        <span className="badge bg-secondary me-1">
                            {m}
                        </span>
                        ))}

                </div>
                <div className="clearfix"></div>
                <hr />
                <table className="table table-compact table-striped">
                    <thead></thead>
                    <tbody>
                    <tr>
                        <td><strong>Title:</strong></td>
                        <td>{movie.title}</td>
                    </tr>
                    <tr>
                        <td><strong>Description:</strong></td>
                        <td>{movie.description}</td>
                    </tr>
                    <tr>
                        <td><strong>Run Time:</strong></td>
                        <td>{movie.runtime} minutes</td>
                    </tr>
                    </tbody>
                </table>
            </Fragment>
        );
    }
}
