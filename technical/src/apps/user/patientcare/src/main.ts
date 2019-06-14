// this import should be first in order to load some required settings (like globals and reflect-metadata)
import { platformNativeScriptDynamic } from "nativescript-angular/platform";
import { registerElement } from "nativescript-angular/element-registry";
import { AppModule } from "./app/app.module";
import { CardView } from 'nativescript-cardview';
import { isDevMode } from '@angular/core';
import * as trace from 'tns-core-modules/trace';
// A traditional NativeScript application starts by initializing global objects, setting up global CSS rules, creating, and navigating to the main page.
// Angular applications need to take care of their own initialization: modules, components, directives, routes, DI providers.
// A NativeScript Angular app needs to make both paradigms work together, so we provide a wrapper platform object, platformNativeScriptDynamic,
// that sets up a NativeScript application and can bootstrap the Angular framework.
platformNativeScriptDynamic().bootstrapModule(AppModule);
registerElement("Ripple", () => require("nativescript-ripple").Ripple);
registerElement('CardView', () => CardView);
import * as app from 'tns-core-modules/application';
import { TraceHelper } from "./app/helpers/trace-helper";

registerElement("DrawingPad", () => require("nativescript-drawingpad").DrawingPad);

app.on(app.launchEvent, (args: app.ApplicationEventData) => {
    console.log('application launch executed with ');
    if (args.android) {
        console.log("Launched Android application with the following intent1: " + args.android + ".");
    } else if (args.ios !== undefined) {
        console.log("Launched iOS application with options: " + args.ios);
    } 
    
    TraceHelper.configure();
});


app.on(app.uncaughtErrorEvent, (args) => {
    console.log('exception occured in applications');
    if (isDevMode()) {
        //trace.write("exception occured in applications", TraceCustomCategory.APP_EXCEPTION, trace.messageType.error);
        //throw args.error;
        console.log(" *** NativeScriptError *** : " + args.android);
        console.log(" *** NativeScriptError stackTrace *** : " + args.android.stackTrace);
        console.log('nativeException',args.android.nativeException);
      //  ErrorHandlerService.getInstance().handleError(args.android.stackTrace); // remove this in production mode.
    } else {
    //    ErrorHandlerService.getInstance().handleError(args.error);
    }

});

app.on(app.lowMemoryEvent, (args: app.ApplicationEventData) => {
    console.log('lowMemoryEvent occured');
});
