import React from 'react';

// Import all styles
import './styles/App.css';
import './styles/Sidebar.css';
import './styles/Topbar.css';

import { Client } from './client/client';
import { Topbar } from './ui/containers/Topbar';
import { WorkspaceRow } from './ui/containers/WorkspaceRow';

function App() {
  // Initialize client.
  // TODO: Is this the best place for the client to be?
  const client = Client("localhost:8000");
  const defaultHeight = 50;
  return (
    <div className="app">
      <Topbar height={defaultHeight} />
      <WorkspaceRow client={client} />
    </div>
  );
}

export default App;
