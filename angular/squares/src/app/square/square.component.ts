import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-square',
  templateUrl: './square.component.html',
  styleUrls: ['./square.component.css']
})
export class SquareComponent implements OnInit {

  @Input() colour: string;
  @Output() squareClicked = new EventEmitter<string>();

  clrStyles = {
    red: 'bg-danger',
    orange: 'bg-warning',
    blue: 'bg-info',
    grey: 'bg-secondary',
    green: 'bg-success'
  };

  constructor() { }

  ngOnInit(): void {
  }

  onSquareClick(data): void {
    this.squareClicked.emit(data);
  }

}
