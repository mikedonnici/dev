import {Injectable} from '@angular/core';
import {UserService} from './user.service';

@Injectable({
  providedIn: 'root'
})
export class CounterService {

  actions = {activations: 0, inactivations: 0};

  constructor(private usrSrvc: UserService) {
  }

  activeUserCount(): number {
    return this.usrSrvc.activeUsers.length;
  }

  inactiveUserCount(): number {
    return this.usrSrvc.inactiveUsers.length;
  }

  logActivation() {
    console.log('activation');
    console.log(this.actions);
    this.actions.activations++;
  }

  logInactivation() {
    console.log('inactivation');
    console.log(this.actions);
    this.actions.inactivations++;
  }
}
