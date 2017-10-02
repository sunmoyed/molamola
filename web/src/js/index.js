import React from 'react'
import { render } from 'react-dom'
import { createStore } from 'redux'
import { Provider } from 'react-redux'

import App from './components/App'
import rootReducer from './reducers'

const initialState = {
  library: [["cowboy bebop", "fav"],
            ["penguindrum", "??? what even is"],
            ["psycho-pass", "I still need to finish this xD"] ]}

const store = createStore(rootReducer, initialState)

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
)
