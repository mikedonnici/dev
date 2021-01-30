import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {DroseraComponent} from './drosera/drosera.component';
import {DionaeaComponent} from './dionaea/dionaea.component';
import {DrosophyllumComponent} from './drosophyllum/drosophyllum.component';
import {NepenthesComponent} from './nepenthes/nepenthes.component';
import {SarraceniaComponent} from './sarracenia/sarracenia.component';
import {HomeComponent} from "./home/home.component";

const routes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'dionaea', component: DionaeaComponent},
  {path: 'drosera', component: DroseraComponent},
  {path: 'drosophyllum', component: DrosophyllumComponent},
  {path: 'nepenthes', component: NepenthesComponent},
  {path: 'sarracenia', component: SarraceniaComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
