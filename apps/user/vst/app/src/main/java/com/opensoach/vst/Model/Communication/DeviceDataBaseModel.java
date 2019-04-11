package com.opensoach.vst.Model.Communication;

import com.opensoach.vst.Helper.AppAction;

/**
 * Created by Mandar on 4/13/2017.
 */

public class DeviceDataBaseModel {

    private AppAction commandType;

    public AppAction getUserActionType() {
        return commandType;
    }

    public void setUserActionType(AppAction commandType) {
        this.commandType = commandType;
    }
}
