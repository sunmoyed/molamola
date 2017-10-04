import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Route, Switch } from 'react-router'

import Navigation from './Navigation'
import Library from './Library'
import Login from './Login'


const ConnectedSwitch = connect(state => ({
  location: {...state.routerReducer.location}
}))(Switch);

const App = ({ location }) => (
   <div className="root">
     <Navigation />
     <div className="page">
       <ConnectedSwitch>
         <Route exact path="/" component={Library}/>
         <Route path="/library" component={Library}/>
         <Route path="/login" component={Login}/>
      </ConnectedSwitch>
     </div>
   </div>
);

export default connect(state => ({location: state.location}))(App)
