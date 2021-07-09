import { Dispatch } from 'redux';
import { Client, VERB_GET } from 'client';

import { dispatcher } from './utils';
import { REQUEST_SILOS, REQUEST_SILOS_ERROR, RECEIVE_SILOS } from '../constants';

// Note: If these are ever used again move to `types` directory
interface silos {
  isFetching: boolean;
  didInvalidate: boolean;
  list: string[];
}

interface IState {
  silos: silos;
}

interface IResp {
  type: string;
  list: string[];
}

function shouldFetchSilos(state: IState) {
  const { silos } = state;
  if (!silos || !Object.keys(silos).length) {
    return true;
  }
  if (silos.isFetching) {
    return false;
  }
  return silos.didInvalidate;
}

function requestSilos() {
  return {
    type: REQUEST_SILOS,
  };
}

function requestSilosError(error: string) {
  return {
    type: REQUEST_SILOS_ERROR,
    error,
  };
}

function receiveSilos(resp: IResp) {
  return {
    type: RECEIVE_SILOS,
    silos: resp.list,
  };
}

export function fetchSilosIfNeeded(client: Client) {
  // TODO: The type of getState depends on the shape of the store object
  //  using "any" is a bit lazy
  return (dispatch: Dispatch, getState: any): void => {
    if (shouldFetchSilos(getState())) {
      // Dispatch action for request
      dispatch(requestSilos());
      client.request(
        VERB_GET,
        'v1/silos',
        {},
        {},
        dispatcher(dispatch, receiveSilos),
        dispatcher(dispatch, requestSilosError),
      );
    }
  };
}
