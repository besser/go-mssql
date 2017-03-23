# Rest API Service #

Version: 1.3

## __NCMs__ ##

API: http://localhost:9000/api/ncms

#### Return (JSON) ####
```json
{
  "ncms": [
    {
      "ncm"            : int,
      "descricao"      : string,
      "vigenciaInicial": datetime,
      "vigenciaFinal"  : datetime,
      "tributo"        : string,
      "instituto"      : string
    }
  ]
}
```

## __NCMs by user__ ##

API: http://locahost:9000/api/ncms/{userID}

#### Return (JSON) ####
```json
{
  "ncms": [
    {
      "ncm"            : int,
      "descricao"      : string,
      "vigenciaInicial": datetime,
      "vigenciaFinal"  : datetime,
      "tributo"        : string,
      "instituto"      : string
    }
  ]
}
```

## __Trial Leads__ ##

API: http://locahost:9000/api/leads/trial

#### Return (JSON) ####
```json
{
  "leads": [
    {
      "email"          : string,
      "cpf"            : string,
      "cnpj"           : string,
      "ncm"            : int,
      "descricao"      : string,
      "vigenciaInicial": datetime,
      "vigenciaFinal"  : datetime,
      "tributo"        : string,
      "instituto"      : string
    }
  ]
}  
```
