export class ActionTransactionResponse<T>{
    admissionid: number;
    patientconfid: number;
    txndate: string;
    txndata: T;
    actionname: string;
    updated_by: string;
    firstname: string;
    lastname: string;
    conftypecode: string;
}

export class TransactionInput {

}
export class ActionTransactionProcessedData {
    transactionkey: string;
    transactiondata: ActionTransactionResponse<ActionTransactionDataValue>[];
}

export class ActionTransactionDataValue {
    comment: string;
    value: string;
}
