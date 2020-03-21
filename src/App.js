import React from 'react';
import './App.css';
import MainPage from './Components/MainPage'
import NewEmployee from './Components/NewEmployee'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
} from "react-router-dom";

function App() {
  return (
     <Router>
       <Switch>
         <Route path="/new-account">
           <NewEmployee />
         </Route>
         <Route path="/">
           <MainPage />
         </Route>
       </Switch>
   </Router>
 );
}

export default App;
