import React from 'react';

export class Workspace extends React.Component {
  render() {
    const { width } = this.props;
    return <div className="workspace" style={{ width }} />;
  }
}
