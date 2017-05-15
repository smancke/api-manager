
import { take, call, put, select, cancel, takeEvery } from 'redux-saga/effects';
import { LOCATION_CHANGE } from 'react-router-redux';
import { LOAD_USERDATA } from 'containers/App/constants';
import { userdataLoaded, userdataLoadingError } from 'containers/App/actions';

function jwt_decode (token) {
  var base64Url = token.split('.')[1];
  var base64 = base64Url.replace('-', '+').replace('_', '/');
  return JSON.parse(window.atob(base64));
};

/**
 * Take the userdata from JWT cookie.
 */
function* getUserdata() {
    var cookieName = "jwt_token"
    var cookieList = document.cookie.match('(^|;)\\s*' + cookieName + '\\s*=\\s*([^;]+)')

    if  (! cookieList) {
        yield put(userdataLoadingError())
    }
    
    var token = cookieList.pop()
    try {
        yield put(userdataLoaded(jwt_decode(token)));
    } catch (exception) {
        console.log(exception )
        yield put(userdataLoadingError("error decoding token"));
    }
}

/**
 * Root saga manages watcher lifecycle
 */
function* userData() {
  yield takeEvery(LOAD_USERDATA, getUserdata);
}

// Bootstrap sagas
export default [
  userData,
];
