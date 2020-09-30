import {Component} from '@angular/core';
import {timestamp} from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'directives';
  pHidden = true;
  loggedClicks: string[] = [];

  logClick(): void {
    this.loggedClicks.push(Date.now().toString());
  }
}
