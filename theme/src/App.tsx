import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import Home from './views/Home';
import Paper from './views/Paper';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/"><Home /></Route>
        <Route path="/:name"><Paper /></Route>
        <Route path="*"><Redirect to={{ pathname: "/" }}/></Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
