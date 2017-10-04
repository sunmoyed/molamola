import React, { Component } from 'react'
import Navigation from './Navigation'
import Library from './Library'
import Login from './Login'

class App extends Component {

  render() {
    return (
      <div className="root">
        <Navigation />
        <div className="page">
          <Login />
          <Library />
        </div>
      </div>
    )
  }
}


export default App
