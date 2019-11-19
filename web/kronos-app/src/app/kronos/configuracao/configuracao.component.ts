import { Component, OnInit, AfterViewInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { SelectionModel } from '@angular/cdk/collections';
import { NotificationsService } from 'angular2-notifications';
import { HttpClient } from '@angular/common/http';


import { HttpService } from '../../model/httpclient';
import { DnsWebService } from '../../model/dns';
import { Empresa } from '../../model/empresa';
import { Mensagem } from 'src/app/model/mensagem';


@Component({
  selector: 'app-configuracao',
  templateUrl: './configuracao.component.html',
  styleUrls: ['./configuracao.component.scss']
})
export class ConfiguracaoComponent implements OnInit, AfterViewInit {
  constructor(
    private httpClient: HttpClient,
    private notificationsService: NotificationsService) { }


  public empresa: Empresa = new Empresa();
  public empresas: Array<Empresa> = new Array<Empresa>();
  displayedColumns: string[] = ['ID', 'Nome', 'Ativa'];
  dataSource = new MatTableDataSource(this.empresas);
  selection = new SelectionModel<Empresa>(true, []);

  ngOnInit() { }

  ngAfterViewInit() {
    this.carregarEmpresas();
  }

  selectionar(emp: Empresa) {
    this.empresa.id = emp.id;
    this.empresa.nome = emp.nome;
    this.empresa.ativa = emp.ativa;
  }

  limpar() {
    this.empresa = new Empresa();
  }

  cadastro(emp: Empresa) {
    new HttpService<Empresa>(this.httpClient)
      .post(DnsWebService.EMPRESA, this.empresa, false, this.empresa, e => {
        const toast = this.notificationsService.error(e.codigo, e.mensagem, {
          timeOut: 3000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: true,
        });
      })
      .subscribe(emp => {
        if (emp !== undefined && emp.id !== 0) {
          this.empresa = new Empresa();
          this.carregarEmpresas();
        }
      });
  }

  atualizar(emp: Empresa) {
    new HttpService<Empresa>(this.httpClient)
      .put(DnsWebService.EMPRESA, this.empresa, false, this.empresa, e => {
        const toast = this.notificationsService.error(e.codigo, e.mensagem, {
          timeOut: 3000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: true,
        });
      })
      .subscribe(emp => {
        if (emp !== undefined && emp.id !== 0) {
          this.empresa = new Empresa();
          this.carregarEmpresas();
        }
      });
  }

  delete(emp : Empresa) {
    new HttpService<Mensagem>(this.httpClient)
      .delete(DnsWebService.EMPRESA + '/' + this.empresa.id, this.empresa, false, new Mensagem(), e => {
        const toast = this.notificationsService.error(e.codigo, e.mensagem, {
          timeOut: 3000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: true,
        });
      })
      .subscribe(msg => {
        const toast = this.notificationsService.success('Sucesso', msg.mensagem, {
          timeOut: 3000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: true,
        });
         this.carregarEmpresas();
      });
  }

  public carregarEmpresas() : void  {
    new HttpService<Array<Empresa>>(this.httpClient)
    .get(DnsWebService.EMPRESA, false, this.empresas, err => {
      const toast = this.notificationsService.error(err.codigo, err.mensagem, {
        timeOut: 3000,
        showProgressBar: true,
        pauseOnHover: true,
        clickToClose: true,
      });
    })
    .subscribe(emp => {
      this.empresas = new Array<Empresa>();
      emp.forEach((e) => this.empresas.push(e));
      this.empresas.sort((a, b) => a.id.valueOf() - b.id.valueOf());
      this.dataSource.data = this.empresas;
      this.dataSource.data;
    });
  }
}
