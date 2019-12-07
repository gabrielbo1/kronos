import { Usuario } from './usuario';

export class DnsWebService {

    // Endereco de onde se encontra o servidor na Web - https://www.zaynapp.tk
    public static  dns = window.location.hostname === 'localhost' ? 'http://localhost:8080' : 'https://' + window.location.hostname;

    // Constante para identificar tokens de segura da aplicacao
    // E dos usuarios.
    public static storageTokenUsuarioAdm  = '';

    // Constante para identifica tokens de seguranca da aplicacao.
    public static storageTokenAplicacao = '';

    public static usuario : Usuario = null;

    public static LOGIN_USUARIO : string = '/usuario/login';
    
    public static LOGIN_USUARIOOK : string = '/usuario/loginok';

    public static EMPRESA : string = '/empresa';

    public static ATENDIMENTO : string = '/atendimento';

    public static ATENDIMENTO_USUARIO : string = '/atendimento/usuario';

    public static PONTO : string = '/ponto';
}