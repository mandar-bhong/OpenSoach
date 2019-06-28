package com.opensoach.hpft.Views;

import android.app.Activity;
import android.os.Bundle;
import android.support.design.widget.TextInputEditText;
import android.util.Log;
import android.view.View;
import android.view.WindowManager;
import android.webkit.JavascriptInterface;
import android.widget.Button;
import android.widget.RatingBar;
import android.widget.Toast;

import com.hsalf.smilerating.BaseRating;
import com.hsalf.smilerating.SmileRating;

import com.opensoach.hpft.R;
import com.opensoach.hpft.Handler.UserFeedbackClickHandler;
import com.opensoach.hpft.Model.View.UserFeedbackModel;

public class ScreenSaverActivity extends Activity {

    private String MY_JSON_EXAMPLE = "{\"name\":\"John Doe\",\"email\":\"jdoe@testco.com\"}";

    RatingBar ratingbar1;
    Button button;
    SmileRating smileRating;
    TextInputEditText txt_feedback;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.activity_screen_saver);


//        LinearLayout ll = (LinearLayout) findViewById(R.id.activity_screen_saver);
//
//        ll.setClickable(true);
//
//        ll.setOnClickListener(new View.OnClickListener() {
//            @Override
//            public void onClick(View v) {
//                finish();
//            }
//        });


        txt_feedback = (TextInputEditText)findViewById(R.id.txt_feedback);

        Button submitBtn = (Button) findViewById(R.id.btn_submit);

        submitBtn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                if (smileRating.getRating() == smileRating.NONE){
                    Toast.makeText(
                            v.getContext(),
                            v.getContext().getResources().getString(R.string.feed_back_input_required),
                            Toast.LENGTH_LONG).show();
                    return;
                }

                UserFeedbackModel userFeedbackOKModel=new     UserFeedbackModel();
                userFeedbackOKModel.UserRating = smileRating.getRating();
                userFeedbackOKModel.Comment = txt_feedback.getText().toString();
                new UserFeedbackClickHandler(userFeedbackOKModel).onClick(v);
                smileRating.setSelectedSmile(smileRating.NONE);
                txt_feedback.setText("");

                Toast.makeText(
                        v.getContext(),
                        v.getContext().getResources().getString(R.string.feed_back_success),
                        Toast.LENGTH_LONG).show();
            }
        });


        Button dismissBtn = (Button) findViewById(R.id.btn_dismiss);
        dismissBtn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                finish();
            }
        });

         smileRating = (SmileRating) findViewById(R.id.smile_rating);

        smileRating.setOnSmileySelectionListener(new SmileRating.OnSmileySelectionListener() {
            @Override
            public void onSmileySelected(@BaseRating.Smiley int smiley, boolean reselected) {
                // reselected is false when user selects different smiley that previously selected one
                // true when the same smiley is selected.
                // Except if it first time, then the value will be false.
                String TAG ="SmileRating";
                switch (smiley) {
                    case SmileRating.BAD:
                        Log.i(TAG, "Bad");
                        break;
                    case SmileRating.GOOD:
                        Log.i(TAG, "Good");
                        break;
                    case SmileRating.GREAT:
                        Log.i(TAG, "Great");
                        break;
                    case SmileRating.OKAY:
                        Log.i(TAG, "Okay");
                        break;
                    case SmileRating.TERRIBLE:
                        Log.i(TAG, "Terrible");
                        break;
                }
            }
        });

    }


    @Override
    public void onWindowFocusChanged(boolean hasFocus) {
        super.onWindowFocusChanged(hasFocus);

        if (hasFocus) {

            getWindow().getDecorView().setSystemUiVisibility(
                    View.SYSTEM_UI_FLAG_LAYOUT_STABLE
                            | View.SYSTEM_UI_FLAG_LAYOUT_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_LAYOUT_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_HIDE_NAVIGATION
                            | View.SYSTEM_UI_FLAG_FULLSCREEN
                            | View.SYSTEM_UI_FLAG_IMMERSIVE_STICKY);

            getWindow().addFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN);


            getWindow().setSoftInputMode(
                    WindowManager.LayoutParams.SOFT_INPUT_STATE_ALWAYS_HIDDEN
            );
        }
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
        //this.finish();
    }

}
