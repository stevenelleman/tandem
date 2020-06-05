import React from 'react';

// Import all styles
import './styles/App.css';
import './styles/Sidebar.css';
import './styles/Topbar.css';

import { Topbar } from './containers/Topbar';
import { WorkspaceRow } from './containers/WorkspaceRow';

function App() {
  const defaultHeight = 50;
  return (
    <div className="app">
      <Topbar height={defaultHeight} />
      <WorkspaceRow />
    </div>
  );
}

export default App;
