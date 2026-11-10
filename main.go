package main

import (
	"fmt"

	"encomiendas/tarifas"
)

func registrarEnvio(codigo string, peso float64, km int, servicio tarifas.TipoServicio) {
	fmt.Printf("\n[Inicio] Envío %s\n", codigo)
	
	// TODO 6: usa defer para imprimir el mensaje al salir de la función
	defer fmt.Printf("[Fin] Envío %s procesado\n", codigo)

	subtotal, err := tarifas.CalcularSubtotal(peso, km, servicio)
	if err != nil {
		fmt.Println("  Error:", err)
		return
	}
	igv, total := tarifas.AplicarIGV(subtotal)
	fmt.Printf("  Subtotal: S/ %.2f | IGV: S/ %.2f | Total: S/ %.2f\n", subtotal, igv, total)
}

func main() {
	// TODO 7: declara la constante empresa y la variable totalEnvios, e imprímelas
	const empresa = "Encomiendas USS"
	totalEnvios := 3
	fmt.Printf("%s - envíos a registrar: %d\n", empresa, totalEnvios)

	registrarEnvio("ENC-001", 3.5, 210, tarifas.Estandar)

	// TODO 8: registra ENC-002 y ENC-003
	registrarEnvio("ENC-002", 12.0, 770, tarifas.Express)
	registrarEnvio("ENC-003", 60.0, 100, tarifas.Prioritario)
}