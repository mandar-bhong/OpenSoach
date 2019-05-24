package com.opensoach.vst.Views.ClickHandler;

import com.opensoach.vst.Model.View.UIViewActionRequestModel;

/**
 * Created by Mandar on 31-10-2017.
 */

public interface IViewAction {

    void ExectuteViewTask(int taskID, UIViewActionRequestModel requestData);
}
