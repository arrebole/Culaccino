import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import Home from "./views/Home";
import Articles from "./views/Articles";
import Admin from "./views/Admin";
import Edit from "./views/Edit";
import Creator from "./views/Creator";
import Login from "./views/Login";

export default function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/" render={() => <Home />}></Route>
        <Route exact path="/login" render={() => <Login />} />
        <Route exact path="/articles/:article" render={() => <Articles />} />

        <Route exact path="/admin" render={() => <Admin />} />
        <Route exact path="/new" render={() => <Creator />} />
        <Route exact path="/articles/:article/edit" render={() => <Edit />} />

        <Route path="*" render={() => <Redirect to="/" />} ></Route>
      </Switch>
    </BrowserRouter>
  );
}
