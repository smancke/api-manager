
import { combineReducers } from 'redux-immutable';
import { fromJS } from 'immutable';
import { LOCATION_CHANGE } from 'react-router-redux';
import {
  LOAD_USERDATA,
  LOAD_USERDATA_SUCCESS,
  LOAD_USERDATA_ERROR,
} from './constants';

const userInitialState = fromJS({});

function userReducer(state = userInitialState, action) {
    console.log(action)
    switch (action.type) {
    case LOAD_USERDATA:
        return state
            .set('user', {})
            .set('error', false);
    case LOAD_USERDATA_SUCCESS:        
        return state
            .set('user', action.user)
            .set('error', false);
    case LOAD_USERDATA_ERROR:
        return state
            .set('error', action.error)
    default:
        return state;
    }
}
    
export default userReducer;
