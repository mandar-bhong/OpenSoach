package spl.hkt.opensoach.splapp.handler;

import android.view.View;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.TimeZone;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.communication.CommunicationManager;
import spl.hkt.opensoach.splapp.helper.AppAction;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.view.UserFeedbackModel;

import static spl.hkt.opensoach.splapp.helper.ApplicationConstants.PACKET_DATE_FORMAT;

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

        ArrayList<PacketFeedbackDataModel> feedbacks = new ArrayList<>();
        feedbacks.add(packetFeedbackDataModel);

        SendPacketManager.Instance().send(AppAction.FEEDBACK_DATA, feedbacks);
    }
}
