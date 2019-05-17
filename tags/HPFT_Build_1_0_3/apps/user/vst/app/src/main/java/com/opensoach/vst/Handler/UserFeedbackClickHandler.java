package com.opensoach.vst.Handler;

import android.view.View;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.TimeZone;

import com.opensoach.vst.Helper.AppAction;
import com.opensoach.vst.Manager.SendPacketManager;
import com.opensoach.vst.Model.Communication.PacketFeedbackDataModel;
import com.opensoach.vst.Model.View.UserFeedbackModel;

import static com.opensoach.vst.Constants.ApplicationConstants.PACKET_DATE_FORMAT;

/**
 * Created by Mandar on 8/14/2017.
 */

public class UserFeedbackClickHandler implements View.OnClickListener {

    UserFeedbackModel feedbackModel;

    public UserFeedbackClickHandler(UserFeedbackModel userFeedbackModel) {
        feedbackModel = userFeedbackModel;
    }

    @Override
    public void onClick(View v) {

        SimpleDateFormat raiseOnDateFormat = new SimpleDateFormat(PACKET_DATE_FORMAT);

        raiseOnDateFormat.setTimeZone(TimeZone.getTimeZone("GMT"));
        PacketFeedbackDataModel packetFeedbackDataModel = new PacketFeedbackDataModel();
        packetFeedbackDataModel.Feedback = feedbackModel.UserRating;
        packetFeedbackDataModel.RaisedOn = raiseOnDateFormat.format(new Date());
        packetFeedbackDataModel.Comment = feedbackModel.Comment;

        ArrayList<PacketFeedbackDataModel> feedbacks = new ArrayList<>();
        feedbacks.add(packetFeedbackDataModel);

        SendPacketManager.Instance().send(AppAction.FEEDBACK_DATA, feedbacks);
    }
}
