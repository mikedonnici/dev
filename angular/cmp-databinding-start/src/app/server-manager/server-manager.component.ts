import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-server-manager',
  templateUrl: './server-manager.component.html',
  styleUrls: ['./server-manager.component.css']
})
export class ServerManagerComponent implements OnInit {

  serverList = [];
  newServerName = '';
  newServerContent = '';

  constructor() {
  }

  ngOnInit(): void {
  }

  onAddServer() {
    this.serverList.push({
      type: 'server',
      name: this.newServerName,
      content: this.newServerContent
    });
  }

  onAddBlueprint() {
    this.serverList.push({
      type: 'blueprint',
      name: this.newServerName,
      content: this.newServerContent
    });
  }

}
