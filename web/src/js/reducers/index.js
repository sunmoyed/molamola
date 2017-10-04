import * as types from '../actions/actions';
import {browserHistory} from 'react-router';

const initialState = {
  user: {
    username: "olas",
    avatar: "photoloas.jpg"
  },
  library: [["cowboy bebop", "fav"],
            ["penguindrum", "??? what even is"],
            ["psycho-pass", "I still need to finish this xD"] ]}

export default (state = initialState, action) => {
  switch (action.type) {
    case types.LOG_IN_SUCCESS:
      return Object.assign({}, state, {user: action.user})
    case types.LOG_OUT:
      // browserHistory.push('/')  // IMPURE! // TODO where is browserHistory?
      !!sessionStorage.jwt
      return Object.assign({}, state, {user: null})
    default:
      return {...state}
  }
}

// const new_obj = {...old_obj, new_val: 2};
