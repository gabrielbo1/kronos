import { Component, OnInit, AfterViewInit, ViewChild, Inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormControl } from '@angular/forms';

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

import { Mensagem } from 'src/app/model/mensagem';

export interface StatusSelect {
  value: StatusAtendimento;
  viewValue: string;
}

export interface ClienteSelect {
  value: Empresa,
  viewValue: string;
  id: Number;
}

export interface DeletarAtendimento {
  value: boolean;
  atendimentoComponent: AtendimentoComponent;
}

@Component({
  selector: 'app-atendimento',
  templateUrl: './atendimento.component.html',
  styleUrls: ['./atendimento.component.scss']
})
export class AtendimentoComponent implements OnInit, AfterViewInit {

  @ViewChild('dataInicioValor', { static: false })
  dataInicioValor: any;
  dataInicioValorForm = new FormControl(new Date());
  public dataInicio: Date = new Date();
  public horaInicio: string = '00:00'

  @ViewChild('dataFimValor', { static: false })
  dataFimValor: any;
  dataFimValorForm = new FormControl(new Date());
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
  public deletarAtendimento: boolean = false;
  public addIntervalo: boolean = false;

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
  private atendDeltHttpClient: HttpService<Mensagem>;

  constructor(
    private httpClient: HttpClient,
    private notificationsService: NotificationsService,
    public dialog: MatDialog) {
    moment.locale('pt-BR');
  }

  ngOnInit() { }

  onChangeDataInicio(s: any) {
    this.dataInicio = moment(s.value, 'DD/MM/YYYY').toDate();
  }

  onChangeDataFim(s: any) {
    this.dataFim = moment(s.value, 'DD/MM/YYYY').toDate();
  }

  ngAfterViewInit() {
    this.atendPagHttpClient = new HttpService<Pagina<Atendimento>>(this.httpClient);
    this.atendHttpClient = new HttpService<Atendimento>(this.httpClient);
    this.atendDeltHttpClient = new HttpService<Mensagem>(this.httpClient);

    new HttpService<Array<Empresa>>(this.httpClient)
      .get(DnsWebService.EMPRESA, false, new Array<Empresa>(), e => { })
      .subscribe(empresas => {
        this.clientsSelects = [];
        empresas.forEach(emp => this.clientsSelects.push({ id: emp.id, value: emp, viewValue: emp.nome }));
      });
    this.buscarAtendimentosIdUsuario();
  }


  cadastro() {
    if (this.validaAtendimento(this.atendimento)) {
      this.atendHttpClient
        .post(DnsWebService.ATENDIMENTO, this.atendimento, false, new Atendimento(), e => {
          this.onErrorMensage(e.codigo, e.mensagem);
        })
        .subscribe(atndReturn => {
          if (atndReturn.id !== undefined && atndReturn.id !== 0) {
            this.limpar();
            this.onSucessMensage("Sucesso", "Chamado registrado com sucesso!");
          }
        });
    }
  }

  atualizar() {
    if (this.validaAtendimento(this.atendimento)) {
      this.atendHttpClient
        .put(DnsWebService.ATENDIMENTO, this.atendimento, false, new Atendimento(), e => {
          this.onErrorMensage(e.codigo, e.mensagem);
        })
        .subscribe(atndReturn => {
          if (atndReturn.id !== undefined && atndReturn.id !== 0) {
            this.limpar();
            this.onSucessMensage("Sucesso", "Chamado atualizado com sucesso!");
          }
        });
    }
  }

  deletar() {
    this.deletarAtendimento = false;
    const dialogRef = this.dialog.open(AtendimentoDeletetarDialogComponent, {
      width: '250px',
      data: { value: false, atendimentoComponent: this }
    });

    dialogRef.afterClosed().subscribe(result => {

      if (this.deletarAtendimento) {
        let msg: Mensagem = new Mensagem();
        this.atendDeltHttpClient
          .delete(DnsWebService.ATENDIMENTO + '/' + this.atendimento.id, msg, false, new Mensagem(), e => {
            this.onErrorMensage(e.codigo, e.mensagem);
          })
          .subscribe(msgReturn => {
            if (msgReturn !== undefined &&
              msgReturn.mensagem !== undefined &&
              msgReturn.mensagem !== '') {
              this.onSucessMensage('Sucesso', msgReturn.mensagem);
              this.limpar();
            }

          });
      }
    });
  }

  compareObjects(o1: any, o2: any): boolean {
    if (o1 === undefined || o1.id === undefined) return false;
    if (o2 === undefined || o2.id === undefined) return false;
    return o1.id === o2.id;
  }

  selecionar(atn: Atendimento) {
    this.limpar();
    this.statusAtendimentoSelecionado = atn.statusAtendimento;
    this.atendimento = new Atendimento(atn.id,
      atn.usuario,
      atn.cliente,
      atn.horariosAtendimento,
      atn.statusAtendimento,
      atn.observacao);

    let length: number = this.atendimento.horariosAtendimento === undefined ? 0 :
      this.atendimento.horariosAtendimento.length;
    if (length > 0) {

      this.dataInicioValorForm.setValue(moment(this.atendimento.horariosAtendimento[length - 1].dataInicio, 'YYYY-MM-DDTHH:mm:ssZ'));
      this.horaInicio = this.atendimento.horariosAtendimento[length - 1].dataInicio.substr(11, 5);

      if (this.atendimento.horariosAtendimento[length - 1].dataFim !== '') {
        this.dataFimValorForm.setValue(moment(this.atendimento.horariosAtendimento[length - 1].dataFim, ''));
        this.horaFim = this.atendimento.horariosAtendimento[length - 1].dataFim.substr(11, 5);
      }
    }
  }

  limpar() {
    this.dataInicio = new Date();
    this.horaInicio = '00:00'
    this.dataFim = new Date();
    this.horaFim = '00:00';
    this.atendimento = new Atendimento();
    this.addIntervalo = false;
    this.dataInicioValorForm.setValue(undefined);
    this.dataFimValorForm.setValue(undefined);
    this.statusAtendimentoSelecionado = StatusAtendimento.Aberto;
    this.addIntervalo = false;
    this.buscarAtendimentosIdUsuario();
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

  disableSalvar(): boolean {
    return this.atendimento !== undefined &&
      this.atendimento !== null &&
      this.atendimento.id !== undefined &&
      this.atendimento.id !== null &&
      this.atendimento.id !== 0;
  }

  disableAtualizar(): boolean {
    return this.atendimento !== undefined &&
      this.atendimento !== null &&
      this.atendimento.id !== undefined &&
      this.atendimento.id !== null &&
      this.disableDeletar();
  }

  statusAtendimentoSelecionado: StatusAtendimento;
  disableDeletar(): boolean {
    return this.statusAtendimentoSelecionado !== undefined && this.statusAtendimentoSelecionado === StatusAtendimento.Fechado;
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

  private validaAtendimento(atnd: Atendimento): boolean {
    let length: number = this.atendimento.horariosAtendimento === undefined ? 0 :
      this.atendimento.horariosAtendimento.length;


    if (this.atendimento.horariosAtendimento === undefined ||
      this.atendimento.horariosAtendimento.length === 0) {
      this.atendimento.horariosAtendimento = new Array<Intervalo>();
      this.atendimento.horariosAtendimento.push(new Intervalo('', ''));
      length = atnd.horariosAtendimento.length;
    } else if (this.atendimento.horariosAtendimento[length - 1].dataInicio !== undefined &&
      this.atendimento.horariosAtendimento[length - 1].dataInicio !== null &&
      this.atendimento.horariosAtendimento[length - 1].dataInicio !== '' &&
      this.atendimento.horariosAtendimento[length - 1].dataFim !== undefined &&
      this.atendimento.horariosAtendimento[length - 1].dataFim !== null &&
      this.atendimento.horariosAtendimento[length - 1].dataFim !== '' &&
      !this.addIntervalo) {
      this.atendimento.horariosAtendimento.push(new Intervalo());
      length = atnd.horariosAtendimento.length;
    }
    this.addIntervalo = true;

    if (this.horaInicio === undefined || this.horaInicio === '00:00') {
      this.onErrorMensage('Erro', 'Informe a hora de início do atendimento.')
      return false;
    } else {
      atnd.horariosAtendimento[length - 1].dataInicio
        = moment(this.dataInicio).format('YYYY-MM-DD') + ' ' + this.horaInicio;
    }

    if (this.horaFim === undefined || this.horaFim === '00:00') {
      atnd.horariosAtendimento[length - 1].dataFim = '';
    }

    if (this.dataFim !== undefined && this.horaFim !== undefined && this.horaFim !== '00:00') {
      atnd.horariosAtendimento[length - 1].dataFim =
        moment(this.dataFim).format('YYYY-MM-DD') + ' ' + this.horaFim;
    }

    if (atnd.cliente === undefined) {
      this.onErrorMensage('Erro', 'Informe o cliente do atendimento.')
      return false;
    }

    if (atnd.statusAtendimento === undefined) {
      this.onErrorMensage('Erro', 'Informe o status do atendimento.')
      return false;
    }
    
    if (this.atendimento.horariosAtendimento.length >= 2) {
      if (this.atendimento.horariosAtendimento[0].dataInicio.replace('T','').replace(':00Z','')
           === this.atendimento.horariosAtendimento[1].dataInicio.replace(' ', '') &&
        this.atendimento.horariosAtendimento[0].dataFim.replace('T','').replace(':00Z','')
         === this.atendimento.horariosAtendimento[1].dataFim.replace(' ', '')) {
        this.atendimento.horariosAtendimento.pop();
      }
    }

    atnd.usuario = DnsWebService.usuario;
    return true;
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


@Component({
  selector: 'deletar-atendimento-dialog',
  template: `
  <h1 mat-dialog-title>Deletar atendimento número {{data.atendimentoComponent.atendimento.id}}</h1>
  <div mat-dialog-actions>
    <button mat-button (click)="nao()">Não</button>
    <button mat-button (click)="sim()">Sim</button>
  </div>
  `
})
export class AtendimentoDeletetarDialogComponent {
  constructor(public dialogRef: MatDialogRef<AtendimentoDeletetarDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DeletarAtendimento) {
  }

  nao(): void {
    this.data.atendimentoComponent.deletarAtendimento = false;
    this.dialogRef.close();
  }

  sim(): void {
    this.data.atendimentoComponent.deletarAtendimento = true;
    this.dialogRef.close();
  }
}