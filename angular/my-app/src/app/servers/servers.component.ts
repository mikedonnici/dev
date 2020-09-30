import {Component, OnInit} from '@angular/core';
import {FormsModule} from '@angular/forms';

@Component({
  selector: 'app-servers',
  templateUrl: './servers.component.html',
  styleUrls: ['./servers.component.css']
})
export class ServersComponent implements OnInit {

  allowNewServer = false;
  serverCreationStatus = 'idle';
  newServerName = '';

  serverList: string[] = [];

  constructor() {
    setTimeout(() => {
      this.allowNewServer = true;
    }, 3000);
  }

  ngOnInit(): void {
  }

  onAddNewServer(): void {
    this.serverCreationStatus = `creating ${this.newServerName} ...`;
    setTimeout(() => {
      this.serverCreationStatus = `${this.newServerName} creation complete`;
      this.serverList.push(this.newServerName);
      this.newServerName = '';
      setTimeout(() => {
        this.serverCreationStatus = 'idle';
      }, 2000);
    }, 5000);
  }
}
