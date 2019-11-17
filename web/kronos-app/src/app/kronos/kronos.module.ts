import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PontoComponent } from './ponto/ponto.component';
import { AtendimentoComponent } from './atendimento/atendimento.component';
import { RelatorioComponent } from './relatorio/relatorio.component';
import { ConfiguracaoComponent } from './configuracao/configuracao.component';
import { KronosComponent } from './kronos.component';
import { KronosRoutingModule } from './kronos-routing.module';
import { MaterialModule } from '../material-module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SimpleNotificationsModule } from 'angular2-notifications';



@NgModule({
  declarations: [
    KronosComponent,
    PontoComponent, 
    AtendimentoComponent, 
    RelatorioComponent, 
    ConfiguracaoComponent
  ],
  imports: [
    CommonModule,
    KronosRoutingModule,
    MaterialModule,
    FormsModule, 
    ReactiveFormsModule,
    SimpleNotificationsModule.forRoot({
      position: ["top", "right"],
    }),
  ]
})
export class KronosModule { }
