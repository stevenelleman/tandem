import React from 'react';

// Import all styles
import './styles/App.css';
import './styles/Sidebar.css';
import './styles/Topbar.css';

import { Client } from './client';
import { Topbar } from './ui/containers/Topbar';
import { WorkspaceRow } from './ui/containers/WorkspaceRow';

type StateType = {client: Client};
class App extends React.Component<any, StateType> {
  constructor(props: any) {
    super(props);
    // TODO: Is this the best place for the client to be?
    this.state = {
      client: new Client('localhost:8000'),
    };
  }

  render() {
    const { client } = this.state;
    return (
      <div className="app">
        <Topbar />
        <WorkspaceRow client={client} />
      </div>
    );
  }
}

export default App;
