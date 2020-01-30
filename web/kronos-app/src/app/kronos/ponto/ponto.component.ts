import { Component, OnInit, AfterViewInit, OnDestroy, Inject } from '@angular/core';

import { Ponto } from '../../model/ponto';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';
import { NotificationsService } from 'angular2-notifications';
import { HttpService } from 'src/app/model/httpclient';
import { Usuario } from 'src/app/model/usuario';
import { DnsWebService } from 'src/app/model/dns';

import * as moment from 'moment';


function atualizaDataAtual(comp: PontoComponent) {
  setTimeout(function () {
    const fevereiro: number = 2;
    const diaFimHorarioVerao2020 = 16;
    const anoImplementacao: number = 2020;
    let hoje: Date = new Date();
    comp.data = hoje;

    //Ajuste horario de verao.
    if (hoje.getFullYear() < anoImplementacao ||
      (hoje.getMonth() <= fevereiro && hoje.getDay() <= diaFimHorarioVerao2020)) {
      comp.data = new Date(hoje.valueOf() - hoje.getTimezoneOffset() * 30000);
    }
    atualizaDataAtual(comp);
  }, 1000);
}

@Component({
  selector: 'app-ponto',
  templateUrl: './ponto.component.html',
  styleUrls: ['./ponto.component.scss']
})
export class PontoComponent implements OnInit, AfterViewInit {
  public msgError: string = '';
  public data: Date = new Date();
  public dataEnvio: Date;
  public onViewOk: boolean = false;
  public pontos: Array<Ponto> = new Array<Ponto>();
  public pontoOk: boolean = false;
  private httpPonto: HttpService<Ponto>;
  private httpPontos: HttpService<Array<Ponto>>;
  displayedColumns: string[] = ['data'];

  constructor(
    private httpClient: HttpClient,
    private notificationsService: NotificationsService,
    public dialog: MatDialog) { }

  ngAfterViewInit() {
    this.onViewOk = true;
    atualizaDataAtual(this);
    this.buscarPontos();
  }

  ngOnInit() {
    this.httpPonto = new HttpService<Ponto>(this.httpClient);
    this.httpPontos = new HttpService<Array<Ponto>>(this.httpClient);
  }

  ngOnDestroy() {
    this.onViewOk = false;
  }

  registrarPonto() {
    this.dataEnvio = this.data;
    this.pontoOk = false;
    const dialogRef = this.dialog.open(ConfirmarPontoComponent, {
      width: '250px',
      data: { pontoComponent: this }
    });

    dialogRef.afterClosed()
             .subscribe(result => {
                console.log(' result ' + this.pontoOk);
                if (this.pontoOk) {
                  let ponto: Ponto = new Ponto();
                  ponto.usuario = new Usuario();
                  ponto.usuario = DnsWebService.usuario;
                  ponto.data = moment(this.dataEnvio).format('YYYY-MM-DD')  + ' ' +
                              this.dataEnvio.toLocaleTimeString();


                  this.httpPonto
                      .post(DnsWebService.PONTO, ponto, false, new Ponto(), e => {
                        this.onErrorMensage(e.codigo, e.mensagem);
                      })
                      .subscribe(atndReturn => {
                        if (atndReturn.id !== undefined && atndReturn.id !== 0) {
                          this.onSucessMensage("Sucesso", "Ponto registrado com sucesso!");
                        }
                        this.limpar();
                        this.buscarPontos();  
                      });
                }
              });
  }

  limpar() {
    this.pontoOk = false;
    this.dataEnvio = null;
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

  private buscarPontos(): void {
    let hoje : string = moment(this.data).format('YYYY-MM-DD')  + ' ' +
                        this.data.toLocaleTimeString();
    this.httpPontos
        .get(DnsWebService.PONTO + '/' + DnsWebService.usuario.id.toString() + '/' + hoje,  false, new Array<Ponto>(), (err) => {
          this.onErrorMensage(err.codigo, err.mensagem);
        })
        .subscribe((pontos) => {
          this.pontos = [];
          pontos.forEach((p) => {
              this.pontos.push(p);
          });
        });
  }
}

export interface ConfirmarPonto {
  pontoComponent: PontoComponent;
}

@Component({
  selector: 'confirmar-ponto-dialog',
  template: `
  <h1 mat-dialog-title>Confirmar ponto</h1>
  <mat-dialog-content>Horário:</mat-dialog-content>
  <mat-dialog-content>{{data.pontoComponent.dataEnvio.toLocaleString()}}</mat-dialog-content>
  <div mat-dialog-actions>
    <button mat-button (click)="nao()">Não</button>
    <button mat-button (click)="sim()">Sim</button>
  </div>
  `
})
export class ConfirmarPontoComponent {
  constructor(public dialogRef: MatDialogRef<PontoComponent>,
    @Inject(MAT_DIALOG_DATA) public data: ConfirmarPonto) {
  }

  sim() {
    this.data.pontoComponent.pontoOk = true;
    this.dialogRef.close();
  }

  nao() {
    this.data.pontoComponent.pontoOk = false;
    this.dialogRef.close();
  }
}
