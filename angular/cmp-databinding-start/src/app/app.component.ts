import {Component} from "@angular/core"
import {ServerData} from "./types"

@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.css"]
})
export class AppComponent {

  serverElements = []

  onServerAdded(data: ServerData) {
    this.serverElements.push({
      type: "server",
      name: data.name,
      content: data.content
    })
  }

  onBlueprintAdded(data: ServerData) {
    this.serverElements.push({
      type: "blueprint",
      name: data.name,
      content: data.content
    })
  }

}
