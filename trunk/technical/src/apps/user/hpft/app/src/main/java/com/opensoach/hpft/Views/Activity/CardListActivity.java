package com.opensoach.hpft.Views.Activity;




import android.app.FragmentManager;
import android.databinding.DataBindingUtil;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.view.WindowManager;

import com.opensoach.hpft.Model.AppNotificationModelBase;
import com.opensoach.hpft.R;
import com.opensoach.hpft.ViewModels.CardBriefViewModel;
import com.opensoach.hpft.ViewModels.CardGridViewModel;
import com.opensoach.hpft.ViewModels.CardListViewModel;
import com.opensoach.hpft.ViewModels.MainViewModel;
import com.opensoach.hpft.ViewModels.MedicalDetailsViewModel;
import com.opensoach.hpft.ViewModels.PatientDetailsViewModel;
import com.opensoach.hpft.ViewModels.TaskDataViewModel;
import com.opensoach.hpft.ViewModels.TaskDetailsViewModel;
import com.opensoach.hpft.ViewModels.TaskTimeDataViewModel;
import com.opensoach.hpft.Views.Fragment.HeaderFragment;
import com.opensoach.hpft.Views.Interfaces.IFragment;
import com.opensoach.hpft.Views.Interfaces.IUIUpdateEvent;
import com.opensoach.hpft.databinding.ActivityCardListBinding;

import java.util.ArrayList;
import java.util.Date;


public class CardListActivity extends AppCompatActivity
        implements IFragment<CardListViewModel>,IUIUpdateEvent,HeaderFragment.OnFragmentInteractionListener  {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        MainViewModel.getInstance().ContextActivity = this;

        setContentView(R.layout.activity_card_list);
        setDataContext(MainViewModel.getInstance().getCardListViewModel());


        //TODO: This step is importent for adding fragment into activity
        android.support.v4.app.FragmentManager fm = getSupportFragmentManager();
        fm.beginTransaction().replace(R.id.headerPlace, HeaderFragment.newInstance("","")).commit();

        hideSoftKeyboard();

    }

    public void hideSoftKeyboard() {
        getWindow().setSoftInputMode(WindowManager.LayoutParams.SOFT_INPUT_STATE_HIDDEN);
    }


    @Override
    public  void setDataContext(CardListViewModel viewModel){
        ActivityCardListBinding activityMainBinding = DataBindingUtil.setContentView(this, R.layout.activity_card_list);

        //viewModel.AppContext = this.getBaseContext();
        viewModel.ContextActivity = this;
        viewModel.getCardGridViewModel().ContextActivity = this;

        CardGridViewModel cardGridViewModel = viewModel.getCardGridViewModel();

        ArrayList list = new ArrayList<CardBriefViewModel>();

        list.add(GenerateData(this,1));
        list.add(GenerateData(this,2));
        list.add(GenerateData(this,3));


        cardGridViewModel.setItemsSource(list);

       // cardGridViewModel.getItemsSource().add(new CardBriefViewModel());

        //jobGridViewModel.setItemsSource(GenerateData());

//        viewModel.setGridViewModel(jobGridViewModel);
//
//        viewModel.setHeaderViewModel(new HeaderViewModel());

        activityMainBinding.setDataContext(cardGridViewModel);
    }

    @Override
    public CardListViewModel getDataContext() {
        return null;
    }

    @Override
    public void OnUIUpdateEvent(final AppNotificationModelBase model) {
        switch (model.DataProcessStatergyID) {

        }
    }





    @Override
    public void onFragmentInteraction(Uri uri){

    }

    private CardBriefViewModel GenerateData(AppCompatActivity ctx,int index){
        CardBriefViewModel cardBriefViewModel = new CardBriefViewModel();
        cardBriefViewModel.ContextActivity = ctx;

        PatientDetailsViewModel patientDetailsViewModel = new PatientDetailsViewModel();
        MedicalDetailsViewModel medicalDetailsViewModel = new MedicalDetailsViewModel();
        TaskDetailsViewModel taskDetailsViewModel =new TaskDetailsViewModel();

        patientDetailsViewModel.setAge(25+index);
        patientDetailsViewModel.setName("Patient-"+index);
        patientDetailsViewModel.setEmergencyContactNo("9898989-"+index);
        patientDetailsViewModel.setRegNo("89898-"+index);
        patientDetailsViewModel.setAdmissionDate(new Date(5000+index));


        medicalDetailsViewModel.setAllergies("Allergies-"+index);
        medicalDetailsViewModel.setHistory("Medical History-"+index);
        medicalDetailsViewModel.setTreatment("Treatment-"+index);


        taskDetailsViewModel.setTaskDataViewModel(new TaskDataViewModel());
        taskDetailsViewModel.setTaskTimeDataViewModel(new TaskTimeDataViewModel());
        taskDetailsViewModel.getTaskTimeDataViewModel().setUp();
        taskDetailsViewModel.getTaskDataViewModel().setUp();
        taskDetailsViewModel.setTitle("This is test for databind ele");
        taskDetailsViewModel.ContextActivity = ctx;

        cardBriefViewModel.setPatientDetails(patientDetailsViewModel);
        cardBriefViewModel.setMedicalDetails(medicalDetailsViewModel);
        cardBriefViewModel.setTaskDetails(taskDetailsViewModel);
        return cardBriefViewModel;
    }
}
