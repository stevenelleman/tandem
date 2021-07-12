import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';

// External Library Imports
import { Client } from 'client';

// Import env var
import { EnvVars } from './env-vars';

// Import all styles
import './styles/App.css';

// Import fonts
import './fonts/Fonts.css';

import { Topbar } from './ui/containers/Topbar';
// @ts-ignore
import { Footer } from './ui/containers/Footer';
import { Storyline } from './ui/components/Storyline';

type StateType = {client: Client};
class App extends React.Component<any, StateType> {
  constructor(props: any) {
    super(props);
    const host = window.location.hostname;
    const { protocol } = EnvVars;
    this.state = {
      client: new Client(protocol, host, null),
    };
  }

  render() {
    // Pass client to all routes/components that need to make API requests
    // const { client } = this.state;
    return (
      <div className="app">
        <Topbar />
        <Router>
          <Switch>
            <Route path="/">
              <Storyline />
            </Route>
          </Switch>
        </Router>
        <Footer />
      </div>
    );
  }
}

export default App;
