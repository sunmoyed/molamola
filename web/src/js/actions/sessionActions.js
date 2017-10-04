import * as types from './actions';

const SESSION_STORAGE_KEY = 'molamola-jwt'

export function loginSuccess(credentials) {
  return {
    type: types.LOG_IN_SUCCESS,
    user: credentials
  }
}

export function loginUser(credentials) {
  return function(dispatch) {
    return login(credentials).then(response => {
      console.log("we got a token!", response.jwt);
      sessionStorage.setItem(SESSION_STORAGE_KEY, response.jwt);
      // TODO this is fake
      dispatch(loginSuccess({username: credentials.username, avatar: null}));
      // dispatch(loginSuccess(response.user));
    }).catch(error => {
      debugger;
      throw(error);
    });
  };
}

export function logOutUser(e) {
  e.preventDefault()
  sessionStorage.removeItem(SESSION_STORAGE_KEY); // IMPURE!

  return {type: types.LOG_OUT}
}

// class Auth {
//   static loggedIn() {
//     return !!sessionStorage.jwt;
//   }
//
//   static logOut() {
//     sessionStorage.removeItem('jwt');
//   }
// }

// IMPURE!
const login = (credentials) => {
  // const request = new Request(`${process.env.API_HOST}/login`, {
  //   method: 'POST',
  //   headers: new Headers({
  //     'Content-Type': 'application/json'
  //   }),
  //   body: JSON.stringify({auth: credentials})
  // });
  //
  // return fetch(request).then(response => {
  //   return response.json();
  // }).catch(error => {
  //   return error;
  // });

  // TODO replace fake network calls \o/
  const request = new Promise((resolve, reject) => {
    setTimeout(function(){
      resolve({json: () => {return {jwt: 'ayyy'}}});
    }, 250);
  });

  return request.then(response => {
    return response.json();
  }).catch(error => {
    return error;
  });
}
