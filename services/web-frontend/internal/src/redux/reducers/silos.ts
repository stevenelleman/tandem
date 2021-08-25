import { REQUEST_SILOS, REQUEST_SILOS_ERROR, RECEIVE_SILOS } from '../constants';

// TODO: type checking for action? Or too generic for type-checking?
export const initialState = {};
export const silos = (state = initialState, action: any) => {
  switch (action.type) {
    case REQUEST_SILOS:
      return {
        ...state,
        ...{ silos: { isFetching: true } },
      };
    case REQUEST_SILOS_ERROR:
      return {
        ...state,
        ...{ silos: { isFetching: false } },
      };
    case RECEIVE_SILOS:
      return {
        ...state,
        ...{ silos: { isFetching: false, list: action.silos } },
      };
    default:
      return state;
  }
};
