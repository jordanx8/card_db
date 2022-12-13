import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import { ApolloProvider, ApolloClient, InMemoryCache } from "@apollo/client";
import CardTable from './components/CardTable';
import Navigation from './components/Navigation';
import NotFound from './components/NotFound';
import 'bootstrap/dist/css/bootstrap.min.css';
import AddPlayerForm from './components/EditComponents/AddPlayerForm';
import AddSeasonForm from './components/EditComponents/AddSeasonForm';
import AddCardForm from './components/EditComponents/AddCardForm';

const client = new ApolloClient({
  uri: "http://localhost:4000/",
  cache: new InMemoryCache()
});

const router = createBrowserRouter([
  {
    path: "/",
    element: <Navigation />,
    errorElement: <NotFound />,
    children: [
      {
        path: "/",
        element: <CardTable tableName={"New Orleans Pelicans"} />
      },
      {
        path: "/tigers",
        element: <CardTable tableName={"LSU Tigers"} />
      },
      {
        path: "/add/player",
        element: <AddPlayerForm />
      },
      {
        path: "/add/season",
        element: <AddSeasonForm />
      },
      {
        path: "/add/card",
        element: <AddCardForm />
      },
    ]
  },
]);

const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <RouterProvider router={router} />
    </ApolloProvider>
  </React.StrictMode>
);
