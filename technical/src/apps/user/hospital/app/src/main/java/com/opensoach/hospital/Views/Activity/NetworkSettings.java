package com.opensoach.hospital.Views.Activity;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;

import com.opensoach.hospital.Helper.Constants;
import com.opensoach.hospital.Helper.SharedPreferencesHelper;
import com.opensoach.hospital.R;

public class NetworkSettings extends Activity {

    Button btnSubmit;
    EditText txtBoxIpAddress;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_network_settings);

        btnSubmit = (Button) findViewById(R.id.btnSubmit);
        txtBoxIpAddress = (EditText) findViewById(R.id.etIPAddress) ;


        String serverIPAddress = SharedPreferencesHelper.getInstance(this).getDataFromSharedPreference("ServerIPAddress");

        if(serverIPAddress != null && serverIPAddress != ""){
            txtBoxIpAddress.setText(serverIPAddress);
            Intent intent = new Intent(this.getBaseContext(), MainActivity.class);
            startActivity(intent);
        }

        btnSubmit.setOnClickListener(new View.OnClickListener() {

            @Override
            public void onClick(View view) {
                Constants.SERVER_HOST_OR_IP = txtBoxIpAddress.getText().toString();
                Constants.WEB_SOCKET_URL = "ws://"+Constants.SERVER_HOST_OR_IP+":8085/ws";;

                SharedPreferencesHelper.getInstance(view.getContext()).updateSharedPreference("ServerIPAddress",Constants.SERVER_HOST_OR_IP);

                Intent intent = new Intent(view.getContext(), MainActivity.class);
                startActivity(intent);
            }

        });
    }


}
