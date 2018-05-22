package spl.hkt.opensoach.splapp.model.communication;
import com.google.gson.annotations.SerializedName;

public  class  PacketHeaderModel{
  @SerializedName("crc")
  public String CRC;
  @SerializedName("category")
  public  int  Category;
  @SerializedName("commandid")
  public  int  CommandID  ;
  @SerializedName("seqid")
  public int  SeqID ;
  @SerializedName("locationid")
  public  int  LocationID  ;

}