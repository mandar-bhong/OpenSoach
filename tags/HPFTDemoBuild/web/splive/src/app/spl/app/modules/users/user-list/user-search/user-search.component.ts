import { Component, OnInit } from '@angular/core';

import { EnumDataSourceItem } from '../../../../../../shared/models/ui/enum-datasource-item';
import { AppUserService } from '../../../../../../shared/services/user/app-user.service';
import { UserFilterRequest } from '../../../../../app/models/api/user-models';
import { UserFilterModel } from '../../../../models/ui/user-models';
import { CorporateService } from '../../../../services/corporate.service';
import { UserService } from '../../../../services/user.service';

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
  constructor(private userService: UserService,
    private userSharedService: AppUserService,
    private corporateService: CorporateService) { }

  ngOnInit() {
    this.userStates = this.userSharedService.getUserStates();
    this.userCategories = this.userService.getUsersCategories();
  }
  search() {
    this.isExpanded = false;
    const userFilterRequest = new UserFilterRequest();
    this.dataModel.copyTo(userFilterRequest);
    this.userService.dataListSubjectTrigger(userFilterRequest);
  }
  panelOpened() {
    this.isExpanded = true;
  }
}
