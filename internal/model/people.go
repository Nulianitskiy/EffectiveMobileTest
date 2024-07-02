package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type People struct {
	Id             int    `db:"id" json:"id"`
	PassportSerie  int    `db:"passport_serie" json:"passport_serie"`
	PassportNumber int    `db:"passport_number" json:"passport_number"`
	Name           string `db:"name" json:"name"`
	Surname        string `db:"surname" json:"surname"`
	Patronymic     string `db:"patronymic" json:"patronymic,omitempty"`
	Address        string `db:"address" json:"address"`
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

	if len(d) != 2 {
		return fmt.Errorf("Неверный формат passportNumber: ожидалось 'serie number', got '%s'", aux.PassportNumber)
	}

	serie, err := strconv.Atoi(d[0])
	if err != nil {
		return fmt.Errorf("Неверный формат passportSerie: %s", err)
	}
	number, err := strconv.Atoi(d[1])
	if err != nil {
		return fmt.Errorf("Неверный формат passportNumber: %s", err)
	}

	p.PassportSerie = serie
	p.PassportNumber = number

	return nil
}
