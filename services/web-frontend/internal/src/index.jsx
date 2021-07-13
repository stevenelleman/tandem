import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createLogger } from 'redux-logger';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';

import reducer from './redux/reducers/index';
import App from './App';

const middleware = [thunk];
if (process.env.NODE_ENV !== 'production') {
  middleware.push(createLogger());
}

const initialState = {};

const store = createStore(
  reducer,
  initialState,
  applyMiddleware(...middleware),
);

export const StoreDispatch = store.dispatch;

// TODO: Convert jsx -> tsx
ReactDOM.render(
  <Provider store={store}>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </Provider>,
  document.getElementById('root'),
);

/*

Store Shape:
{
  facets: [
    {id: string, banner: string, name: string},...
  ], // Order does not matter
  silos: [
    {id: string, banner: string, name: string}...
  ], // Order does not matter
  forums: [
    {id: string, banner: string, name: string}...
  ], // Order does not matter
  facet_groups: [
    {name: string, facets: [{id: string},...], subgroups:[...]},...
  ], // Order _does_ matter
  silo_groups: [
    {name: string, silos: [{id: string},...], subgroups:[...]},...
  ], // Order _does_ matter
  forum_groups: [
    {name: string, silos: [{id: string},...], subgroups:[...]},...
  ], // Order _does_ matter
  cached_documents: {
    [id: string] : {
      active: boolean,
      // TODO: finish
    },...
  },
}
 */
