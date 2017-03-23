package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func getNcms(c echo.Context) (err error) {
	if ncms, count := dbGetNcms(); count > 0 {
		n := NCMS{ncms}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}

func getNcmsByUser(c echo.Context) (err error) {
	userID := c.Param("id")

	if ncms, count := dbGetNcmsByUser(userID); count > 0 {
		n := NCMS{ncms}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}

func getLeadsTrial(c echo.Context) (err error) {
	if leads, count := dbGetLeadsTrial(); count > 0 {
		l := Leads{leads}
		err = c.JSON(http.StatusOK, l)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}
