import React from 'react';

type PropsType = {};
export class Storyline extends React.Component<PropsType> {
  render() {
    return (
      <div className="storyline-container">
        <div className="intro-image-container">
          <div className="image-wrapper">
            <div className="storyline-intro-image" />
          </div>
          <div className="storyline-intro-image-text">
            Change starts at home.
          </div>
        </div>
        <div className="transition-block"></div>
        <div className="storyline-section-dark">
          <div className="storyline-section-title">
            <a id="problem">
              Problem
            </a>
          </div>
          <div className="storyline-section-hook">
            American community is disappearing.
          </div>
            <img src="/images/grouphouse-entrance-nighttime.jpg" className="storyline-image-left" alt="Grouphouse Entrance Nighttime"></img>
            <div className="storyline-section-text-right">
              There is a shared but divided struggle of feeling unsupported, unsafe, alone.
              Lack of community has raised the stakes in romantic relationships, marriage, and family, placing extraordinary and unhealthy levels of stress on these relations.
            </div>
        </div>
        <div className="transition-skew-right-flush-top"></div>
        <div className="storyline-section-light">
          <div className="storyline-section-title">
            <a id="mission">
              Mission
            </a>
          </div>
          <div className="storyline-section-hook">
            Rebuild safety by rebuilding community.
          </div>
          <img src="/images/grouphouse-flower.jpg" className="storyline-image-right"alt="Pink Flower"></img>
          <div className="storyline-section-text-right">
            Home means safety. Our mission is to rebuild community through intentional group-living, so that you can find your people and your home.
          </div>
        </div>
        <div className="transition-skew-left-flush-bottom"></div>
        <div className="storyline-section-dark">
          <div className="storyline-section-title">
            <a id="method">
              Method
            </a>
          </div>
          <div className="storyline-section-hook">
            We build tools for better living.
          </div>
          <img src="/images/grouphouse-livingroom-with-person.jpg" className="storyline-image-left"alt="Living Room"></img>
          <div className="storyline-section-text-right">
            Group-living requires skin in the game.
            It's hard. But you get out what you put in. We crowdsource tools to make it easier.
          </div>
        </div>
        <div className="transition-skew-left-flush-top"></div>
        <div className="storyline-section-light">
          <div className="storyline-section-title">
            <a id="community">
              Community
            </a>
          </div>
          <div className="storyline-section-hook">
            The smallest steps are the bravest.
          </div>
          <div className="storyline-section-text-right">
            Join the <a className="embedded-link" href={"https://discord.gg/SejPVdUnM3"}>server</a>,
            contribute to the <a className="embedded-link" href={"https://github.com/grouphouse-io/grouphouse"}>code</a>,
            fork the <a className="embedded-link" href={"https://github.com/stevenelleman/tandem"}>parent</a>.
          </div>
        </div>
      </div>
    );
  }
}
