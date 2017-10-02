import React from 'react'
import ReactDOM from 'react-dom'
import { createStore } from 'redux'
import App from './components/App'
import rootReducer from './reducers'

const initialState = {}

const store = createStore(rootReducer, initialState)
const rootElement = document.getElementById('app')

const render = () => ReactDOM.render(
  <App />,
  rootElement
)

render()
store.subscribe(render)
