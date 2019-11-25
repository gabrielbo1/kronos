import { Usuario } from './usuario';
import { Empresa } from './empresa';

export enum StatusAtendimento {
    Aberto = 1,
    Espera,
    Fechado
}

export class Intervalo {
    public dataInicio : string = '';
    public dataFim : string = '';
    constructor(dtInic ?: string, dtFim ?: string) {
        this.dataInicio = dtInic;
        this.dataFim = dtFim;
    }
}

export class Atendimento {
    public id : Number;
    public usuario : Usuario;
    public cliente : Empresa;
    public horariosAtendimento: Array<Intervalo>;
    public statusAtendimento : StatusAtendimento;
    public observacao : string;

    constructor(id ?: Number,
                usuario ?: Usuario,
                cliente ?: Empresa, 
                horariosAtendimento ?: Array<Intervalo>,
                status ?: StatusAtendimento,
                obs ?: string) {
        this.id = id;
        this.usuario = usuario;
        this.cliente = cliente;
        this.horariosAtendimento = horariosAtendimento;
        this.statusAtendimento = status;
        this.observacao = obs;
    }
}