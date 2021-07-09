import { Dispatch } from 'redux';
import { Client, VERB_GET } from 'client';

import { dispatcher } from './utils';
import { REQUEST_SILOS, REQUEST_SILOS_ERROR, RECEIVE_SILOS } from '../constants';

function shouldFetchSilos(state: any) {
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

// TODO: update resp type as API becomes clearer
function receiveSilos(resp: any) {
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
