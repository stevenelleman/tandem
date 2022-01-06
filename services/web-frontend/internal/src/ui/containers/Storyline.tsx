import React from 'react';

type PropsType = {};
export class Storyline extends React.Component<PropsType> {
  render() {
    return (
      <div className="storyline-container">
        <div className="image-container">
          <div className="image-wrapper">
            <div className="storyline-intro-image" />
          </div>
          <div className="storyline-intro-image-text storyline-intro-centered">
            Change starts at home.
          </div>
          <div className="storyline-intro-image-text-blur storyline-intro-centered">
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
            <img src="/images/grouphouse-entrance-nighttime.jpg" className="storyline-image-left" alt="Gloom"></img>
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
          <img src="/images/grouphouse-flower.jpg" className="storyline-image-right"alt="Hope"></img>
          <div className="storyline-section-text-left">
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
          <img src="/images/grouphouse-porch-daytime.jpg" className="storyline-image-left" alt="Porch"></img>
          <div className="storyline-section-text-right">
            Group-living requires skin in the game.
            It's hard. But you get out what you put in. We crowdsource tools to make it easier.
          </div>
        </div>

        <div className="community-section-transition"></div>
        <div className="storyline-section-grey">
          <div className='image-container'>
            <div className='community-image-wrapper'>
              <div className='community-image-triangle-right triangle-1-right'></div>
              <div className='community-image-triangle-right triangle-2-right'></div>
              <div className='community-image-triangle-right triangle-3-right'></div>
              <div className='community-image-triangle-right triangle-4-right'></div>
              <div className='community-image-triangle-right triangle-5-right'></div>
              <div className='community-image-triangle-right triangle-6-right'></div>
              <div className='community-image-triangle-right triangle-7-right'></div>
              <div className='community-image-triangle-right triangle-8-right'></div>
              <div className='community-image-triangle-right triangle-9-right'></div>
              <div className='community-image-triangle-right triangle-10-right'></div>
              <div className='community-image-triangle-right triangle-11-right'></div>
              <div className='community-image-triangle-right triangle-12-right'></div>
              <div className='community-image-triangle-right triangle-13-right'></div>
              <div className='community-image-triangle-left triangle-1-left'></div>
              <div className='community-image-triangle-left triangle-2-left'></div>
              <div className='community-image-triangle-left triangle-3-left'></div>
              <div className='community-image-triangle-left triangle-4-left'></div>
              <div className='community-image-triangle-left triangle-5-left'></div>
              <div className='community-image-triangle-left triangle-6-left'></div>
              <div className='community-image-triangle-left triangle-7-left'></div>
              <div className='community-image-triangle-left triangle-8-left'></div>
              <div className='community-image-triangle-left triangle-9-left'></div>
              <div className='community-image-triangle-left triangle-10-left'></div>
              <div className='community-image-triangle-left triangle-11-left'></div>
              <div className='community-image-triangle-left triangle-12-left'></div>
              <div className='community-image-triangle-left triangle-13-left'></div>
              <div className='community-image'></div>
            </div>
            <div className='community-centered community-image-text'>
              <div className="community-section-title">
                <a id="community">Community</a>
              </div>
              <div className='community-section-hook'>The smallest steps are the bravest.</div>
              Join the <a className="embedded-link" href={"https://discord.gg/SejPVdUnM3"}>server</a>,
              contribute to the <a className="embedded-link" href={"https://github.com/grouphouse-io/grouphouse"}>code</a>,
              fork the <a className="embedded-link" href={"https://github.com/stevenelleman/tandem"}>parent</a>.
            </div>
            <div className='community-centered community-image-text-blur'></div>
          </div>
          <div className="storyline-section-text-right">

          </div>
        </div>
      </div>
    );
  }
}
