package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type People struct {
	Id             int    `db:"id"`
	PassportSerie  int    `db:"passport_serie"`
	PassportNumber int    `db:"passport_number"`
	Name           string `db:"name"`
	Surname        string `db:"surname"`
	Patronymic     string `db:"patronymic"`
	Address        string `db:"address"`
}

type passport struct {
	PassportNumber string `json:"passportNumber"`
}

func (p *People) UnmarshalJSON(data []byte) error {
	var aux passport
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	d := strings.Split(aux.PassportNumber, " ")

	if len(d) == 2 {
		//TODO доделать проверку
		p.PassportSerie, _ = strconv.Atoi(d[0])
		p.PassportNumber, _ = strconv.Atoi(d[1])
	} else {
		return fmt.Errorf("invalid passportNumber")
	}

	return nil
}
