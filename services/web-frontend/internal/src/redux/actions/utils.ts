// TODO: keep in this file, mov
// After Client.fetcher has finished dispatcher is used to dispatch ERROR/RECEIVE action.
export function dispatcher(dispatch: any, action: any) {
  return (state: any) => dispatch(action(state));
}
