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
    - Arg/opts handler
    - Make a ts Client interface
    - Make object types
*/

export class Client {
  // Props Types + Class Constructor
  origin: string;
  constructor(origin: string) {
    this.origin = origin;
    // TODO: include cookies when instantiated.
  }

  // Static Utility Methods: static methods are class-level methods
  // emitError(msg) null: error handling
  // defaultStateHandler(state) obj: default state handler for fetcher
  // defaultErrorHandler(verb) func: default error handler for fetcher
  // TODO: why can I not make this method static?
  private emitError(msg: string): void {
    console.log(`ERROR: ${msg}`);
  }

  // Instance Methods: methods called on the instance itself
  // TODO: When active FIs are updated the client cookies need to be updated as well
  updateAuthz(): void {
    return;
  }

  // 1. Request API: by design this will be the _one_ method used 90% of the time by the reducers.
  // The other 10% of the time will be methods to update the state of the Client (such as FI add/remove).
  // TODO: args/opts validation?
  //  - Better stateHandler type
  //  - Pass in a generic error action in App
  public request(verb: string, route: string, args={}, opts={}, stateHandler: any, errorHandler: any) {
    const url = `http://${this.origin}/${route}`;
    if (!isValidUrl(url)) {
      return this.emitError(`invalid url: ${url}`);
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
        return this[VERB_GET](url, {method: VERB_GET}, {}, stateHandler, errorHandler);
      case VERB_POST:
        return this[VERB_POST](url, body, {}, stateHandler, errorHandler);
      case VERB_PUT:
        return this[VERB_PUT](url, body, {}, stateHandler, errorHandler);
      case VERB_DELETE:
        return this[VERB_DELETE](url, {method: VERB_DELETE}, {}, stateHandler, errorHandler);
      default:
        return this.emitError(`verb ${verb} is not recognized`);
    }
  }

  // TODO: Why can I not makes these methods static?
  // 2. Fetch Methods:
  private [VERB_GET](url: string, body: any, opts={}, stateHandler: any, errorHandler: any) {
    /*const errorHandler = (err: string): void => {
      this.emitError(`${VERB_GET} error: ${err}`);
      // Temporary solution for spacing our GET request following the PULL request - continue making request until POST unblocks
      // Add some kind of exponential dropoff
      setTimeout(() => {
        this[VERB_GET](url, body, opts, stateHandler)
      }, 200)
    };*/
    return this.fetcher(url, body, stateHandler, errorHandler);
  }

  private [VERB_POST](url: string, body: any, opts={}, stateHandler: any, errorHandler: any) {
    return this.fetcher(url, body, stateHandler, errorHandler);
  }

  private [VERB_PUT](url: string, body: any, opts={}, stateHandler: any, errorHandler: any) {
    return this.fetcher(url, body, stateHandler, errorHandler);
  }

  private [VERB_DELETE](url: string, body: any, opts={}, stateHandler: any, errorHandler: any) {
    return this.fetcher(url, body, stateHandler, errorHandler);
  }

  // 3. Fetcher: by design this is the only fetch call. Let's see if it can be generalized for all the requests.
  private fetcher(url: string, body: any, stateHandler: (state: any) => any, errorHandler: (err: string) => void) {
    // TODO: check out cross-fetch?
    return fetch(url, body).then(resp => resp.json()).then(stateHandler).catch(errorHandler);
  }
}
