import { Component, OnInit } from '@angular/core';
import { PatientService } from 'app/services/patient.service';
import { DataListRequest } from '../../../../../shared/models/api/data-list-models';
import { TransactionDetailsFilter, GroupTransaction } from 'app/models/api/transaction-details';
import { ActionTransactionResponse, ActionTransactionProcessedData, ActionTransactionDataValue } from 'app/models/api/transaction-details-response';

@Component({
  selector: 'app-transactions-details',
  templateUrl: './transaction-details.component.html',
  styleUrls: ['./transaction-details.component.css']
})
export class TransactionDetailsComponent implements OnInit {

  constructor(private patientService: PatientService) { }
  transactionResponse: ActionTransactionResponse<ActionTransactionDataValue>[] = [];
  groupTransaction: GroupTransaction[] = [];
  // groupTransactionMap = new Map<string, ActionTransactionResponse[]>();
  actionTransactionProcessedData: ActionTransactionProcessedData[] = [];
  panelOpenState = false;
  ngOnInit() {
    //  fucntion call for geting transaction data.

   // this.test();


  }
  // service fucntion for getting action transaction details.
  getActionTransaction() {
    const dataListRequest = new DataListRequest<TransactionDetailsFilter>();
    dataListRequest.orderdirection = 'asc';
    dataListRequest.page = 1;
    dataListRequest.limit = 100;
    dataListRequest.orderby = 'patientconfid';
    dataListRequest.filter = new TransactionDetailsFilter();
    dataListRequest.filter.admissionid = 1;

    this.patientService.getActionTransaction(dataListRequest).subscribe((transactionPayloadResponse) => {
      if (transactionPayloadResponse.issuccess) {
        if (transactionPayloadResponse.data != null) {
          transactionPayloadResponse.data.records.forEach((item: any) => {
            const ActionTransactionData = new ActionTransactionResponse<ActionTransactionDataValue>();
            Object.assign(ActionTransactionData, item);
            console.log('item', item.txndata);
            const txnJsonData = JSON.parse(item.txndata);
            ActionTransactionData.txndata = txnJsonData;
            this.transactionResponse.push(ActionTransactionData);
          });
        }
        console.log('transactionPayloadResponse', transactionPayloadResponse);
        const result = this.transactionResponse.reduce(function (r, a) {
          r[a.conftypecode] = r[a.conftypecode] || [];
          r[a.conftypecode].push(a);
          return r;
        }, Object.create(null));

        console.log('result', result);

        this.actionTransactionProcessedData.push({ transactionkey: 'Medicine', transactiondata: result.Medicine || [] });
        this.actionTransactionProcessedData.push({ transactionkey: 'Monitor', transactiondata: result.Monitor || [] });
        this.actionTransactionProcessedData.push({ transactionkey: 'Intake', transactiondata: result.Intake || [] });
        this.actionTransactionProcessedData.push({ transactionkey: 'Output', transactiondata: result.Output || [] });
        console.log(' this.actionTransactionProcessedData', this.actionTransactionProcessedData);
      }
    });






    // this.groupTransactionMap.set("Medicine", result.Medicine || []);
    // this.groupTransactionMap.set("Monitor", result.Monitor || []);
    // this.groupTransactionMap.set("Intake", result.Intake || []);
    // this.groupTransactionMap.set("Output", result.Output || []);
    // console.log('groupTransactionMap', this.groupTransactionMap);


  }
  test() {
    // for (let i = 0; i <= 3; i++) {
    //   const x = new ActionTransactionResponse<ActionTransactionDataValue>();
    //   x.actionname = 'Take Cipla 4 Times ' + i;
    //   x.firstname = 'Sarjerao';
    //   x.lastname = 'Ghadage';
    //   x.txndate = '2019-03-01T16:37:01Z'
    //   x.conftypecode = 'Medicine';
    //   this.transactionResponse.push(x);
    // }
    // for (let i = 0; i <= 3; i++) {
    //   const x = new ActionTransactionResponse<ActionTransactionDataValue>();
    //   x.actionname = 'intake' + i;
    //   x.firstname = 'Sarjerao';
    //   x.lastname = 'Ghadage';
    //   x.txndate = '2019-03-01T16:37:01Z'
    //   x.conftypecode = 'Intake';
    //   this.transactionResponse.push(x);
    // }
    // for (let i = 0; i <= 3; i++) {
    //   const x = new ActionTransactionResponse<ActionTransactionDataValue>();
    //   x.actionname = 'monitoring' + i;
    //   x.firstname = 'Sarjerao';
    //   x.lastname = 'Ghadage';
    //   x.txndate = '2019-03-01T16:37:01Z'
    //   x.conftypecode = 'Monitor';
    //   this.transactionResponse.push(x);
    // }
    // console.log(' this.transactionResponse', this.transactionResponse);
    this.getActionTransaction();
  }
  gotoTop() {
    console.log('goto top clicked');
    let el = document.getElementById('Medicine');
    console.log('el', el);
    el.scrollIntoView({ behavior: "smooth" });
  }
  // code block cehcking obejct is emppty.
  checkEmptyObjects(object): boolean {
    if (Object.keys(object).length > 0) {
      return true;
    } else {
      return false;
    }
  }
}
