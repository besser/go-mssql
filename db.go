package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

// Database variables
var connString string

func init() {
	connString = "server=172.30.0.237;user id=luke;password=tiprodutos;port=1433;database=datathon2017"
}

func dbGetConn() *sql.DB {
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	return db
}

func dbGetNcms() (ncms []NCM, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT [NCM_PESQUISADA],[DESCRICAO_NCM]," +
		"ISNULL([VIGENCIA_INICIAL],CONVERT(DATETIME, '1900-01-01', 102)) AS VIGENCIA_INICIAL," +
		"ISNULL([VIGENCIA_FINAL],CONVERT(DATETIME, '1900-01-01', 102)) AS VIGENCIA_FINAL," +
		"[TRIBUTO],[INSTITUTO] FROM ST_PESQUISAS_REGRAS WHERE DT_ENVIADO IS NULL AND TRIAL = 0 ORDER BY NCM_PESQUISADA")

	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var ncm, ncmDesc, tributo, instituto string
		var dtVigIni, dtVigFim time.Time

		err = rows.Scan(&ncm, &ncmDesc, &dtVigIni, &dtVigFim, &tributo, &instituto)
		if err != nil {
			log.Printf(err.Error())
		}

		ncms = append(ncms, NCM{
			Ncm:            ncm,
			Descricao:      ncmDesc,
			VigenciaIncial: dtVigIni,
			VigenciaFinal:  dtVigFim,
			Tributo:        tributo,
			Instituto:      instituto,
		})

		count++
	}

	return ncms, count
}

func dbGetNcmsByUser(userID string) (ncms []NCM, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT [NCM_PESQUISADA],[DESCRICAO_NCM],"+
		"ISNULL([VIGENCIA_INICIAL],CONVERT(DATETIME, '1900-01-01', 102)) AS VIGENCIA_INICIAL,"+
		"ISNULL([VIGENCIA_FINAL],CONVERT(DATETIME, '1900-01-01', 102)) AS VIGENCIA_FINAL,"+
		"[TRIBUTO],[INSTITUTO] FROM ST_PESQUISAS_REGRAS WHERE DT_ENVIADO IS NULL AND LOGIN_USUARIO = ?1 ORDER BY NCM_PESQUISADA", userID)

	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var ncm, ncmDesc, tributo, instituto string
		var dtVigIni, dtVigFim time.Time

		err = rows.Scan(&ncm, &ncmDesc, &dtVigIni, &dtVigFim, &tributo, &instituto)
		if err != nil {
			log.Printf(err.Error())
		}

		ncms = append(ncms, NCM{
			Ncm:            ncm,
			Descricao:      ncmDesc,
			VigenciaIncial: dtVigIni,
			VigenciaFinal:  dtVigFim,
			Tributo:        tributo,
			Instituto:      instituto,
		})

		count++
	}

	if *debug == false {
		dbSetAlertsSendedToUser(db, userID)
	}

	return ncms, count
}

func dbSetAlertsSendedToUser(db *sql.DB, userID string) (affected int64) {
	query := "UPDATE ST_PESQUISAS_REGRAS SET DT_ENVIADO = GETDATE() WHERE LOGIN_USUARIO = ?1 AND DT_ENVIADO IS NULL;"

	if result, err := db.Exec(query, userID); err != nil {
		log.Printf("Error: ", err.Error())
	} else {
		affected, _ = result.RowsAffected()
	}

	return
}

func dbGetLeadsTrial() (leadsTrial []Lead, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT [EMAIL],ISNULL([CPF],''),ISNULL([CNPJ],''),[NCM_PESQUISADA],[DESCRICAO_NCM]," +
		"ISNULL([VIGENCIA_INICIAL],CONVERT(DATETIME, '1900-01-01', 102)) AS VIGENCIA_INICIAL," +
		"ISNULL([VIGENCIA_FINAL],CONVERT(DATETIME, '1900-01-01', 102)) AS VIGENCIA_FINAL," +
		"[TRIBUTO],[INSTITUTO] FROM ST_PESQUISAS_REGRAS WHERE DT_ENVIADO IS NULL AND TRIAL = 1 ORDER BY NCM_PESQUISADA")

	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var email, cpf, cnpj, ncm, ncmDesc, tributo, instituto string
		var dtVigIni, dtVigFim time.Time

		err = rows.Scan(&email, &cpf, &cnpj, &ncm, &ncmDesc, &dtVigIni, &dtVigFim, &tributo, &instituto)
		if err != nil {
			log.Printf(err.Error())
		}

		leadsTrial = append(leadsTrial, Lead{
			Email:          email,
			Cpf:            cpf,
			Cnpj:           cnpj,
			Ncm:            ncm,
			Descricao:      ncmDesc,
			VigenciaIncial: dtVigIni,
			VigenciaFinal:  dtVigFim,
			Tributo:        tributo,
			Instituto:      instituto,
		})

		count++
	}

	return leadsTrial, count
}
