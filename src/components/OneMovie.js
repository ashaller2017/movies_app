import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";


function OneMovie(props) {
    const [movie, setMovie]=useState({})
    const {id} = useParams();

    useEffect(() => {
        setMovie({
            id: `${id}`,
            title: "Some Movie",
            runtime: 150,
    });
}, []);

        return(
            <>
            <h2>
                Movie: {movie.title} - {movie.id}
            </h2>
            <table className="table table-compact table-striped">
                <thead></thead>
                <tbody>
                <tr>
                    <td><strong>Title:</strong></td>
                    <td>{movie.title}</td>
                </tr>
                <tr>
                    <td><strong>Run Time:</strong></td>
                    <td>{movie.runtime} minutes</td>
                </tr>
                </tbody>
            </table>
            </>
        );
}
export default OneMovie;