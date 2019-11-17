export class Rotina {
    public id : Number;
    public rotina : string;
}

export class Acesso {
    public rotina : Rotina;
    public criar : boolean;
    public atualizar : boolean;
    public deletar : boolean;
    public visualizar : boolean;
}

export class Usuario {
    public id : Number;
    public nome: string;
    public login : string;
    public senha : string;
    public acesso : Acesso[];
}