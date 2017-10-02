import React, { Component } from 'react'
import { connect } from 'react-redux'
import Navigation from './Navigation'
import Library from './Library'

class App extends Component {

  render() {
    return (
      <div className="root">
        <Navigation />
        <div className="page">
          <LibraryContainer />
        </div>
      </div>
    )
  }
}

const mapStateToProps = (state) => ({
  library: state.library
})

const mapDispatchToProps = {}

const LibraryContainer = connect(mapStateToProps, mapDispatchToProps)(Library)


export default App
