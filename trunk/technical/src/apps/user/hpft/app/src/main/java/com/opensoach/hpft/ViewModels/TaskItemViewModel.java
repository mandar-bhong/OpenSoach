/*
 * Copyright (c) 2018 Phunware Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:

 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.

 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package com.opensoach.hpft.ViewModels;

import android.databinding.Bindable;
import android.text.TextUtils;

import com.android.databinding.library.baseAdapters.BR;
import com.opensoach.hpft.AppRepo.AppRepo;
import com.opensoach.hpft.Model.View.TaskItemDataModel;

/**
 * Created by Gregory Rasmussen on 7/26/17.
 */
public class TaskItemViewModel extends BaseViewModel {
    private TaskItemDataModel dataModel;


    public TaskItemViewModel(TaskItemDataModel dataModel) {
        this.dataModel = dataModel;
    }


    @Bindable
    public String getTitle() {
        return !TextUtils.isEmpty(dataModel.getTitle()) ? dataModel.getTitle() : "";
    }


    @Bindable
    public boolean getIsCompleted(){
        return  dataModel.getIsCompleted();
    }


    public String getObservationValue(){
        return dataModel.getObservationValue();
    }

    public void setObservationValue(String  observationValue){
        dataModel.setObservationValue(observationValue);

    }

    public String getComment(){
        return dataModel.getComment();
    }

    public void setComment(String  comment){
        dataModel.setComment(comment);

    }


    @Bindable
    public void setIsCompleted(boolean isCompleted){
         dataModel.setIsCompleted(isCompleted);

        notifyPropertyChanged(BR.isCompleted);

        if (isCompleted == true) {
             AppRepo.getInstance().getSelectedTaskDataViewModels().add(dataModel);
         }else{
             AppRepo.getInstance().getSelectedTaskDataViewModels().remove(dataModel);
         }

         if (AppRepo.getInstance().getSelectedTaskDataViewModels().size() > 0){
             MainViewModel.getInstance().getHeaderViewModel().setUploadEnabled(true);
         }else{
             MainViewModel.getInstance().getHeaderViewModel().setUploadEnabled(false);
         }
    }

}
