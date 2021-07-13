import React, { ReactNode } from 'react';

type PropsType = {
  children: ReactNode,
  width: number
};
export class Workspace extends React.Component<PropsType> {
  render() {
    const { width } = this.props;
    return <div className="workspace" style={{ width }}>
      {this.props.children}
    </div>;
  }
}
