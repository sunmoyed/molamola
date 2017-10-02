import React, { Component } from 'react'
import PropTypes from 'prop-types'

class Library extends Component {

  render() {
    return (
      <div className="library">
        <div className="list">
          <div className="row">
            <div className="column">
              title
            </div>
            <div className="column">
              information
            </div>
          </div>

          <div className="row">
            <div className="column">
              I am an anime
            </div>
            <div className="column">
              an image
            </div>
          </div>

          <div className="row">
            <div className="column">
              hello
            </div>
            <div className="column">
              potato!
            </div>
          </div>

          <div className="row">
            <div className="column">
              I am an anime
            </div>
            <div className="column">
              an image
            </div>
          </div>

        </div>
      </div>
    );
  }
}

export default Library
