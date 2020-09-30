import {Component, OnInit} from '@angular/core';
import {Ingredient} from '../../shared/ingredient.model';

@Component({
  selector: 'app-shopping-list',
  templateUrl: './shopping-list.component.html',
  styleUrls: ['./shopping-list.component.css']
})
export class ShoppingListComponent implements OnInit {

  ingredients: Ingredient[] = [
    new Ingredient('tomatoes', 4, ''),
    new Ingredient('basil', 1, 'bunch'),
    new Ingredient('pasata', 1, 'litre')
  ];

  constructor() {
  }

  ngOnInit(): void {
  }

}
