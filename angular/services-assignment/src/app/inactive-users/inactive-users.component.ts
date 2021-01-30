import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {UserService} from '../services/user.service';

@Component({
  selector: 'app-inactive-users',
  templateUrl: './inactive-users.component.html',
  styleUrls: ['./inactive-users.component.css']
})
export class InactiveUsersComponent implements OnInit {

  users: string[] = [];

  constructor(private usrSrvc: UserService) {
  }

  ngOnInit() {
    this.users = this.usrSrvc.inactiveUsers;
  }

  onSetToActive(id: number) {
    this.usrSrvc.setToActive(id);
  }
}
