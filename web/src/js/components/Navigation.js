import React, { Component } from 'react'
import PropTypes from 'prop-types'

class Navigation extends Component {

  render() {
    return (
      <div className="navigation">
        <h1>mola mola</h1>
        <div className="navigation-buttons">
          <button className="navigation-button">library</button>
          <button className="navigation-button">hello person</button>
        </div>
      </div>
    );
  }
}

export default Navigation
