import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';

@Component({
  selector: 'app-chart-configure-preview',
  templateUrl: './chart-configure-preview.component.html',
  styleUrls: ['./chart-configure-preview.component.css']
})
export class ChartConfigurePreviewComponent implements OnInit {

  editableForm: FormGroup;
  constructor(private dynamicContextService: DynamicContextService) {
  }

  ngOnInit() {
    this.createControls();
  }

  createControls(): void {
    this.editableForm = new FormGroup({
      configNameControl: new FormControl('', [Validators.required]),
      categoryControl: new FormControl('', [Validators.required])
    });
  }

  previousClick() {
    if (this.editableForm.invalid) { return; }
    this.dynamicContextService.changeComponent('app-chart-configure-task');
  }

  save() {
    if (this.editableForm.invalid) { return; }
  }
}

