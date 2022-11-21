import React from "react";
import ReactDOM from "react-dom/client";
import App from './App'
import Home from './components/Home'
import Movies from './components/Movies'
import Genres from './components/Genres'
import ErrorPage from "./components/ErrorPage"
import EditMovie from "./components/EditMovie"
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import ManageCatalogue from "./components/ManageCatalogue";
import GraphQL from "./components/GraphQL";
import Login from "./components/Login";
import Movie from "./components/Movie";

const router= createBrowserRouter([
    {
        path: "/",
        element: <App />,
        errorElement: <ErrorPage/>,
        children: [
            {index: true, element: <Home/>},
            {
                path: "/movies",
                element: <Movies/>
            },
            {
                path: "/movies/:id",
                element: <Movie/>
            },
            {
                path: "/genres",
                element: <Genres/>
            },
            {
                path: "/movies",
                element: <Movies/>
            },
            {
                path: "/admin/movie/0",
                element: <EditMovie/>
            },
            {
                path: "/manage-catalogue",
                element: <ManageCatalogue/>
            },
            {
                path: "/graphql",
                element: <GraphQL/>
            },
            {
                path: "/login",
                element: <Login/>
            },
        ]
    }
])

const root= ReactDOM.createRoot(
    document.getElementById("root")
);

root.render(
    <React.StrictMode>
        <RouterProvider router={router}/>
    </React.StrictMode>
)