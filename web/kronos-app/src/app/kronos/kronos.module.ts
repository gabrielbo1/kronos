import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PontoComponent, ConfirmarPontoComponent } from './ponto/ponto.component';
import { AtendimentoComponent } from './atendimento/atendimento.component';
import { AtendimentoDeletetarDialogComponent } from './atendimento/atendimento.component';
import { RelatorioComponent } from './relatorio/relatorio.component';
import { ConfiguracaoComponent } from './configuracao/configuracao.component';
import { KronosComponent } from './kronos.component';
import { KronosRoutingModule } from './kronos-routing.module';
import { MaterialModule } from '../material-module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SimpleNotificationsModule } from 'angular2-notifications';
import { MaterialTimePickerModule } from '@candidosales/material-time-picker';


import {
  MAT_MOMENT_DATE_FORMATS,
  MomentDateAdapter,
  MAT_MOMENT_DATE_ADAPTER_OPTIONS,
} from '@angular/material-moment-adapter';
import {DateAdapter, MAT_DATE_FORMATS, MAT_DATE_LOCALE} from '@angular/material/core';


@NgModule({
  declarations: [
    KronosComponent,
    PontoComponent,
    ConfirmarPontoComponent, 
    AtendimentoComponent, 
    AtendimentoDeletetarDialogComponent,
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
    MaterialTimePickerModule,
  ],
  entryComponents: [AtendimentoComponent, 
                    AtendimentoDeletetarDialogComponent, 
                    PontoComponent, 
                    ConfirmarPontoComponent],
  providers: [
    {provide: MAT_DATE_LOCALE, useValue: 'pt-BR'},
    {
      provide: DateAdapter,
      useClass: MomentDateAdapter,
      deps: [MAT_DATE_LOCALE, MAT_MOMENT_DATE_ADAPTER_OPTIONS],
    },
    {provide: MAT_DATE_FORMATS, useValue: MAT_MOMENT_DATE_FORMATS,},
  ]
})
export class KronosModule { }
