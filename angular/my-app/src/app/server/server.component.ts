import {Component} from '@angular/core';

@Component({
  selector: 'app-server',
  templateUrl: './server.component.html'
})

export class ServerComponent {
  serverStatusCssClass(): string {
    const s = Math.random();
    if (s > 0.5) {
      return 'alert-success';
    } else {
      return 'alert-danger';
    }
  }
}
