import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';

import {IndexComponent} from './index/index.component';
import {OneComponent} from './one/one.component';
import {AComponent} from './one/a/a.component';
import {DogComponent} from './one/a/dog/dog.component';
import {BComponent} from './one/b/b.component';
import {TwoComponent} from './two/two.component';


const routes: Routes = [
  {path: '', component: IndexComponent},
  {path: 'one', component: OneComponent},
  {path: 'one/a', component: AComponent},
  {path: 'one/a/dog', component: DogComponent},
  {path: 'one/b', component: BComponent},
  {path: 'two', component: TwoComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
