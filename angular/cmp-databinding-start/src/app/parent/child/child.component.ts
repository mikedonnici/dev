import {Component, EventEmitter, Output} from '@angular/core';

@Component({
  selector: 'app-child',
  templateUrl: './child.component.html',
  styleUrls: ['./child.component.css']
})
export class ChildComponent {
  message: string;
  @Output() speak = new EventEmitter<string>();
  onSpeak() {
    this.speak.emit(this.message);
  }
}
