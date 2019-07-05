import { Component, OnInit } from '@angular/core';

import { EnumDataSourceItem } from '../../../../../shared/models/ui/enum-datasource-item';
import { AppUserService } from '../../../../../shared/services/user/app-user.service';
import { UserFilterRequest } from '../../../../models/api/user-models';
import { UserFilterModel } from '../../../../models/ui/user-models';
import { ProdUserService } from '../../../../services/user/prod-user.service';

@Component({
  selector: 'app-user-search',
  templateUrl: './user-search.component.html',
  styleUrls: ['./user-search.component.css']
})
export class UserSearchComponent implements OnInit {
  dataModel = new UserFilterModel();
  isExpanded = false;
  userStates: EnumDataSourceItem<number>[];
  userCategories: EnumDataSourceItem<number>[];
  constructor(private prodUserService: ProdUserService,
    public appUserService: AppUserService) { }

  ngOnInit() {
    this.userStates = this.appUserService.getUserStates();
  }
  search() {
    this.isExpanded = false;
    const userFilterRequest = new UserFilterRequest();
    this.dataModel.copyTo(userFilterRequest);
    this.prodUserService.dataListSubjectTrigger(userFilterRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }
}
