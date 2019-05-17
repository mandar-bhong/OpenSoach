package com.opensoach.vst.Model.DB;

import com.opensoach.vst.DAL.DBConstants;
import com.opensoach.vst.DAL.DBTableSchema;

/**
 * Created by samir.s.bukkawar on 2/26/2017.
 */
@DBTableSchema(TableName = DBConstants.TABLE_TASKS)
public class DBTaskTableRowModel {

    private int taskId;
    private String taskName;
    private int taskOrder;

    public int getTaskId() {
        return taskId;
    }

    public void setTaskId(int taskId) {
        this.taskId = taskId;
    }

    public String getTaskName() {
        return taskName;
    }

    public void setTaskName(String taskName) {
        this.taskName = taskName;
    }

    public int getTaskOrder() {
        return taskOrder;
    }

    public void setTaskOrder(int taskOrder) {
        this.taskOrder = taskOrder;
    }
}
