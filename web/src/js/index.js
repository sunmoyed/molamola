import React from 'react'
import { render } from 'react-dom'
import { createStore, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import thunk from 'redux-thunk';

import App from './components/App'
import rootReducer from './reducers'

const initialState = {
  user: {
    username: "olas",
    avatar: "photoloas.jpg"
  },
  library: [["cowboy bebop", "fav"],
            ["penguindrum", "??? what even is"],
            ["psycho-pass", "I still need to finish this xD"] ]}

const store = createStore(rootReducer, initialState, applyMiddleware(thunk))

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
)
