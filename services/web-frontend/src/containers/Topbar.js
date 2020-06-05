import React from 'react';

export class Topbar extends React.Component {
  render() {
    return <div className="topbar" style={{height: this.props.height}}>
        <div className="searchbar">
          <div className="searchbar-text">Search...</div>
        </div>
      </div>;
  }
}