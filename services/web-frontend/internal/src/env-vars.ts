const envSettings = window as any;
export class EnvVars {
  static protocol: string = envSettings.REACT_APP_PROTOCOL;
}
