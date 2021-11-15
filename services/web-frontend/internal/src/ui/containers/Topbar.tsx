import React from 'react';

export class Topbar extends React.Component {
  render() {
    return (
      <div className="topbar">
        <div className="header-logo">
          <img src="/images/grouphouse-logo.png" alt="Mushroom Header"className="header-logo-img"/>
        </div>
        <a href="/" className="header-logo-title">Grouphouse</a>
        <div className="header-links">
          <div className="header-link">
            <a href="#problem">Problem</a>
          </div>
          <div className="header-link">
            <a href="#mission">Mission</a>
          </div>
          <div className="header-link">
            <a href="#process">Process</a>
          </div>
          <div className="header-link">
            <a href="#methods">Methods</a>
          </div>
          <div className="header-link header-emphasis">
            <a href="#community">Join the Community</a>
          </div>
        </div>
          <a href="/login" id="login-icon">
              <img src="/images/logos/user_icon.png" alt="Login Icon" id="login-icon-header"/>
          </a>
      </div>
    );
  }
}
