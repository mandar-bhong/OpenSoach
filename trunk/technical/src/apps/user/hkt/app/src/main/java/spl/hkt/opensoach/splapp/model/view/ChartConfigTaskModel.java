package spl.hkt.opensoach.splapp.model.view;

import android.support.annotation.NonNull;

/**
 * Created by Mandar on 3/27/2017.
 */

public class ChartConfigTaskModel implements Comparable<ChartConfigTaskModel>  {

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

    @Override
    public int compareTo(@NonNull ChartConfigTaskModel chartTaskModel) {

        if (taskOrder > chartTaskModel.getTaskOrder()) {
            return 1;
        }
        else if (taskOrder <  chartTaskModel.getTaskOrder()) {
            return -1;
        }
        else {
            return 0;
        }
    }
}
