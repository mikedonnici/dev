import { Component, EventEmitter, Input, Output, OnInit } from '@angular/core';
import {UserService} from '../services/user.service';

@Component({
  selector: 'app-active-users',
  templateUrl: './active-users.component.html',
  styleUrls: ['./active-users.component.css']
})
export class ActiveUsersComponent implements OnInit {

  users: string[] = [];

  constructor(private usrSrvc: UserService) {
  }

  ngOnInit() {
    this.users = this.usrSrvc.activeUsers;
  }

  onSetToInactive(id: number) {
    this.usrSrvc.setToInactive(id);
  }
}
