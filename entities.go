package main

import "time"

// NCM type
type NCM struct {
	Ncm            string    `json:"ncm"`
	Descricao      string    `json:"descricao"`
	VigenciaIncial time.Time `json:"vigenciaInicial"`
	VigenciaFinal  time.Time `json:"vigenciaFinal"`
	Tributo        string    `json:"tributo"`
	Instituto      string    `json:"instituto"`
}

// NCMS is a slice of NCM
type NCMS struct {
	NCMS []NCM `json:"ncms"`
}

// Lead type
type Lead struct {
	Email          string    `json:"email"`
	Cpf            string    `json:"cpf"`
	Cnpj           string    `json:"cnpj"`
	Ncm            string    `json:"ncm"`
	Descricao      string    `json:"descricao"`
	VigenciaIncial time.Time `json:"vigenciaInicial"`
	VigenciaFinal  time.Time `json:"vigenciaFinal"`
	Tributo        string    `json:"tributo"`
	Instituto      string    `json:"instituto"`
}

// Leads is a slice of lead
type Leads struct {
	Leads []Lead `json:"leads"`
}
