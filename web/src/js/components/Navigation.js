import React, { Component } from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'

class Navigation extends Component {

  render() {
    const {username} = this.props.user

    return (
      <div className="navigation">
        <h1>mola mola</h1>
        <div className="navigation-buttons">
          {username && <button className="navigation-button">{`hello ${username} (log out)`}</button>}
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state) => ({user: {...state.user}})

const mapDispatchToProps = {}

export default connect(mapStateToProps, mapDispatchToProps)(Navigation)
