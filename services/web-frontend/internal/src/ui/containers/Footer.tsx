import React from 'react';

export class Footer extends React.Component {
  render() {
    return (
      <div className="footer-container">
        <div className="footer-content">
          <div className="footer-text">
            Grouphouse v0.0.1
          </div>
          <div className="footer-links">
            <div className="footer-link">
              <a href="/#">
                <img src="/images/logos/instagram_logo.png" alt="Instagram Logo" height="20px" width="20px"></img>
              </a>
            </div>
            <div className="footer-link">
              <a href="/#">
                <img src="/images/logos/facebook_logo.png" alt="Facebook Logo" height="20px" width="20px"></img>
              </a>
            </div>
            <div className="footer-link">
              <a href="/#">
                <img src="/images/logos/twitter_logo.png" alt="Twitter Logo" height="20px" width="20px"></img>
              </a>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
