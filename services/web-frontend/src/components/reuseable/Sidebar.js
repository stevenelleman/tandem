import React from 'react';

/*
TODO:
- Enable tabs
- group silos, group forums, group tabs - works out quite nicely actually

Arguments:
- type title
- list of objects (silo, forum, tabs)
- left/right/bottom (just right and left initially, expand to bottom later but prepare it for that
- switch dropdown -> function
 */

export class Sidebar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      active: true,
    };
  }

  // Use for collapsing and opening sidebar
  toggleBar() {
    const { left, changeWidth } = this.props;
    const { active } = this.state;
    const orientation = left ? 'left' : 'right';
    if (active) {
      changeWidth(orientation, 20);
    } else {
      changeWidth(orientation, 120);
    }

    this.setState({
      active: !active,
    });
  }

  openSymbol() {
    const { left } = this.props;
    const { active } = this.state;
    // Left may be undefined instead of false. !! converts it to false.
    if (!!left !== active) {
      return '->';
    }
    return '<-';
  }

  // Populates side buttons: Collapse/Open, Drag, and Options
  // Collapse/Open Buttons: Close and open sidebars
  // Drag Button: Change the width of the sidebars
  // Options Button: Dropdown to switch one sidebar with another
  sidebarButtons() {
    // Constants
    const { width, left, type } = this.props;
    const { active } = this.state;

    // Title width
    const titleWidth = width - 80;

    // Generate classes
    const floatClass = left ? 'float-right' : 'float-left';
    const titleClass = !active ? 'sidebar-title-collapsed' : 'sidebar-title';
    const sidebarClass = `sidebar-buttons ${(left ? 'left' : 'right')}`;

    // Generate divs
    const openSymbol = this.openSymbol();
    const openBtn = <div key="open" onClick={() => this.toggleBar()} className={`sidebar-open-btn ${floatClass}`}>{openSymbol}</div>;
    const collapseBtn = <div key="collapse" onClick={() => this.toggleBar()} className={`sidebar-collapse-btn ${floatClass}`}>{openSymbol}</div>;
    const createBtn = <div key="create" className={`sidebar-create-btn ${floatClass}`}>+</div>;
    const dragBtn = <div key="drag" className={`sidebar-drag-btn ${floatClass}`}/>;
    const optionsBtn = <div key="options" className={`sidebar-options-btn ${floatClass}`}/>;
    const title = <div key="title" className={`${titleClass} ${floatClass}`} style={{ titleWidth }}>{type}</div>;

    // Push divs into array
    const buttons = [];
    if (!active) {
      buttons.push(openBtn);
    } else {
      buttons.push(dragBtn, collapseBtn, createBtn, optionsBtn, title);
    }
    return (
      <div className={sidebarClass}>
        {buttons}
      </div>
    );
  }


  render() {
    const { width, left } = this.props;
    const { active } = this.state;
    const orientation = left ? 'left' : 'right';
    const sidebarClass = `${orientation} ${(!active ? 'sidebar-collapsed' : 'sidebar')}`;
    const contentsClass = `${orientation} ${(!active ? 'sidebar-contents-collapsed' : 'sidebar-contents')}`;
    return (
      <div className={sidebarClass} style={{ width }}>
        {this.sidebarButtons()}
        <div className={contentsClass} />
      </div>
    );
  }
}
