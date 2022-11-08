import React, {Component} from 'react';
import {useLocation} from "react-router-dom";

function Categories(){
    const location=useLocation();
    const {title}=location.state;
    return <h2>Category - {title}</h2>
}
export default Categories;