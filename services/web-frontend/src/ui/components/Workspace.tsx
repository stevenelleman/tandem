import React from 'react';

type PropsType = {width: number};
export class Workspace extends React.Component<PropsType> {
  render() {
    const { width } = this.props;
    return <div className="workspace" style={{ width }} />;
  }
}
