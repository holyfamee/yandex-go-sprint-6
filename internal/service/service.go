package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(s string) string {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			return morse.ToMorse(s)
		}
	}
	return morse.ToText(s)
}
