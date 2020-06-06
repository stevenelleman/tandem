import React from 'react';
import { Sidebar } from '../components/reuseable/Sidebar';
import { Workspace } from './Workspace';

/*
  TODO:
    - Need to disp

 */

export class WorkspaceRow extends React.Component {
  constructor(props) {
    super(props);

    // Need to bind method to pass them down to child component
    this.changeWidth = this.changeWidth.bind(this);

    const defaultWidth = 120;
    this.state = {
      left: defaultWidth,
      right: defaultWidth,
    };
  }

  // Used for updating the component dimensions on resize events
  componentDidMount() {
    window.addEventListener('resize', () => this.forceUpdate());
  }

  // Used for changing width of sidebars
  changeWidth(name:string, w:number) {
    const state: any = {};
    state[name] = w;
    this.setState(state);
  }

  render() {
    const { left, right } = this.state;
    const width = window.innerWidth - (left + right);
    return (
      <div className="row">
        <Sidebar left type="silos" width={left} changeWidth={this.changeWidth} />
        <Workspace width={width} />
        <Sidebar type="forums" width={right} changeWidth={this.changeWidth} />
      </div>
    );
  }
}
