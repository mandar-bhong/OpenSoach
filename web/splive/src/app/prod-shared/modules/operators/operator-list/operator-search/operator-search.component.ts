import { Component, OnInit } from '@angular/core';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { OperatorFiltrRequest } from '../../../../models/api/operator-models';
import { OperatorFilterModel } from '../../../../models/ui/operator-models';
import { ProdOperatorService } from '../../../../services/operator/prod-operator.service';

@Component({
  selector: 'app-operator-search',
  templateUrl: './operator-search.component.html',
  styleUrls: ['./operator-search.component.css']
})
export class OperatorSearchComponent implements OnInit {
  dataModel = new OperatorFilterModel();
  isExpanded = false;
  operatorStates: EnumDataSourceItem<number>[];
  constructor(private prodOperatorService: ProdOperatorService) { }

  ngOnInit() {
    this.operatorStates = this.prodOperatorService.getOperatorStates();
  }
  search() {
    this.isExpanded = false;
    const operatorFiltrRequest = new OperatorFiltrRequest();
    this.dataModel.copyTo(operatorFiltrRequest);
    this.prodOperatorService.dataListSubjectTrigger(operatorFiltrRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }
}
