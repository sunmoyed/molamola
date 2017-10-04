import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Route, Switch } from 'react-router'

import Navigation from './Navigation'
import Library from './Library'
import Login from './Login'


const ConnectedSwitch = connect(state => ({
  location: {...state.routerReducer.location}
}))(Switch);

const App = ({ location, user }) => {
  const loggedIn = !!user.username

  return (
    <div className="root">
      <Navigation />
      <div className="page">
        <ConnectedSwitch>
          <Route exact path="/" component={loggedIn ? Library : Login}/>
          {loggedIn &&
            <Route path="/library" component={Library}/>}
          {!loggedIn &&
            <Route path="/login" component={Login}/>}
        </ConnectedSwitch>
      </div>
    </div>
  )
}

export default connect(state => ({
  user: {...state.rootReducer.user},
  location: {...state.routerReducer.location}
}))(App)
