import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';

@Component({
  selector: 'app-server-manager',
  templateUrl: './server-manager.component.html',
  styleUrls: ['./server-manager.component.css']
})
export class ServerManagerComponent implements OnInit {

  serverList = [];
  newServerName = '';
  // newServerContent = '';
  @ViewChild('newServerContent') newServerContent: ElementRef;

  constructor() {
  }

  ngOnInit(): void {
  }

  onAddServer(serverName: HTMLInputElement) {
    this.serverList.push({
      type: 'server',
      name: serverName.value,
      content: this.newServerContent.nativeElement.value
    });
  }

  onAddBlueprint(serverName: HTMLInputElement) {
    this.serverList.push({
      type: 'blueprint',
      name: serverName,
      content: this.newServerContent.nativeElement.value
    });
  }

}
