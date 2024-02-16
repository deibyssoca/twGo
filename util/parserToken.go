package util

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func ObtenerToken(txtToken string) string {
	var desde, hasta int
	var salida string
	desde = strings.Index(txtToken, "access_token=")
	hasta = strings.Index(txtToken[desde:], "&token_type=")
	if desde != -1 && hasta != -1 {
		//fmt.Println(txtToken[desde+13 : desde+hasta])
		//salida := txtToken[desde+12 : desde+hasta]
		//tf = t[desde+12 : desde+hasta]
		// Copiar texto al portapapeles
		// Dividir la URL en partes utilizando "&" como separador
		parametros := strings.Split(txtToken[desde:], "&")
		//println("===>>>" + parametros[2])
		// Iterar sobre los parámetros e imprimirlos
		// for _, parametro := range parametros {
		// 	fmt.Println(parametro)
		// }

		salida = txtToken[desde+13:desde+hasta] + "&" + parametros[2]
		err := clipboard.WriteAll(salida)
		if err != nil {
			fmt.Println("Error al copiar texto al portapapeles:", err)
		}
	} else {
		fmt.Println("No se encontró el token en la URL.")
	}
	return salida
}
