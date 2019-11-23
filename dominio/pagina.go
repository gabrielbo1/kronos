package dominio

type Pagina struct {
	NumPagina     int         `json:"numPagina" example:"1"`
	QtdPorPagina  int         `json:"qtdPagina" example:"10"`
	TotalRegistro int         `json:"totalRegistro" example:"1000"`
	TotalPagina   int         `json:"totalPagina" example:"100"`
	Conteudo      interface{} `json:"conteudo"`
}

func CalcQtdPaginas(totalRegistro int, qtdPorPagina int) (totalPagina int) {
	totalPagina = (totalRegistro + (qtdPorPagina - (totalRegistro % qtdPorPagina))) / qtdPorPagina
	return totalPagina
}
