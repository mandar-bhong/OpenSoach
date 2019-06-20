import { environment as devEnvironment } from './environment.dev';
import { environment as prodEnvironment } from './environment.prod';


declare var process:any;

export const environment = (() => {
  let envVars;

  if (
    typeof process !== 'undefined' && process &&
    Object.prototype.hasOwnProperty.call(process, 'env') && process.env &&
    Object.prototype.hasOwnProperty.call(process.env, 'buildmode') && process.env.buildmode
  ) {
    switch (process.env.buildmode) {
      case 'prod':
        console.log("Build mode is prod");
        envVars = prodEnvironment;
        break;
      // TODO: Add additional environment (e.g. uat) if required. 
      default:
          console.log("Build mode is dev");
        envVars = devEnvironment;
        break;
    }
  } else {
    envVars = devEnvironment;
  }

  return envVars;
})();