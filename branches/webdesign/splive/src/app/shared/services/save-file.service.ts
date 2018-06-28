import { Injectable } from '@angular/core';

@Injectable()
export class SaveFileService {
    constructor() { }

    saveFile(data: Blob, filename: string) {
        const blob = new Blob([data]);
        const url = window.URL.createObjectURL(blob);
        console.log('url dowload');
        console.log(url);
        const anchor = document.createElement('a');
        anchor.download = filename;
        anchor.href = url;
        anchor.click();
    }
}
