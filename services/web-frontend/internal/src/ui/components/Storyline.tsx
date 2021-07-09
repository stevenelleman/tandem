import React from 'react';

type PropsType = {};
export class Storyline extends React.Component<PropsType> {
  render() {
    return (
      <div className="storyline-container">
        <div className="storyline-intro-image">
          <div className="storyline-intro-text"> Change starts at home.</div>
        </div>
        { /* Change starts at home. */ }
        <div className="storyline-section">
          <div className="storyline-section-title">
            <a id="problem" href="/">
              Problem
            </a>
          </div>
          <div className="storyline-section-content">
            <div className="storyline-section-text">
              { 'American community has ' }
              <a href="http://bowlingalone.com/">collapsed</a>
              { '.' }
              <br />
              <br />
              { 'People feel alone, unsupported, unsafe. Lack of community has forced the stakes to increase in "strong ties" \
              of romantic relationships, marriage, and family, \
              placing extraordinary and unhealthy levels of stress on these relations. \
              Social media has filled the void to offer the illusion of community while deepening the sense of isolation \
              and personal dissatisfaction.' }
              <br />
              <br />
              { 'We have all seen the social media cry for help, we have tried to help only to find the problem cannot be \
              addressed with our individual efforts. \
              Instead the last-gasp accelerates the process of abandonment, as friends turn their back to preserve their \
              own finite security and energy.' }
              <br />
              <br />
              { 'The net-effect is that the people who need the help the most are least likely to get it, creating a vicious \
              cycle of abandonment and isolation.' }
              <br />
              <br />
              { 'The vicious cycle is a kind of debt with interest, compounding upon itself, creating the conditions for \
              implosion (depression, obesity, addiction, suicide) and explosion (hatred, anger, bigotry, \
              violence).' }
              <br />
              <br />
              { 'No individual can escape these forces. A community is vital to step in to stop the vicious cycle.\n\n' }
              { 'But where can we rebuild community?' }
              <br />
              <br />
            </div>
          </div>
          <div className="storyline-section-transition" />
        </div>
        <div className="storyline-section">
          <div className="storyline-section-title">
            <a id="our-mission" href="/">
              Our Mission
            </a>
          </div>
          <div className="storyline-section-content">
            <div className="storyline-section-text">
              { 'Home is safety.' }
              <br />
              <br />
              It&apos;s a comforting embrace. It&apos;s breaking bread among friends. A warm bed after a weary day.
              A place to drop one&apos;s burden to come back to it on a new day.
              <br />
              <br />
              { 'We all need that safety. We all need a home.' }
              <br />
              <br />
              { 'At Grouphouse we believe that ' }
              <i>housing</i>
              { ' could become a vehicle of ' }
              <i>homing</i>
              { '.' }
              <br />
              <br />
              { 'Our mission is to rebuild community through intentional group-living, so that you can find ' }
              <i>
                {' '}
                your
                people
                {' '}
              </i>
              {' '}
              and
              {' '}
              <i> your home</i>
              { '. Our mission is to rebuild safety.' }
              <br />
              <br />
            </div>
          </div>
          <div className="storyline-section-transition" />
        </div>
        <div className="storyline-section">
          <div className="storyline-section-title">
            <a id="values" href="/">
              Values
            </a>
          </div>
          <div className="storyline-section-content">
            <div className="storyline-section-text">
              { '1. Change starts at home.' }
              <br />
              The smallest steps are often the most courageous.
              This is where revelations and revolutions truly begin.
              We&apos;re rooting for you.
              <br />
              <br />
              { '2. Bring out the best in people.' }
              <br />
              { 'People are good. People are beautiful. People are phenomenal.' }
              <i> You are beautiful. </i>
              { 'Good systems bring out the best. Bad systems and bad circumstances bring out the worst. \
              The burden of responsibility falls on processes, not people.' }
              <br />
              <br />
              { '3. Build a win-win world.' }
              <br />
              In English there&apos;s a false dichotomy in &quot;selfish&quot; and &quot;selfless&quot;.
              The distinction suggests that needs are selfish and therefore bad.
              The most beautiful things in life are both &quot;selfish&quot; and &quot;selfless&quot;: they&apos;re aligned, they&apos;re barycentric.
              Win-win scenarios
              {' '}
              <i>always</i>
              {' '}
              { 'exist, but it often takes trust and creativity to find them.' }
              <br />
              <br />
              { '4. Honor every story by honoring your own.' }
              <br />
              { 'You are the herald of your own story, and every story is profound.' }
              {' '}
              <i>Knowing</i>
              {' '}
              your story is profound is
              the path to knowing that everyone elses is as well.
              Caring about others&apos; stories starts with caring about your own.
              <br />
              <br />
            </div>
          </div>
          <div className="storyline-section-transition" />
        </div>
        <div className="storyline-section">
          <div className="storyline-section-title">
            <a id="code" href="/">
              Code
            </a>
          </div>
          <div className="storyline-section-content">
            <div className="storyline-section-text">
              We&apos;re taking a different approach.
              <br />
              <br />
              Software isn&apos;t like typical work. You don&apos;t need to rebuild the wheel every time.
              In fact, in an ideal world the only thing you would ever need to build is the logic that&apos;s essential to your
              product. The rest can be shared from others&apos; efforts.
              <br />
              <br />
              { 'In this world a generic ' }
              <a href="https://en.wikipedia.org/wiki/Library_(computing)">library</a>
              { ' can be\
              reused a huge number of times, multiplying the value the original efforts tenfold, hundredfold, even\
              thousandfold!' }
              <br />
              { 'We want to work our way to that reality.' }
              <br />
              <br />
              <a href="https://github.com/stevenelleman/tandem"> Tandem </a>
              is a generic root library for building and deploying web applications that
              {' '}
              <i> anyone </i>
              {' '}
              can
              {' '}
              <a
                href="https://en.wikipedia.org/wiki/Fork_(software_development)"
              >
                fork
              </a>
              {' '}
              and use,
              provided generic code is pushed back for the broader Tandem community to use.
              <br />
              <br />
              <a href="https://github.com/grouphouse-io/grouphouse"> Grouphouse </a>
              {' '}
              is forked from Tandem and adds its
              own specific business logic on top.
              <br />
              <br />
            </div>
          </div>
          <div className="storyline-section-transition" />
        </div>
        <div className="storyline-section">
          <div className="storyline-section-title">
            <a id="community" href="/">
              Community
            </a>
          </div>
          <div className="storyline-section-content">
            <div className="storyline-section-text">
              Join our public Discord server:
              {' '}
              <a href="https://discord.gg/SejPVdUnM3">https://discord.gg/SejPVdUnM3</a>
            </div>
          </div>
          <div className="storyline-section-transition" />
        </div>
      </div>
    );
  }
}
