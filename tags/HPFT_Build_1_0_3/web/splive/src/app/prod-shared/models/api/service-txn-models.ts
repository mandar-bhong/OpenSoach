export class ServiceTxnRequest {
    spid: number;
    startdate: Date;
    enddate: Date;

}

export class ServiceInstanceTransactionResponse {
    servintxnid: number;
    servinid: number;
    fopcode: string;
    fopname: string;
    status: number;
    txndata: string;
    txndate: Date;
}
