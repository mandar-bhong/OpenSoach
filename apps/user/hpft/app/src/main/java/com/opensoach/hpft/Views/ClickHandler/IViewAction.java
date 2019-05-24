package com.opensoach.hpft.Views.ClickHandler;

import com.opensoach.hpft.Model.View.UIViewActionRequestModel;

/**
 * Created by Mandar on 31-10-2017.
 */

public interface IViewAction {

    void ExectuteViewTask(int taskID, UIViewActionRequestModel requestData);
}
