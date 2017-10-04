import React, { Component } from 'react'
import { connect } from 'react-redux'
import {bindActionCreators} from 'redux';
import { Link } from 'react-router-dom'
import PropTypes from 'prop-types'

import { logOutUser } from '../actions/sessionActions'

class Navigation extends Component {

  logout(e) {
    e.preventDefault()
    this.props.onLogoutClick()
  }

  render() {
    const {username} = this.props.user

    return (
      <div className="navigation">
        <h1><Link to="/">mola mola</Link></h1>
        <div className="navigation-buttons">
          {username &&
            <Link to="/library" className="navigation-button">{`hello ${username}`}</Link>}
          {username?
            <button
              className="navigation-button"
              onClick={this.props.onLogoutClick}>logout</button> :
            <Link to="/login" className="navigation-button">login</Link>}
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state) => ({user: {...state.rootReducer.user}})

const mapDispatchToProps = (dispatch) => ({onLogoutClick: bindActionCreators(logOutUser, dispatch)})

export default connect(mapStateToProps, mapDispatchToProps)(Navigation)
