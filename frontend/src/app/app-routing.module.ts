import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { SearchComponent } from './pages/search/search.component';
import { IndicesComponent } from './pages/indices/indices.component';
import { EntriesComponent } from './pages/entries/entries.component';

const routes: Routes = [
  { path: "dashboard", component: DashboardComponent },
  { path: "indices", component: IndicesComponent },
  { path: "entries", component: IndicesComponent },
  { path: 'entries/:name', component: EntriesComponent },
  { path: "search", component: SearchComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
