
import { StoreDispatch } from '../index';

// Store Types
export type AppDispatch = typeof StoreDispatch;

// Interface Definitions
interface EmptyAction {
  type: string
}

interface ErrorAction {
  type: string,
  error: string,
}

// Type Definitions
export type RequestSilosActionType = () => EmptyAction;
export type RequestSilosErrorType = (error: string) => ErrorAction;
// TODO: Update "any" to Silo interface when it's better defined
export type ReceiveSiloActionType = (
  resp: {list: Array<any>}
) => ({type: string, silos: Array<any>});

// Dispatch Method Example:
// (client: Client) => Dispatch<Action<receiveSiloActionType>>
