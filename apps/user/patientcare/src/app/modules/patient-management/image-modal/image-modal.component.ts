import { Component, OnInit } from "@angular/core";
import { ModalDialogParams } from "nativescript-angular/modal-dialog";
import { registerElement } from 'nativescript-angular/element-registry';
import { PDFView } from 'nativescript-pdf-view';
registerElement('ImageZoom', () => require('nativescript-image-zoom').ImageZoom);
registerElement('PDFView', () => PDFView);

@Component({
    moduleId: module.id,
    selector: 'image-modal',
    templateUrl: './image-modal.component.html',
    styleUrls: ['./image-modal.component.css']
})
export class ImageModalComponent implements OnInit {

    ctx: any;

    constructor(private params: ModalDialogParams) { }

    ngOnInit() {
        this.ctx = this.params.context;
    }

    close() {
        this.params.closeCallback();
    }
}

