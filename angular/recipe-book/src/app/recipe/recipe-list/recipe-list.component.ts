import {Component, OnInit} from '@angular/core';
import {Recipe} from '../recipe.model';

@Component({
  selector: 'app-recipe-list',
  templateUrl: './recipe-list.component.html',
  styleUrls: ['./recipe-list.component.css']
})
export class RecipeListComponent implements OnInit {

  recipes: Recipe[] = [
    new Recipe('hommus', 'chickpea dip / spread', 'https://img.taste.com.au/5DdgePxp/w720-h480-cfill-q80/taste/2016/11/classic-hommus-82014-1.jpeg'),
    new Recipe('hommus', 'chickpea dip / spread', 'https://img.taste.com.au/5DdgePxp/w720-h480-cfill-q80/taste/2016/11/classic-hommus-82014-1.jpeg'),
  ];

  constructor() {
  }

  ngOnInit(): void {
  }

}
