import React, { Component } from 'react'
import Navigation from './Navigation'
import Library from './Library'

class App extends Component {

  render() {
    // const { children, inputValue } = this.props
    return (
      <div className="root">
        <Navigation />
        <div className="page">
          <Library />
        </div>
      </div>
    )
  }
}

export default App
