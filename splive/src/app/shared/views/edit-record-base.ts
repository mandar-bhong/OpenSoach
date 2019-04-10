import { FormGroup } from '@angular/forms';

export enum FORM_MODE {
    VIEW = 0,
    EDITABLE = 1,
}

export enum EDITABLE_RECORD_STATE {
    ADD = 0,
    UPDATE = 1,
}

export abstract class EditRecordBase {
    editableForm: FormGroup;
    formMode = FORM_MODE.VIEW;
    recordState = EDITABLE_RECORD_STATE.ADD;
    callbackUrl: string;
    formModeEnum = FORM_MODE;
    recordStateEnum = EDITABLE_RECORD_STATE;
    isEditable = true;
    showBackButton = true;
    pageTitle = '';
    subTitle = '';
    iconCss: string;
    iconName: string;
    inProgress = false;

    editForm() {
        this.setFormMode(FORM_MODE.EDITABLE);
    }

    setFormMode(mode: FORM_MODE) {
        this.formMode = mode;
        switch (this.formMode) {
            case FORM_MODE.VIEW:
                this.editableForm.disable();
                break;
            case FORM_MODE.EDITABLE:
                this.editableForm.enable();
                break;
        }
    }

    abstract closeForm();

    onCancelHandler(){}

}
