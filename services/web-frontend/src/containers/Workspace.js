import React from 'react';
import {Sidebar} from "../components/reuseable/Sidebar";

export class Workspace extends React.Component {
  render() {
    const margin = this.props.margin;
    // style={{marginLeft: margin, marginRight: margin}}
    return <div  className="workspace"/>;
  }
}
