import React from 'react';

export class Footer extends React.Component {
  render() {
    return (
      <div className="footer-container">
        <div className="footer-links">
          <div className="footer-link">
            <a href="">Link 1</a>
          </div>
          <div className="footer-link">
            <a href="">Link 2</a>
          </div>
          <div className="footer-link">
            <a href="">Link 3</a>
          </div>
        </div>
      </div>
    );
  }
}
