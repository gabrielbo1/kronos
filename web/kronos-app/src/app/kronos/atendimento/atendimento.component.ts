import { Component, OnInit, AfterViewInit, ViewChild } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import * as moment from 'moment';
import { NotificationsService } from 'angular2-notifications';

import { Atendimento, Intervalo, StatusAtendimento } from 'src/app/model/atendimento';
import { Pagina } from 'src/app/model/pagina';
import { Empresa } from 'src/app/model/empresa';
import { HttpService } from 'src/app/model/httpclient';
import { DnsWebService } from 'src/app/model/dns';
import { MatTableDataSource } from '@angular/material/table';
import { SelectionModel } from '@angular/cdk/collections';
import { PageEvent } from '@angular/material/paginator';

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

  displayedColumns: string[] = ['ID', 'Usuario', 'Cliente', 'Status'];
  dataSource = new MatTableDataSource(this.atendimentos);
  selection = new SelectionModel<Atendimento>(true, []);

  // MatPaginator Inputs
  length = 100;
  pageSize = 10;
  pageSizeOptions: number[] = [5, 10, 25, 100];
  // MatPaginator Output
  pageEvent: PageEvent;


  private atendPagHttpClient: HttpService<Pagina<Atendimento>>;
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
    this.atendPagHttpClient = new HttpService<Pagina<Atendimento>>(this.httpClient);
    this.atendHttpClient = new HttpService<Atendimento>(this.httpClient);
    this.atendsHttpClient = new HttpService<Array<Atendimento>>(this.httpClient);
    this.empresasHttpClient = new HttpService<Array<Empresa>>(this.httpClient);

    this.empresasHttpClient
      .get(DnsWebService.EMPRESA, false, new Array<Empresa>(), e => { })
      .subscribe(empresas => {
        this.clientsSelects = [];
        empresas.forEach(emp => this.clientsSelects.push({ value: emp, viewValue: emp.nome }));
      });
    this.buscarAtendimentosIdUsuario();
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
            this.buscarAtendimentosIdUsuario();
            this.onSucessMensage("Sucesso", "Chamado registrado com sucesso!");
          }
        });
  }

  selecionar(atn: Atendimento) {
    this.atendimento = atn;
  }

  limpar() {
    this.dataInicio = new Date();
    this.horaInicio = '00:00'
    this.dataFim = new Date();
    this.horaFim = '00:00';
    this.atendimento = new Atendimento();
    this.atendimento.horariosAtendimento = new Array<Intervalo>();
    this.atendimento.horariosAtendimento.push(new Intervalo());
  }

  statusName(s: StatusAtendimento): string {
    return StatusAtendimento[s];
  }

  onChangePageEvent(p: PageEvent) {
    this.pageEvent = p;
    if (this.textoPesquisa !== undefined && this.textoPesquisa !== '') {
       this.onChangePesquisar();
       return;
    }
    this.buscarAtendimentosIdUsuario();
  }

  textoPesquisa: string = '';
  onChangePesquisar() {
    this.atendPagHttpClient
      .get(DnsWebService.ATENDIMENTO_USUARIO
        + '/' + DnsWebService.usuario.id
        + '/' + this.getNumPag(this.pageEvent)
        + '/' + this.getPageSize(this.pageEvent)
        + '/' + this.textoPesquisa,
        false, new Pagina<Atendimento>(), e1 => {
          this.onErrorMensage(e1.codigo, e1.mensagem);
        })
      .subscribe(atendReturs => {
        this.atendimentos = [];
        atendReturs.conteudo.forEach((a) => this.atendimentos.push(a));
        this.dataSource.data = this.atendimentos;
        this.length = atendReturs.totalRegistro.valueOf();
      });
  }

 

  private buscarAtendimentosIdUsuario() {
    this.atendPagHttpClient
      .get(DnsWebService.ATENDIMENTO_USUARIO + '/' + DnsWebService.usuario.id
        + '/' + this.getNumPag(this.pageEvent)
        + '/' + this.getPageSize(this.pageEvent),
        false, new Pagina<Atendimento>(), e1 => {
          this.onErrorMensage(e1.codigo, e1.mensagem);
        })
      .subscribe(atendReturs => {
        this.atendimentos = [];
        atendReturs.conteudo.forEach((a) => this.atendimentos.push(a));
        this.dataSource.data = this.atendimentos;
        this.length = atendReturs.totalRegistro.valueOf();
      });
  }

  private getNumPag(event: PageEvent): number {
    if (event === undefined || event === null) {
      return 0;
    }
    return event.pageIndex;
  }

  private getPageSize(event: PageEvent): number {
    if (event === undefined || event === null) {
      return 10;
    }
    return event.pageSize;
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
