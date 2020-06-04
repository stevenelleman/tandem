import React from 'react';
import './App.css';
import { Workspace } from './containers/Workspace';
import {Sidebar} from "./components/reuseable/Sidebar";
import {Topbar} from "./containers/Topbar";

function App() {
  const defaultWidth = 120;
  // Apply the same inline display / block pattern in buttons here, want workspace to take the remaining space, just like the title
  return (
    <div className="app">
      <Topbar/>
      <div className="row">
        <Sidebar left={true} type={"silos"} width={defaultWidth}/>
        <Workspace margin={defaultWidth}/>
        <Sidebar type={"forums"} width={defaultWidth}/>
      </div>
    </div>
  );
}

export default App;
