import { Component, OnInit, Output, EventEmitter, ViewChild } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { EditRecordBase, EDITABLE_RECORD_STATE, FORM_MODE } from '../../../../../../shared/views/edit-record-base';
import { PathologyReportAddRequest } from 'app/models/api/pathology-report-add-request';
import { GetUUID } from '../../../../../../shared/helpers/get-uuid';
import { PatientService } from 'app/services/patient.service';
import { FloatingButtonMenuService } from '../../../../../../shared/services/floating-button-menu.service';
import { AppNotificationService } from '../../../../../../shared/services/notification/app-notification.service';

@Component({
  selector: 'app-add-pathology-report',
  templateUrl: './add-pathology-report.component.html',
  styleUrls: ['./add-pathology-report.component.css']
})
export class AddPathologyReportComponent extends EditRecordBase implements OnInit {
  @Output() restForm = new EventEmitter();
  @ViewChild('fileUploadInput') fileUploadInputVariable: any;
  closeForm() {
    throw new Error("Method not implemented.");
  }
  uploadDocumentsIsEnable: boolean;
  doc_Array: File;
  uploadedDocuments: any[] = [];
  documentFile: any = null;
  constructor(private patientService: PatientService,
    private appNotificationService: AppNotificationService) {
    super();
  }

  ngOnInit() {
    this.createControls();
    this.iconCss = 'medical-icon-i-pathology'
    this.pageTitle = "Pathology Report";
    this.subTitle = 'Add Pathology Report Details';
    this.recordState = EDITABLE_RECORD_STATE.ADD;
    this.setFormMode(FORM_MODE.EDITABLE);
    this.showBackButton = false;
    // this.callbackUrl = params['callbackurl'];
  }
  createControls(): void {
    this.editableForm = new FormGroup({
      testperformed: new FormControl('', [Validators.required]),
      testresult: new FormControl('', [Validators.required]),
      comments: new FormControl(),
      testperformedtime: new FormControl('', [Validators.required])
    });
  }
  // code block for save 
  onFormSubmit() {
    if (this.editableForm.valid) {
      const pathologyReportAddRequest = new PathologyReportAddRequest();
      pathologyReportAddRequest.testperformed = this.editableForm.get('testperformed').value;
      pathologyReportAddRequest.testresult = this.editableForm.get('testresult').value;
      pathologyReportAddRequest.comments = this.editableForm.get('comments').value;
      pathologyReportAddRequest.testperformedtime = this.editableForm.get('testperformedtime').value;      
      const uuid = new GetUUID();
      pathologyReportAddRequest.uuid = uuid.getUuid();
      pathologyReportAddRequest.admissionid = this.patientService.admissionid;
      pathologyReportAddRequest.documentuuidlist = [];
      if (this.uploadedDocuments.length > 0) {
        this.uploadedDocuments.forEach((item) => {
          console.log('item.documentid', item);
          pathologyReportAddRequest.documentuuidlist.push(item.documentid);
        });
      }
      this.patientService.addReportData(pathologyReportAddRequest).subscribe(reportsPayloadResponse => {
        if (reportsPayloadResponse) {
          if (reportsPayloadResponse.issuccess) {
            this.appNotificationService.success('Report Added Successfully');
            this.restForm.emit(0);
          }
        }
      });

    }
  }
  uploadDocuments(): void {
    const formdata = new FormData();
    // code block for appending value to form data

    // formdata.append('UUID', uuid.getUuid());
    console.log('this.doc_Array', this.doc_Array);
    formdata.append('file', this.doc_Array, this.doc_Array.name);
    console.log(' formdata', formdata);
    this.patientService.uploadReportsDocuments(formdata).subscribe(documentsPayloadResponse => {
      if (documentsPayloadResponse) {
        if (documentsPayloadResponse.issuccess) {
          this.appNotificationService.success('Document Uploaded Successfully');
          // tslint:disable-next-line:max-line-length
          this.uploadedDocuments.push({ 'documentid': documentsPayloadResponse.data.uuid, 'documentname': this.doc_Array.name });
          this.doc_Array = null;
          this.fileUploadInputVariable.nativeElement.value = '';
        }
      }
    });
  }

  onChange(event: any) {
    this.documentFile = [].slice.call(event.target.files);
    this.doc_Array = this.documentFile[0];
    if (this.documentFile.length > 0) {
      // document size validations
    }
    console.log('this.documentFile', this.documentFile);
  }// end of fucntion.

  cancle() {
    // emiting 1 for rest.
    this.restForm.emit(1);
  }
}

