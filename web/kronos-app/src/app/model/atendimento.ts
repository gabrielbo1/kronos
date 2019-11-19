import { Usuario } from './usuario';
import { Empresa } from './empresa';

export enum StatusAtendimento {
    Aberto = 1,
    Espera,
    Fechado
}

export class Intervalo {
    public dataInicio : string = '';
    public dataFim : string = ''
}

export class Atendimento {
    public id : Number;
    public usuario : Usuario;
    public cliente : Empresa;
    public horariosAtendimento: Array<Intervalo> = new Array<Intervalo>();
    public statusAtendimento : StatusAtendimento;
    public observacao : string;
}