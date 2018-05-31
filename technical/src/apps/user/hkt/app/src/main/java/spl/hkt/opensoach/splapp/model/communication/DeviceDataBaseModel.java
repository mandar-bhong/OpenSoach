package spl.hkt.opensoach.splapp.model.communication;

import spl.hkt.opensoach.splapp.helper.AppAction;

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
