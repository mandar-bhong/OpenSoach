import { Component, OnInit } from "@angular/core";
import { ModalDialogParams } from "nativescript-angular/modal-dialog";
import { registerElement } from 'nativescript-angular/element-registry';
registerElement('ImageZoom', () => require('nativescript-image-zoom').ImageZoom);

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

