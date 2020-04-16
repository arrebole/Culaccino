import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";

// <--
import Home from "./views/Home";
import Articles from "./views/Articles";
// -->

export default function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/"><Home/></Route>
        <Route exact path="/articles/:article"><Articles /></Route>
        <Route path="*"><Redirect to={{ pathname: "/" }} /></Route>
      </Switch>
    </BrowserRouter>
  );
}
