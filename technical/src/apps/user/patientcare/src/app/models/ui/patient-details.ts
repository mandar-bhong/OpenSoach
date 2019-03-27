// export class PatientDetails {
// 	name: string;
// }
// export class JSONBaseDataModel<T> {
// 	version: number;
// 	data: T;
// }
// export class DataList {
// 	test: string;
// }
// export class PersonalHistoryModel{
// 	weight: string;
// 	weighttendency: string;
// 	alcohalquantity: string;
// 	alcohalcomment: string;
// 	smokingquantity: string;
// 	smokingcomment: string;
// 	other: string;
// }
export class DataList {
	test: string;
}
export class PersonalHistoryModel {
	medicaldetialsid: number;
	weight: WeightData;
	alcohol: AlcoholData;
	smoking: SmokData;
	other: string;
}
export class SmokData {
	applicable: boolean;
	quantity: string;
	remarks: string;
}

export class AlcoholData {
	applicable: boolean;
	quantity: string;
	remarks: string;
}

export class WeightData {
	weight: string;
	weight_tendency: string;
}

export class PatientDetailsModel {
	fname: string;
	lname: string;
	mob_no: string;
	age: string;
	blood_grp: string;
	gender: number;
}
export class AdmissionDetailsModel {
	patient_uuid: string;
	patient_reg_no: string;
	bed_no: string;
	sp_uuid: number;
	dr_incharge: number;
	admitted_on: string;
}