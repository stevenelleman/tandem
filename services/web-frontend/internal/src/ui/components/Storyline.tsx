import React from 'react';

type PropsType = {};
export class Storyline extends React.Component<PropsType> {
  render() {
    return <div className="storyline-container">
      <div className="storyline-intro-image" /*style={{backgroundImage: `url(grouphouse-background.jpeg)`}}*//>
      {/* Change starts at home. */}
      <div className="storyline-section">
        <div className="storyline-section-title">
          <a id="problem">
            Problem
          </a>
        </div>
        <div className="storyline-section-content"></div>
        <div className="storyline-section-transition"></div>
      </div>
      <div className="storyline-section">
        <div className="storyline-section-title">
          <a id="approach">
            Approach
          </a>
        </div>
        <div className="storyline-section-content"></div>
        <div className="storyline-section-transition"></div>
      </div>
      <div className="storyline-section">
        <div className="storyline-section-title">
          <a id="values">
            Values
          </a>
        </div>
        <div className="storyline-section-content"></div>
        <div className="storyline-section-transition"></div>
      </div>
      <div className="storyline-section">
        <div className="storyline-section-title">
          <a id="code">
            Code
          </a>
        </div>
        <div className="storyline-section-content"></div>
        <div className="storyline-section-transition"></div>
      </div>
      <div className="storyline-section">
        <div className="storyline-section-title">
          <a id="community">
            Community
          </a>
        </div>
        <div className="storyline-section-content"></div>
        <div className="storyline-section-transition"></div>
      </div>
    </div>;
  }
}