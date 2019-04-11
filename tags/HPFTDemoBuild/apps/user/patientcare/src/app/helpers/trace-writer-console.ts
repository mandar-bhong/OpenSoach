import * as trace from 'trace';
export class TraceConsoleWriter {
    constructor() { }

    public write(message, category, type) {
        if (!console) {
            return;
        }

        const traceMessage = new Date().toISOString() + " " + category + ": " + message;

        switch (type) {
            case trace.messageType.info:
                console.info(traceMessage);
                break;
            case trace.messageType.info:
                console.warn(traceMessage);
                break;
            case trace.messageType.error:
                console.error(traceMessage);
            default:
                console.log(traceMessage);
                break;
        }
    }
}