package spl.hkt.opensoach.splapp.handler;

import android.view.View;

import spl.hkt.opensoach.splapp.apprepo.AppRepo;
import spl.hkt.opensoach.splapp.communication.CommunicationManager;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.view.UserFeedbackModel;

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

        PacketFeedbackDataModel packetFeedbackDataModel = new PacketFeedbackDataModel();
        packetFeedbackDataModel.Rating = feedbackModel.UserRating;

        //String packetJson1 = PacketHelper.GetFeedbackPacket(packetFeedbackDataModel);
        if(AppRepo.getInstance().IsServerConnected()) {
            String packetJson = PacketHelper.GetFeedbackPacket(packetFeedbackDataModel);
            CommunicationManager.getInstance().SendPacket(packetJson);
        }else{
            //TODO: Log Message
        }
    }
}
