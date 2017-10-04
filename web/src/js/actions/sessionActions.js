import * as types from './actions';

export function loginSuccess() {
  return {type: types.LOG_IN_SUCCESS}
}

export function loginUser(credentials) {
  return function(dispatch) {
    return login(credentials).then(response => {
      console.log("we got a token!", response.jwt);
      sessionStorage.setItem('molamola-jwt', response.jwt);
      dispatch(loginSuccess());
    }).catch(error => {
      debugger;
      throw(error);
    });
  };
}

// IMPURE!
export function logOutUser() {
  sessionStorage.removeItem('jwt');

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
