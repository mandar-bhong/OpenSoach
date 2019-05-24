import { Injectable } from "@angular/core";
import { DialogOptions, CFAlertStyle, CFAlertActionStyle, CFAlertActionAlignment, CFAlertDialog, CFAlertGravity } from "nativescript-cfalert-dialog";

@Injectable()
export class AppNotificationService extends CFAlertDialog {


    private isConfirmed: boolean;
    constructor() {
        super();
    }
    notify(text: string): void {
        let cfalertDialog = new CFAlertDialog();
        let onSelection = response => {
        };

        const options: DialogOptions = {
            dialogStyle: CFAlertStyle.NOTIFICATION,
            textAlignment: CFAlertGravity.CENTER_HORIZONTAL,
            title: text,
            message: null,
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
        cfalertDialog.show(options);
    }
    async confirm(text: string): Promise<boolean> {
        let cfalertDialogConfirm = new AppNotificationService();
        let onSelection = response => {
            cfalertDialogConfirm.isConfirmed = true;
        };

        let onRejection = response => {
            cfalertDialogConfirm.isConfirmed = false;
        }
        const options: DialogOptions = {
            dialogStyle: CFAlertStyle.ALERT,
            textAlignment: CFAlertGravity.CENTER_HORIZONTAL,
            title: null,
            message: text,
            backgroundBlur: false,
            buttons: [
                {
                    text: "Yes",
                    buttonStyle: CFAlertActionStyle.POSITIVE,
                    buttonAlignment: CFAlertActionAlignment.JUSTIFIED,
                    onClick: onSelection
                },
                {
                    text: "No",
                    buttonStyle: CFAlertActionStyle.NEGATIVE,
                    buttonAlignment: CFAlertActionAlignment.JUSTIFIED,
                    onClick: onRejection
                }]
        };
        const userResponse = <boolean>await super.show(options).then(success => {
            return cfalertDialogConfirm.isConfirmed;
        }, error => {
            console.log('error occured in confirm function', cfalertDialogConfirm.isConfirmed);
        });
        return userResponse;
    }


}