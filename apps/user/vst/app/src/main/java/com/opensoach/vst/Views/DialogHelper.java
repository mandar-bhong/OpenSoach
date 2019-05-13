package com.opensoach.vst.Views;

import android.app.AlertDialog;
import android.content.Context;
import android.content.DialogInterface;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import com.opensoach.vst.R;

/**
 * Created by samir.s.bukkawar on 4/12/2017.
 */

public class DialogHelper {

    public interface DialogCallBack {
        public boolean onSucess(String strData);

        public void onSucess(String strData1, String strData2);

        public void onSucess(String strData1, String strData2, String strData3);
    }

    public static void showSingleLineEditTextAlert(Context context, String title, final DialogCallBack callback) {

        LayoutInflater layoutInflater = LayoutInflater.from(context);
        View promptsView = layoutInflater.inflate(R.layout.dialog_input_singleline, null);
        TextView tvTitle = (TextView) promptsView.findViewById(R.id.dialogTitle);
        tvTitle.setText(title);

        AlertDialog.Builder alertDialogBuilder = new AlertDialog.Builder(context);
        alertDialogBuilder.setView(promptsView);
        final EditText userInput = (EditText) promptsView.findViewById(R.id.editTextDialogUserInput);
        final Context lContext = context;

        // set dialog message
        alertDialogBuilder
                .setCancelable(false)
                .setPositiveButton(context.getResources().getString(R.string.btn_ok),
                        new DialogInterface.OnClickListener() {
                            public void onClick(DialogInterface dialog, int id) {
                            }
                        })
                .setNegativeButton(context.getResources().getString(R.string.btn_cancel),
                        new DialogInterface.OnClickListener() {
                            public void onClick(DialogInterface dialog, int id) {
                                dialog.cancel();
                            }
                        });
        final AlertDialog alertDialog = alertDialogBuilder.create();
        alertDialog.show();

        //Overriding the handler immediately after show to prevent the defuals dismiss behaviour
        alertDialog.getButton(AlertDialog.BUTTON_POSITIVE).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (userInput.getText() != null && userInput.getText().length() != 0) {
                  boolean isSuccess =  callback.onSucess(userInput.getText().toString());

                  if(isSuccess) {
                      alertDialog.dismiss();
                  }
                } else {
                    Toast.makeText(
                            lContext,
                            lContext.getResources().getString(R.string.enter_valid_authcode),
                            Toast.LENGTH_LONG).show();
                }

            }
        });
    }

    public static void showComplaintDialog(Context context, String title, final DialogCallBack callback) {

        LayoutInflater layoutInflater = LayoutInflater.from(context);
        View promptsView = layoutInflater.inflate(R.layout.dialog_raise_complaint, null);
        // TextView tvTitle = (TextView) promptsView.findViewById(R.id.dialogTitle);
        //tvTitle.setText(title);

        AlertDialog.Builder alertDialogBuilder = new AlertDialog.Builder(context);
        alertDialogBuilder.setView(promptsView);
        final EditText etComplaintBy = (EditText) promptsView.findViewById(R.id.etComplaintBy);
        final EditText etComplaintTitle = (EditText) promptsView.findViewById(R.id.etComplaintTitle);
        final EditText etDescription = (EditText) promptsView.findViewById(R.id.etComplaintDescription);

        // set dialog message
        alertDialogBuilder
                .setCancelable(false)
                .setPositiveButton(context.getResources().getString(R.string.btn_ok),
                        new DialogInterface.OnClickListener() {
                            public void onClick(DialogInterface dialog, int id) {
                                if (etComplaintBy.getText() != null && etComplaintTitle.getText() != null) {
                                    callback.onSucess(etComplaintBy.getText().toString(),etComplaintTitle.getText().toString(), etDescription.getText().toString());
                                }
                                dialog.dismiss();
                            }
                        })
                .setNegativeButton(context.getResources().getString(R.string.btn_cancel),
                        new DialogInterface.OnClickListener() {
                            public void onClick(DialogInterface dialog, int id) {
                                dialog.cancel();
                            }
                        });
        AlertDialog alertDialog = alertDialogBuilder.create();
        alertDialog.show();
    }

}
