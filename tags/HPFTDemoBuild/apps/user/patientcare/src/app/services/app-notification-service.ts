import { Injectable } from "@angular/core";
import { DialogOptions, CFAlertStyle, CFAlertActionStyle, CFAlertActionAlignment, CFAlertDialog, CFAlertGravity } from "nativescript-cfalert-dialog";

@Injectable()
export class AppNotificationService {
   
    private cfalertDialog: CFAlertDialog;
    constructor(){
        this.cfalertDialog = new CFAlertDialog();
    }
    notify(text:string): void {
        let onSelection = response => {
        };
      
        const options: DialogOptions = {
            dialogStyle: CFAlertStyle.NOTIFICATION,
            textAlignment: CFAlertGravity.CENTER_HORIZONTAL,
            title: null,
            message: text,
            backgroundBlur: false,
            onDismiss: () => console.log("showAlert dismissed"),
            buttons: [
                {
                    text: "Ok",
                    buttonStyle: CFAlertActionStyle.NEGATIVE,
                    buttonAlignment: CFAlertActionAlignment.END,
                    onClick: onSelection
                }]
        };
        this.cfalertDialog.show(options);
    }
}