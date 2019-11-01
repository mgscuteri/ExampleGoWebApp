import React, { Component } from 'react';
import piLogo from '../../images/piLogoSmall.png';
import ReactDOM from 'react-dom'
import { Route, Link, BrowserRouter as Router } from 'react-router-dom'
import { Switch } from 'react-router-dom';

import AdminActions from '../../pages/AdminActions/AdminActions.js'
import UserActions from '../../pages/UserActions/UserActions.js'
import HomePage from '../../pages/HomePage/HomePage.js'
import LogIn from '../../pages/LogIn/LogIn.js';


class Navbar extends React.Component {  
    constructor(props) {
      super(props);
    }

    render() {
      return (
          <Router>
            <div>
              <div class="navBarContainer">
                <div class="navbar">
                  <ul class="topnav">
                    <li><Link to="/site"><img src={piLogo} alt="logo"/></Link></li>
                    <li class="topnav"><Link to="/site">Home</Link></li>                    
                    <li class="topnav"><Link to="/site/AdminActions">Admin Actions</Link></li>
                    <li class="topnav"><Link to="/site/UserActions">User Actions</Link></li>
                    <li class="topnav, right"><Link to="/site/LogIn">Log In</Link></li>
                  </ul>
                </div>
              </div>
              <div className="appBody">
                <Switch>
                  <Route exact path="/site" component={HomePage} />
                  <Route path="/site/AdminActions" component={AdminActions} />
                  <Route path="/site/UserActions" component={UserActions} />                  
                  <Route path="/site/LogIn" component={LogIn} />                  
                  <Route render={ () => <h1>404 Error</h1> } />
                </Switch>
              </div>
            </div>
          </Router>
      );
    }
  }

export default Navbar