import React from 'react';

export class Workspace extends React.Component {
  render() {
    return <div className="workspace" style={{width: this.props.width}}/>;
  }
}
