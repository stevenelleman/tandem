const envSettings = window as any;
export class EnvVars {
  static protocol = envSettings.REACT_APP_PROTOCOL;
}
