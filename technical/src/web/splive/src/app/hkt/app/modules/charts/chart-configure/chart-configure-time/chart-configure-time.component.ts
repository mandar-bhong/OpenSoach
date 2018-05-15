import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { DynamicContextService } from '../../../../../../shared/modules/dynamic-component-loader/dynamic-context.service';


@Component({
  selector: 'app-chart-configure-time',
  templateUrl: './chart-configure-time.component.html',
  styleUrls: ['./chart-configure-time.component.css']
})
export class ChartConfigureTimeComponent implements OnInit {

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

  save() {
    if (this.editableForm.invalid) { return; }

  }

  previousClick() {
    if (this.editableForm.invalid) { return; }
    this.dynamicContextService.changeComponent('app-chart-configure-basic');
  }

  nextClick() {
    if (this.editableForm.invalid) { return; }
    this.dynamicContextService.changeComponent('app-chart-configure-task');
  }
}
