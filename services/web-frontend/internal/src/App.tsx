import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';

// Local Library Imports
import { Client } from 'client';

// Import env var
import { EnvVars } from './env-vars';

// Import all styles
import './styles/App.css';

// Import fonts
import './fonts/Fonts.css';

import { Topbar } from './ui/containers/Topbar';
// @ts-ignore
import { Storyline } from './ui/components/Storyline';
import { Footer } from './ui/containers/Footer';
import { Login } from './auth/Login'
import { Registration } from './auth/Registration'
 
type StateType = {client: Client};
class App extends React.Component<unknown, StateType> {
  constructor(props: any) {
    super(props);
    const host: string = window.location.hostname;
    const protocol: string = EnvVars.protocol;
    this.state = {
      client: new Client(protocol, host, null),
    };
  }

  render() {
    // Pass client to all routes/components that need to make API requests
    // const { client } = this.state;
    return (
      <div className="app">
        <Router>
          <Switch>                
            <Topbar />
              <Route exact path="/">
                <Storyline />
              </Route>
              <Route path="/login">
                <Login />
              </Route>
              <Route path="/register">
                <Registration />
              </Route>
          </Switch>
        </Router>
        <Footer />
      </div>
    );
  }
}

export default App;
