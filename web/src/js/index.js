import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'

import { applyMiddleware, combineReducers, compose, createStore } from 'redux'
import thunk from 'redux-thunk';
import { ConnectedRouter, routerReducer, routerMiddleware } from 'react-router-redux'

import createHistory from 'history/createBrowserHistory'

import App from './components/App'
import rootReducer from './reducers'


const history = createHistory()

// TODO env
// this is for redux-devtools browser extension
// it's useful, you might want to install it \o/
let composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

let store = createStore(
  combineReducers({ routerReducer, rootReducer }),
  composeEnhancers(
    applyMiddleware(thunk),
    applyMiddleware(routerMiddleware(history))
  ))

// dispatch navigation actions from anywhere: `store.dispatch(push('/thing'))`
// also: import {push} from 'react-router-redux'

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <App />
    </ConnectedRouter>
  </Provider>,
  document.getElementById('app'),
);
