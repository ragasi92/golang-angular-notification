import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomePageComponent } from './components/home-page/home-page.component';


const routes: Routes = [
  { path: '', component: HomePageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }


// import { NgModule } from '@angular/core';
// import { Routes, RouterModule } from '@angular/router';
// import { RoleGuard } from './core/role.guard';
// import { AppComponent } from './app.component';
// import { PageNotFoundComponent } from './shared/page-not-found/page-not-found.component';

// const routes: Routes = [
//   { path: '', component: AppComponent, canActivate: [RoleGuard] },
//   { path: 'admin', loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule) },
//   { path: '**', component: PageNotFoundComponent },
//   { path: 'not-found', component: PageNotFoundComponent }
// ];

// @NgModule({
//   imports: [RouterModule.forRoot(routes, { scrollPositionRestoration: 'enabled' })],
//   exports: [RouterModule]
// })
// export class AppRoutingModule {
// }
