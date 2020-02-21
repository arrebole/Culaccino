import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import Home from './views/Home';
import Paper from './views/Paper';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/"><Home /></Route>
        <Route path="/code/:name"><Paper /></Route>
        <Route path="/alive/:name"><Paper /></Route>
        <Route path="/spirit/:name"><Paper /></Route>
        <Route path="/zero/:name"><Paper /></Route>
        <Route path="*"><Redirect to={{ pathname: "/" }}/></Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
