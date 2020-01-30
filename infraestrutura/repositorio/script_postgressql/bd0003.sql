CREATE SEQUENCE intervalo_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;

CREATE TABLE INTERVALO (
 ID BIGINT DEFAULT nextval('intervalo_id_seq') PRIMARY KEY,
 IDATENDIMENTO BIGINT NOT NULL,
 DATA_INICIO TIMESTAMP NOT NULL,
 DATA_FIM TIMESTAMP,
 CONSTRAINT INTERVALO_ATENDIMENTO FOREIGN KEY (IDATENDIMENTO) REFERENCES ATENDIMENTO(ID)
)