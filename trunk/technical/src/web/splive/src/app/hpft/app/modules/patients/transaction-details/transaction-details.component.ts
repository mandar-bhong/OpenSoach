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
  actionTransactionProcessedData: ActionTransactionProcessedData[] = [];
  panelOpenState = false;
  ngOnInit() {

  }
  // service fucntion for getting action transaction details.
  getActionTransaction() {
    const dataListRequest = new DataListRequest<TransactionDetailsFilter>();
    dataListRequest.orderdirection = 'asc';
    dataListRequest.page = 1;
    dataListRequest.limit = 100;
    dataListRequest.orderby = 'patientconfid';
    dataListRequest.filter = new TransactionDetailsFilter();
    dataListRequest.filter.admissionid =this.patientService.admissionid;

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

  }
  test() {
    this.getActionTransaction();
  }

  gotoTop() {
    let el = document.getElementById('Medicine');
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
