import {
  VERB_GET,
  VERB_POST,
  VERB_PUT,
  VERB_DELETE,
} from './constants';

import { isValidUrl } from '../utils';

/*
   TODO:
   - Authz context. Must include cookie(s) in request.
   - Test fetch logic
   - Improved error handling
   - Get vs List state handler
   - Arg/opts handler
   - Make a ts Client interface

   Notes:
   - static method calls class itself, rather than instance
*/

class Client {
  // 0. Initialize Client

  // Props + Types
  origin: string;
  constructor(origin: string) {
    this.origin = origin;
    // TODO: include cookies when instantiated.
  }

  // Static Utility Methods: static methods are class-level methods
  // emitError(msg) null: error handling
  // defaultStateHandler(state) obj: default state handler for fetcher
  // defaultErrorHandler(verb) func: default error handler for fetcher
  private static emitError(msg: string): void {
    console.log(`ERROR: ${msg}`);
  }

  private static defaultStateHandler(state: object): object {
    return state;
  }

  private static defaultErrorHandler(verb: string): (err: string) => void {
    return (err: string) => {
      this.emitError(`${verb} error: ${err}`);
    }
  }

  // Instance Methods: methods called on the instance itself
  // TODO: When active FIs are updated the client cookies need to be updated as well
  updateAuthz() {
    return;
  }

  // 1. Request API: by design this will be the _one_ method used 90% of the time by the reducers.
  // The other 10% of the time will be methods to update the state of the Client (such as FI add/remove).
  // TODO: args/opts validation?
  public request(verb: string, route: string, args={}, opts={}) {
    const url = `${this.origin}/${route}`;
    if (!isValidUrl(url)) {
      return Client.emitError(`invalid url: ${url}`);
    }

    let body = {
      method: verb,
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(args),
    };

    switch (verb) {
      case VERB_GET:
        return Client[VERB_GET](url);
      case VERB_POST:
        return Client[VERB_POST](url, body);
      case VERB_PUT:
        return Client[VERB_PUT](url, body);
      case VERB_DELETE:
        return Client[VERB_DELETE](url);
      default:
        return Client.emitError(`verb ${verb} is not recognized`);
    }
  }

  // 2. Fetch Methods:
  private static [VERB_GET](url: string, body= {method: VERB_GET}, opts={}) {
    const stateHandler = (state: object): object => {
      // TODO: GET vs LIST?
      return state;
    };

    const errorHandler = (err: string): void => {
      Client.emitError(`${VERB_GET} error: ${err}`);
      // Temporary solution for spacing our GET request following the PULL request - continue making request until POST unblocks
      // Add some kind of exponential dropoff
      setTimeout(() => {
        Client[VERB_GET](url, body, opts)
      }, 200)
    };

    return Client.fetcher(url, body, stateHandler, errorHandler);
  }

  private static [VERB_POST](url: string, body= {method: VERB_POST}, opts={}) {
    return Client.fetcher(url, body, Client.defaultStateHandler, Client.defaultErrorHandler(VERB_POST));
  }

  private static [VERB_PUT](url: string, body= {method: VERB_PUT}, opts={}) {
    return Client.fetcher(url, body, Client.defaultStateHandler, Client.defaultErrorHandler(VERB_PUT));
  }

  private static [VERB_DELETE](url: string, body = {method: VERB_DELETE}, opts={}) {
    return Client.fetcher(url, body, Client.defaultStateHandler, Client.defaultErrorHandler(VERB_DELETE));
  }

  // 3. Fetcher: by design this is the only fetch call. Let's see if it can be generalized for all the requests.
  private static fetcher(url: string, body: object, stateHandler: (state: object) => object, errorHandler: (err: string) => void) {
    return fetch(url, body).then(resp => resp.json()).then(stateHandler).catch(errorHandler);
  }
}
