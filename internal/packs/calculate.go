package packs

import (
	"errors"
)

type Config struct {
	Packs []int `yaml:"packs"`
}

var errInputMustBeGreaterThanZero = errors.New("input must be greater than 0")
var errPacksEmpty = errors.New("packs are not defined")

type Calculator struct {
	cfg Config
}

func NewPacksCalculator(cfg Config) *Calculator {
	return &Calculator{
		cfg: cfg,
	}

}

func (c *Calculator) Calculate(input int) (map[int]int, int, error) {

	if len(c.cfg.Packs) == 0 {
		return nil, 0, errPacksEmpty
	}

	if input <= 0 {
		return nil, 0, errInputMustBeGreaterThanZero
	}

	list := c.cfg.Packs

	m := make(map[int]int)

	n := input / list[0]

	if input%list[0] != 0 {
		n++
	}

	m[list[0]] = n

	for i := len(list) - 1; i >= 0; i-- {

		d := list[i] / list[0]
		if d%2 != 0 {
			d--
		}

		for d != 0 && d <= n {
			m[list[i]] += 1
			n -= d
			m[list[0]] = n

			if m[list[0]] == 0 {
				delete(m, list[0])
			}

		}

	}

	total := 0

	for _, v := range m {
		total += v
	}

	return m, total, nil
}
