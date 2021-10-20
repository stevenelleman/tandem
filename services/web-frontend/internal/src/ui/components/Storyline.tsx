import React from 'react';

type PropsType = {};
export class Storyline extends React.Component<PropsType> {
  render() {
    return (
      <div className="storyline-container">
        <div className="transition-block"></div>
        <div className="storyline-section-dark">
          <div className="storyline-section-title">
            <a id="problem" href="/">
              Problem
            </a>
          </div>
          <div className="storyline-section-hook">
            American community is disappearing.
          </div>
          <img src="/images/grouphouse-entrance-nighttime.jpg" className="storyline-image-left"alt="Grouphouse Entrance Nighttime"></img>
          <div className="storyline-section-text-right">
              There is a shared but divided struggle of feeling along, unsupported, unsafe.
              Lack of community has forced the stakes to increase in romantic relationships, marriage, and family, placing extraordinary and unhealthy levels of stress on these relations.
          </div>
          
        </div>
        <div className="transition-skew-right-flush-top"></div>
        <div className="storyline-section-light">
          <div className="storyline-section-title">
            <a id="our-mission" href="/">
              Mission
            </a>
          </div>
          <div className="storyline-section-hook">
            Our mission is to rebuild safety.
          </div>
          <img src="/images/grouphouse-flower.jpg" className="storyline-image-right"alt="Pink Flower"></img>
          <div className="storyline-section-text-left">
              We all need safety. We all need a home. Our mission is to rebuild community through intentional group-living, so that you can find your people and your home.
          </div>
        </div>
        <div className="transition-skew-left-flush-bottom"></div>
        <div className="storyline-section-dark">
          <div className="storyline-section-title">
            <a id="process" href="/">
              Process
            </a>
          </div>
          <div className="storyline-section-hook">
            Build processes that fulfill needs and bring out the best in us.
          </div>
          <img src="/images/grouphouse-livingroom-with-person.jpg" className="storyline-image-right"alt="Living Room"></img>
          <div className="storyline-section-text-left">
              We're stuck in a terrible loop: needs to unfulfilled, depravation brings out the worst in each of us, "responsibility" falls on the individual,
              blame and isolation compound the initial depravation. So many people are stuck in this loop, trapped as their worst selves and feeling it’s their fault. Individual responsibility is scapegoating. The burden of responsibility falls on process, ensuring circumstances and systems exist to bring out the best in all of us. 
              Safety begets process, process begets safety.
          </div> 
        </div>
        <div className="transition-skew-left-flush-top"></div>
        <div className="storyline-section-light">
          <div className="storyline-section-title">
            <a id="methods" href="/">
              Methods
            </a>
          </div>
        </div>
        <div className="transition-skew-left-flush-bottom"></div>
        <div className="storyline-section-dark">
          <div className="storyline-section-title">
            <a id="community" href="/">
              Community
            </a>          
          </div> 
        </div>
        <div className="transition-block"></div>
      </div>
    );
  }
}
