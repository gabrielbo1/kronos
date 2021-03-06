CREATE SEQUENCE empresa_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE EMPRESA (
    ID BIGINT DEFAULT nextval('empresa_id_seq') PRIMARY KEY,
    NOME_EMPRESA VARCHAR(150),
    ATIVA BOOLEAN
);

CREATE SEQUENCE rotina_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE ROTINA (
  ID BIGINT DEFAULT nextval('rotina_id_seq') PRIMARY KEY,
  ROTINA VARCHAR(150) NOT NULL
);

CREATE SEQUENCE usuario_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE USUARIO (
    ID BIGINT DEFAULT nextval('usuario_id_seq') PRIMARY KEY,
    NOME VARCHAR(150) NOT NULL,
    LOGIN VARCHAR(50) NOT NULL,
    SENHA VARCHAR(65) NOT NULL,
    ACESSO JSON NOT NULL
);

CREATE SEQUENCE ponto_id_seq 
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE PONTO (
   ID BIGINT DEFAULT nextval('ponto_id_seq') PRIMARY KEY,
   IDUSUARIO BIGINT NOT NULL,
   DATA TIMESTAMP NOT NULL,
   CONSTRAINT PONTO_USUARIO FOREIGN KEY (IDUSUARIO) REFERENCES USUARIO(ID) 
);

CREATE SEQUENCE atendimento_id_seq 
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE ATENDIMENTO (
    ID BIGINT DEFAULT nextval('atendimento_id_seq') PRIMARY KEY,
    IDUSUARIO BIGINT NOT NULL,
    IDCLIENTE BIGINT NOT NULL,
    HORARIOS_ATENDIMENTO JSON,
    STATUS_ATENDIMENTO SMALLINT NOT NULL,
    OBSERVACAO TEXT,
    CONSTRAINT ATENDIMENTO_USUARIO FOREIGN KEY (IDUSUARIO) REFERENCES USUARIO(ID),
    CONSTRAINT ATENDIMENTO_EMPRESA FOREIGN KEY (IDCLIENTE) REFERENCES EMPRESA(ID) 
);