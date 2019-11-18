import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError, tap } from 'rxjs/operators';
import { Erro } from './erro';
import { DnsWebService } from './dns';
import { Usuario } from './usuario';

export type TratamentoErro = (erro: Erro) => void;

export class HttpService<T> {
    constructor(private http: HttpClient) {
    }

    public post(servico: string, body: any, aplicacao: boolean, retornoPadrao: T, erroFunc: TratamentoErro): Observable<T> {
        return this.http.post<T>(DnsWebService.dns + servico, body, { headers: this.getHeaders(aplicacao) })
            .pipe(
                tap(o => this.log('LOG - ' + DnsWebService.dns + servico)),
                catchError(this.handleError(servico + ' - POST', erroFunc, retornoPadrao)));
    }

    public put(servico: string, body: any, aplicacao: boolean, retornoPadrao: T, erroFunc: TratamentoErro): Observable<T> {
        return this.http.put<T>(DnsWebService.dns + servico, body, { headers: this.getHeaders(aplicacao) })
            .pipe(
                tap(o => this.log('LOG - ' + DnsWebService.dns + servico)),
                catchError(this.handleError(servico + ' - PUT', erroFunc, retornoPadrao)));
    }

    public delete(servico: string, b: any, aplicacao: boolean, retornoPadrao: T, erroFunc: TratamentoErro): Observable<T> {
        return this.http.request<T>('delete', DnsWebService.dns + servico, {
            headers: this.getHeaders(aplicacao),
            body: b
        })
            .pipe(
                tap(o => this.log('LOG - ' + DnsWebService.dns + servico)),
                catchError(this.handleError(servico + ' - DELETE', erroFunc, retornoPadrao)));
    }

    public get(servico: string, aplicacao: boolean, retornoPadrao: T, erroFunc: TratamentoErro): Observable<T> {
        return this.http.get<T>(DnsWebService.dns + servico, { headers: this.getHeaders(aplicacao) })
            .pipe(
                tap(o => this.log('LOG - ' + DnsWebService.dns + servico)),
                catchError(this.handleError(servico + ' - GET', erroFunc, retornoPadrao)));
    }

    public login(servico: string, body: Usuario, retornoPadrao: Usuario, erroFunc: TratamentoErro): Observable<Usuario> {
        let h: HttpHeaders = new HttpHeaders();
        h = h.append('Content-Type', 'application/json');
        h = h.append('Authorization', 'Basic ' + btoa(body.login + ':' + body.senha));
        
    
        return this.http.post<Usuario>(DnsWebService.dns + servico, body, { headers: h  })
          .pipe(
            tap(o => this.log('LOG - ' + DnsWebService.dns + servico)),
            catchError(this.handleError(servico + ' - POST', erroFunc, retornoPadrao)));
      }
    

    public getHeaders(aplicacao: boolean): HttpHeaders {
        let token: string = '';
        token += aplicacao ? DnsWebService.storageTokenAplicacao : DnsWebService.storageTokenUsuarioAdm;
        let h: HttpHeaders = new HttpHeaders();
        h = h.append('Content-Type', 'application/json');
        h = h.append('Authorization', 'Basic ' + token);
        return h;
    }
    /**
     * Handle Http operation that failed.
     * Let the app continue.
     * @param operation - name of the operation that failed
     * @param result - optional value to return as the observable result
     */
    // tslint:disable-next-line
    private handleError<T>(operation = 'operation', erroFunc: TratamentoErro, result?: T) {
        return (error: HttpErrorResponse): Observable<T> => {
            if (error.status !== 400 && error.status !== 401 && error.status !== 422) {
                // TODO: send the error to remote logging infrastructure
                console.error(error); // log to console instead

                // TODO: better job of transforming error for user consumption
                this.log(`${operation} failed: ${error.message}`);

            }

            if (error.status === 401 || error.status == 403) {
                const erroAutorizacao : Erro = new Erro();
                erroAutorizacao.codigo = 'ERRO';
                erroAutorizacao.mensagem = 'Erro requisição não autorizada.';
                erroFunc(erroAutorizacao);
            }

            // Tratamento de exceções de négocio.
            if (error.status === 400 || error.status === 422) {
                const erroNegocio: Erro = new Erro();

                if (error.error.codigo !== undefined) {
                    erroNegocio.codigo = error.error.codigo;
                }

                if (error.error.mensagem !== undefined) {
                    erroNegocio.mensagem = error.error.mensagem;
                }
                erroFunc(erroNegocio);
            }

            // Let the app keep running by returning an empty result.
            return of(result as T);
        };
    }

    /** Log a HeroService message with the MessageService */
    private log(message: string) {
        // tslint:disable-next-line console.log(message);
    }
}