package com.opensoach.hospital.Helper;

import android.support.v7.app.AppCompatActivity;

import com.opensoach.hospital.DAL.DatabaseManager;
import com.opensoach.hospital.Model.DB.DBEnggPartTableRowModel;
import com.opensoach.hospital.Model.DB.DBJobCardTableRowModel;
import com.opensoach.hospital.Model.DB.DBLocationTableRowModel;
import com.opensoach.hospital.ViewModels.JobBriefViewModel;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Random;

/**
 * Created by Mandar on 9/8/2017.
 */

public class TestDataHelper {


   public static List<JobBriefViewModel> GenerateData(AppCompatActivity activity, int count) {
        List<JobBriefViewModel> jobBriefViewModels = new ArrayList<>();

//        for (int i = 0; i < count; i++) {
//            JobBriefViewModel jobBriefViewModel = new JobBriefViewModel();
//            //jobBriefViewModel.AppContext = this.getBaseContext();
//            jobBriefViewModel.ContextActivity = activity;
//            jobBriefViewModel.setPart("Part: " + getRandomString());
//            jobBriefViewModel.setCustomer("Customer: " + getRandomString());
//            jobBriefViewModel.setJob(getRandomString());
//
//            int colorInt = new Random().nextInt(3);
//            if (colorInt == 0) {
//                jobBriefViewModel.setStatusColor(activity.getResources().getColor(R.color.color_status_delayed_notstarted));
//                jobBriefViewModel.setStatus("Not Started");
//            } else if (colorInt == 1) {
//                jobBriefViewModel.setStatusColor(activity.getResources().getColor(R.color.color_status_delayed_started));
//                jobBriefViewModel.setStatus("Delayed Started");
//            } else {
//                jobBriefViewModel.setStatusColor(activity.getResources().getColor(R.color.color_status_inprogress));
//                jobBriefViewModel.setStatus("In Progress");
//            }
//
//            jobBriefViewModels.add(jobBriefViewModel);
//
//
//        }

        return jobBriefViewModels;
    }

    public static String getRandomString() {
        Random rand = new Random();
        int stringSize = rand.nextInt(600) + 1;
        //String tempString = "Ust use LinearLayout in place of TableRow where you need colspan, that's all. I think you need to wrap a layout around another one. Have one Layout list vertically, inside have another one (or in this case, two) list horizontally. I'm still finding it hard to nicely split interface to 50-50 portion in ust use LinearLayout in place of TableRow where you need colspan, that's all. I think you need to wrap a layout around another one. Have one Layout list vertically, inside have another one (or in this case, two) list horizontally. I'm still finding it hard to nicely split interface to 50-50 portion in";
        return (String.valueOf( stringSize));
    }


    public static void InsertTestDataIntoDatabase(){

        DBLocationTableRowModel dbLocationTableRowModel = new DBLocationTableRowModel();
        dbLocationTableRowModel.setLocationName("Location1");
        dbLocationTableRowModel.setLocationId(1);
        DatabaseManager.InsertRow(dbLocationTableRowModel);

        DBLocationTableRowModel dbLocationTableRowModel1 = new DBLocationTableRowModel();
        dbLocationTableRowModel.setLocationName("Location2");
        dbLocationTableRowModel.setLocationId(2);
        DatabaseManager.InsertRow(dbLocationTableRowModel);


        DBJobCardTableRowModel dbJobCardTableRowModel = new DBJobCardTableRowModel();
        dbJobCardTableRowModel.setJobCardId(1);
        dbJobCardTableRowModel.setPartId(1);
        //dbJobCardTableRowModel.setPartNumber("Part 1");
        dbJobCardTableRowModel.setPartCount(10);
        dbJobCardTableRowModel.setJobCode("JobCode-1");
        dbJobCardTableRowModel.setStartDate(new Date());
        dbJobCardTableRowModel.setEndDate(new Date());
        dbJobCardTableRowModel.setState(0);
        dbJobCardTableRowModel.setComments("This is test job card");

        DatabaseManager.InsertRow(dbJobCardTableRowModel);


        dbJobCardTableRowModel = new DBJobCardTableRowModel();
        dbJobCardTableRowModel.setJobCardId(2);
        dbJobCardTableRowModel.setPartId(2);
        //dbJobCardTableRowModel.setPartNumber("Part 2");
        dbJobCardTableRowModel.setPartCount(11);
        dbJobCardTableRowModel.setJobCode("JobCode-2");
        dbJobCardTableRowModel.setStartDate(new Date());
        dbJobCardTableRowModel.setEndDate(new Date());
        dbJobCardTableRowModel.setState(1);
        dbJobCardTableRowModel.setComments("This is test job card 2");

        DatabaseManager.InsertRow(dbJobCardTableRowModel);


        dbJobCardTableRowModel = new DBJobCardTableRowModel();
        dbJobCardTableRowModel.setJobCardId(3);
        dbJobCardTableRowModel.setPartId(3);
        dbJobCardTableRowModel.setPartCount(12);
        //dbJobCardTableRowModel.setPartNumber("Part 3");
        dbJobCardTableRowModel.setJobCode("JobCode-3");
        dbJobCardTableRowModel.setStartDate(new Date());
        dbJobCardTableRowModel.setEndDate(new Date());
        dbJobCardTableRowModel.setState(1);
        dbJobCardTableRowModel.setComments("This is test job card 2");

        DatabaseManager.InsertRow(dbJobCardTableRowModel);


        DBEnggPartTableRowModel dbEnggPartTableRowModel = new DBEnggPartTableRowModel();
        dbEnggPartTableRowModel.setPartId(1);
        dbEnggPartTableRowModel.setProcess("This is test Process for part 1");
        dbEnggPartTableRowModel.setInternalPartNo("INP1");
        dbEnggPartTableRowModel.setPartRevision("2");
        dbEnggPartTableRowModel.setPartNo("ABCD1");
        dbEnggPartTableRowModel.setToolJSON("This is tool 1");

        DatabaseManager.InsertRow(dbEnggPartTableRowModel);


        dbEnggPartTableRowModel = new DBEnggPartTableRowModel();
        dbEnggPartTableRowModel.setPartId(2);
        dbEnggPartTableRowModel.setProcess("This is test Process for part 2");
        dbEnggPartTableRowModel.setInternalPartNo("INP2");
        dbEnggPartTableRowModel.setPartRevision("3");
        dbEnggPartTableRowModel.setPartNo("ABCD2");
        dbEnggPartTableRowModel.setToolJSON("This is tool 2");

        DatabaseManager.InsertRow(dbEnggPartTableRowModel);


        dbEnggPartTableRowModel = new DBEnggPartTableRowModel();
        dbEnggPartTableRowModel.setPartId(2);
        dbEnggPartTableRowModel.setProcess("This is test Process for part 3");
        dbEnggPartTableRowModel.setInternalPartNo("INP3");
        dbEnggPartTableRowModel.setPartRevision("4");
        dbEnggPartTableRowModel.setPartNo("ABCD3");
        dbEnggPartTableRowModel.setToolJSON("This is tool 3");

        DatabaseManager.InsertRow(dbEnggPartTableRowModel);

    }

    public static void InsertJoBDataIntoDatabase(){
        DBJobCardTableRowModel dbJobCardTableRowModel = new DBJobCardTableRowModel();
        dbJobCardTableRowModel.setJobCardId(1);
        dbJobCardTableRowModel.setPartId(1);
        //dbJobCardTableRowModel.setPartNumber("Part 1");
        dbJobCardTableRowModel.setPartCount(10);
        dbJobCardTableRowModel.setJobCode(getRandomString());
        dbJobCardTableRowModel.setStartDate(new Date());
        dbJobCardTableRowModel.setEndDate(new Date());
        dbJobCardTableRowModel.setState(0);
        dbJobCardTableRowModel.setComments("This is test job card");

        DatabaseManager.InsertRow(dbJobCardTableRowModel);
    }


}
