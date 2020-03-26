import React from 'react';
import './App.css';
import MainPage from './Components/MainPage'
import NewAccount from './Components/NewAccount'
import Login from './Components/Login'
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
           <NewAccount />
         </Route>
         <Route path="/login">
           <Login />
         </Route>
         <Route path="/">
           <MainPage />
         </Route>
       </Switch>
   </Router>
 );
}

export default App;
