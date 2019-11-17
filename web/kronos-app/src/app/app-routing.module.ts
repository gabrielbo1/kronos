import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';


const routes: Routes = [
  {
    path: '', 
    component: LoginComponent
  },
  {
    path: 'login', component: LoginComponent
  },
  {
    path: 'kronos',
    loadChildren: 'src/app/kronos/kronos.module#KronosModule'
  }
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
