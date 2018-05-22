package spl.hkt.opensoach.splapp.view;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.webkit.JavascriptInterface;
import android.webkit.WebView;
import android.widget.ImageButton;
import android.widget.LinearLayout;

import spl.hkt.opensoach.splapp.R;
import spl.hkt.opensoach.splapp.handler.UserFeedbackClickHandler;
import spl.hkt.opensoach.splapp.helper.CommandConstants;
import spl.hkt.opensoach.splapp.helper.CommonHelper;
import spl.hkt.opensoach.splapp.helper.PacketHelper;
import spl.hkt.opensoach.splapp.manager.SendPacketManager;
import spl.hkt.opensoach.splapp.model.communication.DeviceChartDataModel;
import spl.hkt.opensoach.splapp.model.communication.PacketFeedbackDataModel;
import spl.hkt.opensoach.splapp.model.view.UserFeedbackModel;

public class ScreenSaverActivity extends AppCompatActivity {

    private String MY_JSON_EXAMPLE = "{\"name\":\"John Doe\",\"email\":\"jdoe@testco.com\"}";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_screen_saver);

        LinearLayout ll = (LinearLayout) findViewById(R.id.activity_screen_saver);

        ll.setClickable(true);

        ll.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                finish();
            }
        });

        ImageButton goodImgBtn = (ImageButton) findViewById(R.id.good_btn);
        UserFeedbackModel userFeedbackGoodModel =new     UserFeedbackModel();
        userFeedbackGoodModel.UserRating = 3;
        goodImgBtn.setOnClickListener(new UserFeedbackClickHandler(userFeedbackGoodModel) );


        ImageButton goodOkBtn = (ImageButton) findViewById(R.id.good_btn);
        UserFeedbackModel userFeedbackOKModel=new     UserFeedbackModel();
        userFeedbackOKModel.UserRating = 2;
        goodImgBtn.setOnClickListener(new UserFeedbackClickHandler(userFeedbackOKModel) );


        ImageButton goodBadBtn = (ImageButton) findViewById(R.id.good_btn);
        UserFeedbackModel userFeedbackBadModel=new     UserFeedbackModel();
        userFeedbackBadModel.UserRating = 1;
        goodImgBtn.setOnClickListener(new UserFeedbackClickHandler(userFeedbackBadModel) );


        //WebView webView = (WebView) findViewById(R.id.webView);
        //webView.loadUrl("file:///android_asset/sleep.gif");

        //webView.getSettings().setJavaScriptEnabled(true);
        //webView.addJavascriptInterface(this, "AndroidMainAct");
        //webView.loadUrl("file:///android_asset/mypage.html");

    }

    @JavascriptInterface
    public String getMyJSONData() {
        return MY_JSON_EXAMPLE;
    }

    @Override
    public void finish() {
        super.finish();
    }

    public void onCloseMeClicked(View view) {
        this.finish();
    }

}
