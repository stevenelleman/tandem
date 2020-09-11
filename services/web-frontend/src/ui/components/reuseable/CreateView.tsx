import React from 'react';
import { Client } from '../../../client';

// General component for creating objects -- would love to generalize this as much as possible to avoid one-off components

type PropsType = {
  client: Client,
  type: string,
};
export class CreateView extends React.Component<PropsType> {
  render() {
    const { type } = this.props;

    return <div>
      {` ${type} woot woot!`}
    </div>;
  }
}
