import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'

class Navigation extends Component {

  render() {
    const {username} = this.props.user

    return (
      <div className="navigation">
        <h1><Link to="/">mola mola</Link></h1>
        <div className="navigation-buttons">
          <Link to="/login" className="navigation-button">login</Link>
          {username &&
            <Link to="/library" className="navigation-button">{`hello ${username}`}</Link>}
          {username?
            <button className="navigation-button">log out</button> :
            <Link to="/login" className="navigation-button">login</Link>}
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {user: {...state.rootReducer.user}}}

const mapDispatchToProps = {}

export default connect(mapStateToProps, mapDispatchToProps)(Navigation)
