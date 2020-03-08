import React from 'react';
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import HomePage from './views/Home';
import PaperPage from './views/Paper';
import EditorPage from './views/Editor';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/"><HomePage /></Route>
        <Route path="/editor"><EditorPage /></Route>
        <Route path="/:name"><PaperPage /></Route>
        <Route path="*"><Redirect to={{ pathname: "/" }}/></Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
