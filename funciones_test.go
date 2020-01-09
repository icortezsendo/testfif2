package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeparaDatos(t *testing.T) {
	lista := [9][2]string{{"11AB398765UJ1A050200N23", "23"},
		{"1nAB398765UJ1A050200N23", "Error en convertir el largo a numerico: 1n"},
		{"11AB398765UJ1C050200N23", "Tipo dato debe ser N o A, no C"},
		{"11AB398765UJ1N050200N23", "Tipo de dato definido como Numerico y es Alfanumerico"},
		{"11AB398765UJ1A050n00N23", "Error en convertir el largo a numerico: 0n"},
		{"11AB398765UJ1A050200C23", "Tipo dato debe ser N o A, no C"},
		{"11AB398765UJ1A05020cN23", "Tipo de dato definido como Numerico y es Alfanumerico"},
		{"", "No se introdujo el dato a leer"},
		{"11AB398765UJ1A050200N23016N66", "23"},
	}
	for _, dato := range lista {
		mapa, err := separaDatos([]byte(dato[0]))
		if err == "" {
			assert.Contains(t, mapa, dato[1])
		} else {
			assert.Equal(t, err, dato[1])
		}

	}
}
