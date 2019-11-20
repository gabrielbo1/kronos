import { Component, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import * as moment from 'moment';
import { NotificationsService } from 'angular2-notifications';

import { Atendimento, Intervalo, StatusAtendimento } from 'src/app/model/atendimento';
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

  @ViewChild('dataInicioValor', { static: false })
  dataInicioValor: any;
  public dataInicio: Date = new Date();
  public horaInicio: string = '00:00'

  @ViewChild('dataFimValor', { static: false })
  dataFimValor: any;
  public dataFim: Date = new Date();
  public horaFim: string = '00:00'


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
  public msgError: string = '';

  private atendHttpClient: HttpService<Atendimento>;
  private atendsHttpClient: HttpService<Array<Atendimento>>;
  private empresasHttpClient: HttpService<Array<Empresa>>;

  constructor(
    private httpClient: HttpClient,
    private notificationsService: NotificationsService) {
    moment.locale('pt-BR');
  }

  ngOnInit() {
    if (this.atendimento.horariosAtendimento === undefined) {
      this.atendimento.horariosAtendimento = new Array<Intervalo>();
      this.atendimento.horariosAtendimento.push(new Intervalo());
      let length: number = this.atendimento.horariosAtendimento.length;
      this.atendimento.horariosAtendimento[length - 1].dataInicio = new Date().toISOString();
      this.atendimento.horariosAtendimento[length - 1].dataFim = new Date().toISOString();
    }
  }

  onChangeDataInicio(s: any) {
    this.dataInicio = moment(s.value, 'DD/MM/YYYY').toDate();
  }

  onChangeDataFim(s: any) {
    this.dataFim = moment(s.value, 'DD/MM/YYYY').toDate();
  }

  ngAfterViewInit() {
    this.atendHttpClient = new HttpService<Atendimento>(this.httpClient);
    this.atendsHttpClient = new HttpService<Array<Atendimento>>(this.httpClient);
    this.empresasHttpClient = new HttpService<Array<Empresa>>(this.httpClient);

    this.empresasHttpClient
      .get(DnsWebService.EMPRESA, false, new Array<Empresa>(), e => { })
      .subscribe(empresas => {
        this.clientsSelects = [];
        empresas.forEach(emp => this.clientsSelects.push({ value: emp, viewValue: emp.nome }));
      });
  }


  cadastro(atnd: Atendimento) {
    const length: number = atnd.horariosAtendimento.length;

    if (this.horaInicio === undefined || this.horaInicio === '00:00') {
      this.onErrorMensage('Erro', 'Informe a hora de início do atendimento.')
      return;
    } else {
      atnd.horariosAtendimento[length - 1].dataInicio
        = moment(this.dataInicio).format('YYYY-MM-DD') + ' ' + this.horaInicio;
    }

    if (this.atendimento.statusAtendimento === StatusAtendimento.Fechado &&
      (this.horaFim === undefined || this.horaFim === '00:00')) {
      this.onErrorMensage('Erro', 'Informe a hora de fim do atendimento.')
      return;
    } else {
      atnd.horariosAtendimento[length - 1].dataFim
        = moment(this.dataFim).format('YYYY-MM-DD') + ' ' + this.horaFim;
    }

    if (atnd.cliente === undefined) {
      this.onErrorMensage('Erro', 'Informe o cliente do atendimento.')
      return;
    }

    if (atnd.statusAtendimento === undefined) {
      this.onErrorMensage('Erro', 'Informe o status do atendimento.')
      return;
    }

    atnd.usuario = DnsWebService.usuario;

    this.atendHttpClient
      .post(DnsWebService.ATENDIMENTO, atnd, false, this.atendimento, e => {
        this.onErrorMensage(e.codigo, e.mensagem);
      })
      .subscribe(atndReturn => {
        if (atndReturn.id !== undefined && atndReturn.id !== 0) {
          this.limpar();
        }
        this.buscarAtendimentosIdUsuario();
      })
  }

  limpar() {
    this.dataInicio = new Date();
    this.horaInicio = '00:00'
    this.dataFim = new Date();
    this.horaFim = '00:00';
    this.atendimento = new Atendimento();
    this.atendimento.horariosAtendimento = new Array<Intervalo>();
  }

  private buscarAtendimentosIdUsuario() {
    this.atendsHttpClient
      .get(DnsWebService.ATENDIMENTO_USUARIO + '/' + DnsWebService.usuario.id,
        false, new Array<Atendimento>(), e1 => {
          this.onErrorMensage(e1.codigo, e1.mensagem);
        })
      .subscribe(atendReturs => {
        this.atendimentos = [];
        atendReturs.forEach((a) => this.atendimentos.push(a));
      })
  }

  private onSucessMensage(title: string, msg: string) {
    const toast = this.notificationsService.success(title, msg, {
      timeOut: 10000,
      showProgressBar: true,
      pauseOnHover: true,
      clickToClose: true,
    });
  }

  private onErrorMensage(title: string, msg: string) {
    const toast = this.notificationsService.error(title, '', {
      timeOut: 10000,
      showProgressBar: true,
      pauseOnHover: true,
      clickToClose: true,
    });
    this.msgError = msg;
    setTimeout(() => { this.msgError = ''; }, 10000);
  }

  limarMsgError() {
    this.msgError = '';
  }
}
