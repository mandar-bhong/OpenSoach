import { ErrorHandler, Injectable } from '@angular/core';
import * as trace from 'tns-core-modules/trace';
import { TraceCustomCategory } from './trace-helper';

@Injectable({ providedIn: 'root' })
export class TraceAngularErrorHandler extends ErrorHandler {
    handleError(e: Error) {
        console.log("Error handled at TraceAngularErrorHandler");        
        e.name = TraceCustomCategory.APP_EXCEPTION;
        trace.error(e);
    }
}