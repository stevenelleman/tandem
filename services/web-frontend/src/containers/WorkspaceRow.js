import React from "react";
import {Sidebar} from "../components/reuseable/Sidebar";
import {Workspace} from "./Workspace";

export class WorkspaceRow extends React.Component {
  constructor(props) {
    super(props);

    // Need to bind method to pass them down to child component
    this.changeWidth = this.changeWidth.bind(this);

    const defaultWidth = 120;
    this.state = {
      left: defaultWidth,
      right: defaultWidth,
    }
  }

  // Used for updating the component dimensions on resize events
  componentDidMount() {
    window.addEventListener('resize', () => this.forceUpdate())
  }

  // Used for changing width of sidebars
  changeWidth(name, w) {
    var state = {};
    state[name] = w;
    this.setState(state)
  }

  render() {
    var width = window.innerWidth - (this.state.left + this.state.right);
    return <div className="row">
      <Sidebar left type={"silos"} width={this.state.left} changeWidth={this.changeWidth}/>
      <Workspace width={width}/>
      <Sidebar type={"forums"} width={this.state.right} changeWidth={this.changeWidth}/>
    </div>
  }
}
