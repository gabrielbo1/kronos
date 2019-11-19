import { Component, OnInit, AfterViewInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormControl } from '@angular/forms';


import { NotificationsService } from 'angular2-notifications';
import { Atendimento, Intervalo, StatusAtendimento } from 'src/app/model/atendimento';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
import { Empresa } from 'src/app/model/empresa';
import { HttpService } from 'src/app/model/httpclient';
import { DnsWebService } from 'src/app/model/dns';

export interface StatusSelect {
  value: StatusAtendimento;
  viewValue: string;
}

export interface ClienteSelect {
  value: Empresa,
  viewValue: string;
}

@Component({
  selector: 'app-atendimento',
  templateUrl: './atendimento.component.html',
  styleUrls: ['./atendimento.component.scss']
})
export class AtendimentoComponent implements OnInit, AfterViewInit {

  public dataInicio: FormControl = new FormControl(new Date());
  public horaInicio: string;
  public dataFim: FormControl = new FormControl(new Date());
  public horaFim: string;
  public statusSelects: StatusSelect[] = [
    {
      value: StatusAtendimento.Aberto, viewValue: 'Aberto'
    },
    {
      value: StatusAtendimento.Espera, viewValue: 'Espera'
    },
    {
      value: StatusAtendimento.Fechado, viewValue: 'Fechado'
    }
  ];
  public clientsSelects: ClienteSelect[] = [];


  public atendimento: Atendimento = new Atendimento();
  public atendimentos: Array<Atendimento> = new Array<Atendimento>();

  constructor(
    private httpClient: HttpClient,
    private notificationsService: NotificationsService) { }

  ngOnInit() {
  }

  ngAfterViewInit() {
    new HttpService<Array<Empresa>>(this.httpClient)
        .get(DnsWebService.EMPRESA, false, new Array<Empresa>(), e => {})
        .subscribe(empresas => {
          this.clientsSelects = [];
          empresas.forEach(emp => this.clientsSelects.push({value: emp, viewValue: emp.nome}));
        });
  }

  addEventDataInicio(event: MatDatepickerInputEvent<Date>) {
    const length: number = this.atendimento.horariosAtendimento.length;
    if (length === 0) {
      this.atendimento.horariosAtendimento.push(new Intervalo());
    }
    if (this.atendimento.horariosAtendimento[length - 1].dataInicio !== '' &&
      this.atendimento.horariosAtendimento[length - 1].dataFim !== '') {
      this.atendimento.horariosAtendimento.push(new Intervalo());
    }
    this.atendimento.horariosAtendimento[length - 1].dataInicio 
        = new Date(this.dataInicio.value).toISOString();
  }

  addEventDataFim(event: MatDatepickerInputEvent<Date>) {
    const length: number = this.atendimento.horariosAtendimento.length;
    if (length === 0) {
      this.onErrorMensage('Data de início.', 'Informe primeiro a data de início do atendimento/chamado.');
    }
    this.atendimento.horariosAtendimento[length-1].dataFim = new Date(this.dataFim.value).toISOString();
  }

  private onSucessMensage(title: string, msg: string) {
    const toast = this.notificationsService.success(title, msg, {
      timeOut: 3000,
      showProgressBar: true,
      pauseOnHover: true,
      clickToClose: true,
    });
  }

  private onErrorMensage(title: string, msg: string) {
    const toast = this.notificationsService.error(title, msg, {
      timeOut: 3000,
      showProgressBar: true,
      pauseOnHover: true,
      clickToClose: true,
    });
  }
}
