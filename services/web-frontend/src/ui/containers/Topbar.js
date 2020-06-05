import React from 'react';

export class Topbar extends React.Component {
  render() {
    const { height } = this.props;
    return (
      <div className="topbar" style={{ height }}>
        <div className="searchbar">
          <div className="searchbar-text">Search...</div>
        </div>
      </div>
    );
  }
}
