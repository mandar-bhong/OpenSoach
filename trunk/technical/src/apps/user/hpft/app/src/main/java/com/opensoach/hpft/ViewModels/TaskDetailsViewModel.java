package com.opensoach.hpft.ViewModels;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Mandar on 01-08-2018.
 */

public class TaskDetailsViewModel extends BaseViewModel {
    ArrayList<String> tasks;

    public ArrayList<String> getTasks() {
        return tasks;
    }

    public void setTasks(ArrayList<String> tasks) {
        this.tasks = tasks;
    }
}
