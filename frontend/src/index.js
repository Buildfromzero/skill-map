import React from 'react';
import ReactDOM from 'react-dom/client';
import 'bootstrap/dist/css/bootstrap.css';
import '@floating-ui/react';
import Home from './routes/home';
import UserLogin from './routes/login';
import reportWebVitals from './reportWebVitals';
import ErrorPage from "./routes/error-pages";
import UserPage from './routes/users';

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import SkillPage from './routes/skills';
import UserDetail from './routes/userDetail';


const router = createBrowserRouter([

  {
    path: "/",
    element: <Home />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/users",
    element: <UserPage />,
    errorElement: <ErrorPage />,
  },

  {
    path: "/users/:userId/",
    element: <UserDetail />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/skills",
    element: <SkillPage />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/login",
    element: <UserLogin />,
    errorElement: <ErrorPage />,
  },
]);


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
