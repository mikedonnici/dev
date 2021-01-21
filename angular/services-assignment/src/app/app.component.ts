import {DoCheck, Component} from '@angular/core';
import {CounterService} from './services/counter.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements DoCheck {

  activeCount = 0;
  inactiveCount = 0;
  activations = 0;
  inactivations = 0;

  constructor(private countSrvc: CounterService) {
  }

  ngDoCheck() {
    this.activeCount = this.countSrvc.activeUserCount();
    this.inactiveCount = this.countSrvc.inactiveUserCount();
    this.activations = this.countSrvc.actions.activations;
    this.inactivations = this.countSrvc.actions.inactivations;
  }
}
