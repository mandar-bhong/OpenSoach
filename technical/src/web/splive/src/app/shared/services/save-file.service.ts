import { Injectable } from '@angular/core';

@Injectable()
export class SaveFileService {
    constructor() { }

    saveFile(data: Blob, filename: string) {
        const blob = new Blob([data]);
        const url = window.URL.createObjectURL(blob);
        const anchor = document.createElement('a');
        anchor.download = filename;
        anchor.href = url;
        document.body.appendChild(anchor);   // added into dom
        anchor.click();
        anchor.remove();   // removed from dom
    }
}
