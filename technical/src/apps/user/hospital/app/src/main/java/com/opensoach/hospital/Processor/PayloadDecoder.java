package com.opensoach.hospital.Processor;

import android.util.Log;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import com.opensoach.hospital.Helper.CommandConstants;
import com.opensoach.hospital.Model.Communication.PacketAuthCodeDataModel;
import com.opensoach.hospital.Model.Communication.PacketDeleteTableRowDataModel;
import com.opensoach.hospital.Model.Communication.PacketEnggPartToolsDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobCardsDataModel;
import com.opensoach.hospital.Model.Communication.PacketJobCardsStatusChangedDataModel;
import com.opensoach.hospital.Model.Communication.PacketLocationsDataModel;
import com.opensoach.hospital.Model.Communication.PacketModel;
import com.opensoach.hospital.Model.Communication.PacketPartDrawingsDataModel;
import com.opensoach.hospital.Model.Communication.PacketSimpleAckModel;
import com.opensoach.hospital.Model.PacketDecodeResultModel;
import com.opensoach.hospital.Utility.AppLogger;

import java.lang.reflect.Type;

/**
 * Created by Mandar on 8/27/2017.
 */

public class PayloadDecoder {
    public static void Decode(PacketDecodeResultModel packetDecodeResultModel, String packet) {

        packetDecodeResultModel.IsSuccess = false;
        try {

            //PacketChartConfigurationModel packetChartConfigurationModel;
            Log.i("PayloadDecoder", "CategoryID : " + packetDecodeResultModel.Packet.Header.Category + " CommandID : " + packetDecodeResultModel.Packet.Header.CommandID);
            switch (packetDecodeResultModel.Packet.Header.Category) {
                case CommandConstants.CMD_CAT_DEVICE_REG: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {
                        case CommandConstants.CMD_DEVICE_REGISTRATION: {
                            //TODO
                            break;
                        }
                    }

                    break;
                }
                case CommandConstants.CMD_CAT_CONFIG: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {

                        case CommandConstants.CMD_CONFIG_DEVICE_SYNC_COMPLETED: {
                            //TODO
                            break;
                        }

                        case CommandConstants.CMD_CONFIG_LOCATION_SYNC: {
                            TypeToken<PacketModel<PacketLocationsDataModel>> typeToken = new TypeToken<PacketModel<PacketLocationsDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new LocationDataProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }

                        case CommandConstants.CMD_CONFIG_JOB_CARD:{
                            TypeToken<PacketModel<PacketJobCardsDataModel>> typeToken = new TypeToken<PacketModel<PacketJobCardsDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new JobCardProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }

                        case CommandConstants.CMD_CONFIG_ENGG_PART:{
                                TypeToken<PacketModel<PacketEnggPartToolsDataModel>> typeToken = new TypeToken<PacketModel<PacketEnggPartToolsDataModel>>() {
                                };
                                Type packetType = typeToken.getType();
                                packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                                packetDecodeResultModel.Processor = new EnggPartProcessor();
                                packetDecodeResultModel.IsSuccess = true;
                            break;
                        }

                        case CommandConstants.CMD_CONFIG_PART_DRAWING: {
                            TypeToken<PacketModel<PacketPartDrawingsDataModel>> typeToken = new TypeToken<PacketModel<PacketPartDrawingsDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new PartDrawingProcessor();
                            packetDecodeResultModel.IsSuccess = true;

                            break;
                        }

                        case CommandConstants.CMD_CONFIG_SERVER_SYNC_COMPLETED: {
                            //TODO
                            packetDecodeResultModel.Processor = new ServerSynCompletedProcessor();
                            packetDecodeResultModel.IsSuccess = true;

                            break;
                        }

                        case CommandConstants.CMD_CONFIG_LOCATION_HCODE: {
                            TypeToken<PacketModel<PacketAuthCodeDataModel>> typeToken = new TypeToken<PacketModel<PacketAuthCodeDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new AuthCodeDataProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }
                    }
                    break;
                }
                case CommandConstants.CMD_CAT_DATA: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {
                        case CommandConstants.CMD_DATA_JOB_CARD_DATA: {
                            //TODO
                            break;
                        }
                        case CommandConstants.CMD_DATA_COMPLAINT_DATA: {
                            //TODO
                            break;
                        }
                    }
                    break;

                }
                case CommandConstants.CMD_CAT_ACK: {
                    Type packetType = new TypeToken<PacketModel<PacketSimpleAckModel>>() {
                    }.getType();
                    packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                    packetDecodeResultModel.Processor = new AcknowledgementProcessor();
                    packetDecodeResultModel.IsSuccess = true;
                    break;
                }

                case CommandConstants.CMD_CAT_DEVICE_STATUS: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {

                        case CommandConstants.CMD_DEVICE_STATUS_BATTERY_STAUS: {
                            //TODO
                            break;
                        }
                    }
                    break;
                }

                case CommandConstants.CMD_CAT_SERVER_DATA_PUSHED: {
                    switch (packetDecodeResultModel.Packet.Header.CommandID) {

                        case CommandConstants.CMD_SERVER_DATA_PUSHED_JOB_CARD: {
                            TypeToken<PacketModel<PacketJobCardsDataModel>> typeToken = new TypeToken<PacketModel<PacketJobCardsDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new NewJobCardProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }
                        case CommandConstants.CMD_SERVER_DATA_JOB_STATE_CHANGED:{
                            TypeToken<PacketModel<PacketJobCardsStatusChangedDataModel>> typeToken = new TypeToken<PacketModel<PacketJobCardsStatusChangedDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new JobCardStateChangeProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }
                        case CommandConstants.CMD_SERVER_DATA_TABLE_ROW_DELETED:{
                            TypeToken<PacketModel<PacketDeleteTableRowDataModel>> typeToken = new TypeToken<PacketModel<PacketDeleteTableRowDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new DeleteTableRowProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }
                        case CommandConstants.CMD_SERVER_DATA_PUSHED_JOB_CARD_UPDATED:{
                            TypeToken<PacketModel<PacketJobCardsDataModel>> typeToken = new TypeToken<PacketModel<PacketJobCardsDataModel>>() {
                            };
                            Type packetType = typeToken.getType();
                            packetDecodeResultModel.Packet.Payload = new Gson().fromJson(packet, packetType);
                            packetDecodeResultModel.Processor = new JobCardUpdateProcessor();
                            packetDecodeResultModel.IsSuccess = true;
                            break;
                        }
                    }
                    break;
                }
            }
        } catch (Exception ex) {
            //result.Error =new ErrorModel();
            packetDecodeResultModel.IsSuccess = false;
            AppLogger.getInstance().Log(ex,"Error occured in PayloadDecoder");
            //TODO: Set Error Model
            //TODO: Log exception error
        }

    }
}
