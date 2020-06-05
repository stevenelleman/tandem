import React from 'react';
import './styles/App.css';
import './styles/Sidebar.css';
import './styles/Topbar.css';
import { Workspace } from './containers/Workspace';
import {Sidebar} from "./components/reuseable/Sidebar";
import {Topbar} from "./containers/Topbar";

function App() {
  const defaultWidth = 120;
  const defaultHeight = 50;
  return (
    <div className="app">
      <Topbar height={defaultHeight}/>
      <div className="row">
        <Sidebar left={true} type={"silos"} width={defaultWidth}/>
        <Workspace margin={defaultWidth}/>
        <Sidebar type={"forums"} width={defaultWidth}/>
      </div>
    </div>
  );
}

export default App;
