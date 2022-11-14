import React, {Fragment} from "react";
import {Link} from 'react-router-dom';
import {useEffect, useState} from "react";

export default function Genres(){
    const [genres, setGenres] = useState([])
    const[error, setError]=useState(null)
    useEffect(()=> {
        fetch("http://localhost:4000/v1/genres")
            .then((response)=>{
                if (response.status !==200){
                    setError("invalid response code: " + response.status);
                }else{
                    setError(null)
                }
                return response.json();
            })
            .then((json)=> {
                setGenres(json.genres)
            })
            },[]);
    if (error != null) {
        console.log(error)
        return <div>Error: {error.message}</div>
    } else {
        return (
            <Fragment>
                <h2>Genres</h2>
                <div className="list-group">

                {genres.map((m) => (

                        <Link to={`/genre/${m.id}`} key={m.id} className="list-group-item list-group-item-action" state={{genreName:m.genre_name}}>{m.genre_name}</Link>

                    ))}
                </div>
            </Fragment>
        )
    }
};