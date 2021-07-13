import React, { Component, ReactNode } from 'react';
import { Dispatch, Action } from 'redux';
import { connect } from 'react-redux';
import { Sidebar } from '../components/reuseable/Sidebar';
import { Workspace } from '../components/Workspace';
import { fetchSilosIfNeeded } from '../../redux/actions/silos';
import { Client } from 'client';
import { ReceiveSiloActionType, AppDispatch } from '../../types';

// Types
type PropsType = {
  children: ReactNode,
  client: Client,
  fetchSilos: (client: Client) => Dispatch<Action<ReceiveSiloActionType>>
};
type StateType = {left:number, right:number};

class PrivateWorkspaceRow extends Component<PropsType, StateType> {
  constructor(props: PropsType) {
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
    this.initProps();
  }

  initProps() {
    const { fetchSilos, client } = this.props;
    fetchSilos(client);
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
        <Workspace width={width}>
          { this.props.children }
        </Workspace>
        <Sidebar left={false} type="forums" width={right} changeWidth={this.changeWidth} />
      </div>
    );
  }
}

const mapStateToProps = (state: any) => {
  const { silos, forums } = state;
  return {
    silos,
    forums,
  };
};

const mapDispatchToProps = (dispatch: AppDispatch) => ({
  fetchSilos: (client: Client) => dispatch(fetchSilosIfNeeded(client)),
});

export const WorkspaceRow = connect(mapStateToProps, mapDispatchToProps)(PrivateWorkspaceRow);
