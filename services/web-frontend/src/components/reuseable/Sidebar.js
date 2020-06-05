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
    }
  }

  // Use for collapsing and opening sidebar
  toggleBar() {
    const orientation = this.props.left ? "left" : "right";
    if (this.state.active) {
      this.props.changeWidth(orientation, 20);
    } else {
      this.props.changeWidth(orientation, 120);
    }

    this.setState({
      active: !this.state.active,
    })
  }

  openSymbol() {
    var left = this.props.left ? this.props.left : false;
    if (left !== this.state.active) {
      return "->";
    }
    return "<-";
  }

  // Populates side buttons: Collapse/Open, Drag, and Options
  // Collapse/Open Buttons: Close and open sidebars
  // Drag Button: Change the width of the sidebars
  // Options Button: Dropdown to switch one sidebar with another
  sidebarButtons() {
    // Title width
    const width = this.props.width - 80;

    // Generate classes
    const floatClass = this.props.left ? "float-right" : "float-left";
    const titleClass = !this.state.active ? "sidebar-title-collapsed" : "sidebar-title";
    const sidebarClass = "sidebar-buttons " + (this.props.left ? "left" : "right");

    // Generate divs
    var openSymbol = this.openSymbol();
    var openBtn = <div key="open" onClick={() => this.toggleBar()} className={`sidebar-open-btn ${floatClass}`}>{openSymbol}</div>
    var collapseBtn = <div key="collapse" onClick={() => this.toggleBar()} className={`sidebar-collapse-btn ${floatClass}`}>{openSymbol}</div>;
    var createBtn = <div key="create" className={`sidebar-create-btn ${floatClass}`}>+</div>;
    var dragBtn = <div key="drag" className={`sidebar-drag-btn ${floatClass}`}/>;
    var optionsBtn = <div key="options" className={`sidebar-options-btn ${floatClass}`}/>;
    var title =  <div key="title" className={`${titleClass} ${floatClass}`} style={{width}}>{this.props.type}</div>

    // Push divs into array
    var buttons = [];
    if (!this.state.active) {
      buttons.push(openBtn)
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
    const width = this.props.width;
    const orientation = this.props.left ? "left" : "right";
    const sidebarClass = orientation + (!this.state.active ? " sidebar-collapsed" : " sidebar");
    const contentsClass = orientation + (!this.state.active ? " sidebar-contents-collapsed" : " sidebar-contents");
    return (
      <div className={sidebarClass} style={{width}}>
        {this.sidebarButtons()}
        <div className={contentsClass}>

        </div>
      </div>
    );
  }
}
