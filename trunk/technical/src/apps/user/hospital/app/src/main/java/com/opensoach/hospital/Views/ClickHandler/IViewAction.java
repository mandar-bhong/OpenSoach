package com.opensoach.hospital.Views.ClickHandler;

import com.opensoach.hospital.Model.View.UIViewActionRequestModel;

/**
 * Created by Mandar on 31-10-2017.
 */

public interface IViewAction {

    void ExectuteViewTask(int taskID, UIViewActionRequestModel requestData);
}
