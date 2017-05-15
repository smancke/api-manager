
import {
  LOAD_USERDATA,
  LOAD_USERDATA_SUCCESS,
  LOAD_USERDATA_ERROR,
} from './constants';

export function loadUserdata() {
  return {
      type: LOAD_USERDATA,
  };
}

export function userdataLoaded(user) {
  return {
      type: LOAD_USERDATA_SUCCESS,
      user: user,
  };
}

export function userdataLoadingError(error) {
  return {
      type: LOAD_USERDATA_ERROR,
      error: error,
  };
}
