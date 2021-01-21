import {EventEmitter, Injectable, Output} from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  activeUsers = ['Mike', 'Christie', 'Maia', 'Leo', 'Milo', 'Bob'];
  inactiveUsers = [];

  @Output() userAction = new EventEmitter<string>();

  setToInactive(id: number) {
    this.inactiveUsers.push(this.activeUsers[id]);
    this.activeUsers.splice(id, 1);
    this.userAction.emit('inactivation');
  }

  setToActive(id: number) {
    this.activeUsers.push(this.inactiveUsers[id]);
    this.inactiveUsers.splice(id, 1);
    this.userAction.emit('activation');
  }
}
