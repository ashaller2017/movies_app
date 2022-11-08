import React, {Fragment} from 'react';
import * as ReactDOM from "react-dom";
import {BrowserRouter as Router, Link, Routes, Route, Outlet, useParams, useRouteMatch} from "react-router-dom";
import Movies from './components/Movies';
import Home from './components/Home';
import Admin from './components/Admin';
import Categories from './components/Categories';
import OneMovie from './components/OneMovie';
export default function App() {
  return (
      <Router>
  <div className="container">
    <div className="row">
      <h1 className="mt-3">
        Go Watch a Movie!
        <hr className="mb-3"></hr>
      </h1>
    </div>
    <div className="row">
      <div className="col-md-2">
        <nav>
          <ul className="list-group">
            <li className="list-group-item">
              <Link to="/">Home</Link>
            </li>
            <li className="list-group-item">
              <Link to="/movies">Movies</Link>
            </li>
            <li className="list-group-item">
              <Link to="/by-category">Categories</Link>
            </li>
            <li className="list-group-item">
              <Link to="/admin">Manage Catalogue</Link>

            </li>
          </ul>
        </nav>

      </div>

      <div className="col-md-10">
        <Routes>
          <Route path="/movies" element={ <Outlet />}>
            <Route index element={<Movies />}/>
            <Route path=":id" element={ <OneMovie />}/>
          </Route>
          <Route path="/by-category" element={ <Outlet />}>
            <Route index element={<CategoryPage/>}/>
            <Route path=":category" element={<Categories title={`title`}/>}/>
          </Route>
          <Route path="/admin" element={<Admin />}/>
          <Route path="/" element={ <Home />}/>
        </Routes>

      </div>
    </div>
  </div>
      </Router>
  );
}

function CategoryPage(){

  const categories=[
    {route: "comedy", title: "Comedy"},
    {route: "drama", title: "Drama"},
  ];
  return (
      <div>
        <h2>Categories</h2>
        <ul>
          {categories.map((cat)=>(
              <li>
                <Link to={cat.route} state={{title: cat.title}}>
                  {cat.title}
                </Link>
              </li>
          ))}
        </ul>
      </div>
  );
}