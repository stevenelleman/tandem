import React from 'react';
import {shallow, mount} from 'enzyme';
import { Sidebar } from './Sidebar';

test('renders sidebar and checks collapse works', () => {
  const sidebarComponent = <Sidebar left={true} width={120} type={"silos"} changeWidth={() => {return}}/>;
  const sidebar = mount(sidebarComponent);
  expect(sidebar.find('div.sidebar-collapse-btn')).toHaveLength(1);
  const closeBtn = sidebar.find('div.sidebar-collapse-btn').at(0);
  closeBtn.simulate('click');
  sidebar.update();
  expect(sidebar.find('div.sidebar-open-btn')).toHaveLength(1);
});
