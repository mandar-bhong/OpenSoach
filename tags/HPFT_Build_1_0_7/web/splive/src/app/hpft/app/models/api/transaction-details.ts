import { ActionTransactionResponse, ActionTransactionDataValue } from "./transaction-details-response";

export class TransactionDetailsFilter {
    admissionid: number;
    conftypecode: string;
}
export class GroupTransaction {
    groupkey: string;
    grouplist: ActionTransactionResponse<ActionTransactionDataValue>[]
}