import {Component, ElementRef, EventEmitter, OnInit, Output, ViewChild} from "@angular/core"
import {ServerData} from "../types"

@Component({
  selector: "app-cockpit",
  templateUrl: "./cockpit.component.html",
  styleUrls: ["./cockpit.component.css"]
})
export class CockpitComponent implements OnInit {

  @Output() serverAdded = new EventEmitter<ServerData>()
  @Output() blueprintAdded = new EventEmitter<ServerData>()

  // ViewChild via a template reference, ie #serverContent
  @ViewChild("serverContent", {static: true}) serverContent: ElementRef

  constructor() {
  }

  ngOnInit(): void {
  }

  onAddServer(serverName: HTMLInputElement) {
    this.serverAdded.emit({name: serverName.value, content: this.serverContent.nativeElement.value})
  }

  onAddBlueprint(serverName: HTMLInputElement) {
    this.blueprintAdded.emit({name: serverName.value, content: this.serverContent.nativeElement.value})
  }

}
