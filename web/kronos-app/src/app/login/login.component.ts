import { Component, OnInit } from '@angular/core';
import { Usuario } from '../model/usuario';
import { NotificationsService } from 'angular2-notifications';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { HttpService } from '../model/httpclient';
import { DnsWebService } from '../model/dns';
import { Mensagem } from '../model/mensagem';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  public usuario : Usuario;

  constructor(
    private router: Router, 
    private httpClient: HttpClient,
    private notificationsService: NotificationsService) { 
    this.usuario = new Usuario();
  }

  ngOnInit() {
    DnsWebService.storageTokenUsuarioAdm = '';
  }

  login() {
    new HttpService<Usuario>(this.httpClient)
    .post(DnsWebService.LOGIN_USUARIO, this.usuario, false, this.usuario, e => {       
        const toast = this.notificationsService.error(e.codigo, e.mensagem, {
            timeOut: 3000,
            showProgressBar: true,
            pauseOnHover: true,
            clickToClose: true,
          });  
    })
    .subscribe(usuario => {
      let m : Mensagem = new Mensagem();
      DnsWebService.usuario = usuario;
      new HttpService<Mensagem>(this.httpClient)
        .get(DnsWebService.LOGIN_USUARIOOK, false, m, e => {
          const toast = this.notificationsService.error('LOGIN', 'Login ou senha incorretos', {
            timeOut: 3000,
            showProgressBar: true,
            pauseOnHover: true,
            clickToClose: true,
          });  
        })
        .subscribe(m => {
          if (m.mensagem !== undefined && m.mensagem === 'OK') {
            this.router.navigate(['kronos']);
          }
        })
    });
  }

}
