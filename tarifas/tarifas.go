package tarifas

import (
	"errors"
	"fmt"
)

// TipoServicio representa la modalidad de envío.
type TipoServicio int

// Declaración con iota de las constantes de servicio
const (
	Estandar TipoServicio = iota
	Express
	Prioritario
)

// Declaración de constantes del negocio
const (
	TarifaBase = 8.00
	PrecioKg   = 2.50
	PrecioKm   = 0.15
	PesoMaximo = 50.0
	IGV        = 0.18
)

// CalcularSubtotal devuelve el costo sin IGV o un error si los datos no son válidos.
func CalcularSubtotal(pesoKg float64, distanciaKm int, servicio TipoServicio) (float64, error) {
	if pesoKg <= 0 || pesoKg > PesoMaximo {
		return 0, fmt.Errorf("peso no válido: %.2f kg (debe estar entre 0 y 50 kg)", pesoKg)
	}
	if distanciaKm <= 0 {
		return 0, errors.New("distancia no válida (debe ser mayor a 0 km)")
	}

	var recargo float64
	switch servicio {
	case Estandar:
		recargo = 1.0
	case Express:
		recargo = 1.3
	case Prioritario:
		recargo = 1.6
	default:
		return 0, errors.New("tipo de servicio no válido")
	}

	subtotal := (TarifaBase + (pesoKg * PrecioKg) + (float64(distanciaKm) * PrecioKm)) * recargo
	return subtotal, nil
}

// AplicarIGV devuelve el IGV y el total usando retornos con nombre.
func AplicarIGV(subtotal float64) (igv, total float64) {
	igv = subtotal * IGV
	total = subtotal + igv
	return
}