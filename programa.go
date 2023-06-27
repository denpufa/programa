// Gerencia do programa  em contato com o usuário
package programa

import (
	"fmt"
	"github.com/denpufa/volume"
	"github.com/denpufa/area"
	"github.com/denpufa/perimetro"
)


func Executar(){

	for {
		var digit int
		fmt.Println("Calculadora de Geometria Plana e Espacial")
		fmt.Println("(1) Triângulo equilátero")
		fmt.Println("(2) Retângulo")
		fmt.Println("(3) Quadrado")
		fmt.Println("(4) Círculo")
		fmt.Println("(5) Pirâmide com base quadrangular")
		fmt.Println("(6) Cubo")
		fmt.Println("(7) Paralelepípedo")
		fmt.Println("(8) Esfera")
		fmt.Println("(0) Sair")
		fmt.Print("Digite a sua opção: ")
		fmt.Scanln(&digit)

		switch digit {
		case 1:
			var base, altura float32
			fmt.Print("\nDigite a base do triângulo: ")
			fmt.Scanln(&base)
			fmt.Print("\nDigite a altura do triângulo: ")
			fmt.Scanln(&altura)

			area,err_a := area.CalcularTrianguloArea(base,altura)
			perimetro,err_p := perimetro.CalcularPerimetroTriangulo(base)
			if err_a != nil {
					fmt.Printf("%v\n", err)
					continue
				}
			if err_p != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			fmt.Printf("Àrea do triângulo: %v\n",area)
			fmt.Printf("Perimetro do triângulo: %v\n", perimetro)

		case 2:
			var base, altura float32
			fmt.Print("\nDigite a base do retângulo: ")
			fmt.Scanln(&base)
			fmt.Print("\nDigite a altura do retângulo: ")
			fmt.Scanln(&altura)
			area,err_a := area.CalcularRetanguloArea(base,altura)
			perimetro,err_p := perimetro.CalcularPerimetroRetangulo(base,altura)
			if err_a != nil {
					fmt.Printf("%v\n", err)
					continue
				}
			if err_p != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			fmt.Printf("Àrea do retângulo: %v\n",area)
			fmt.Printf("Perimetro do retângulo: %v\n", perimetro)

		case 3:
			var lado float32
			fmt.Print("\nDigite o lado do quadrado: ")
			fmt.Scanln(&lado)
			area,err_a := area.CalcularQuadradoArea(lado)
			perimetro,err_p := perimetro.CalcularPerimetroQuadrado(lado)
			if err_a != nil {
					fmt.Printf("%v\n", err)
					continue
				}
			if err_p != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			fmt.Printf("Àrea do quadrado: %v\n",area)
			fmt.Printf("Perimetro do quadrado: %v\n", perimetro)


		case 4:

			var raio float32
			fmt.Print("\nDigite o raio do círculo: ")
			fmt.Scanln(&raio)
			area,err_a := area.CalcularCirculoArea(raio)
			perimetro,err_p := perimetro.CalcularPerimetroCirculo(raio)
			if err_a != nil {
					fmt.Printf("%v\n", err)
					continue
				}
			if err_p != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			fmt.Printf("Àrea do circulo: %v\n",area)
			fmt.Printf("Perimetro do circulo: %v\n", perimetro)
		case 5:

			var baseQuadrado, faceTriangulo, alturaTriangulo, alturaPiramide float32
			fmt.Print("\nDigite o tamanho do lado da base do quadrado da pirâmide: ")
			fmt.Scanln(&baseQuadrado)
			fmt.Print("\nDigite o tamanho do lado  do triângulo de face da pirâmide: ")
			fmt.Scanln(&baseTriangulo)
			fmt.Print("\nDigite a altura do triângulo lateral da pirâmide: ")
			fmt.Scanln(&alturaTriangulo)
			fmt.Print("\nDigite a altura da pirâmide: ")
			fmt.Scanln(&alturaPiramide)

			volume,err_v := volume.CalcularVolumePiramide(baseQuadrado,alturaPiramide)
			area,err_a := area.CalcularPiramideArea(baseQuadrado,faceTriangulo,alturaTriangulo)

			if err_v != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			if err_a != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			fmt.Printf("O valor do volume da pirâmide é: %v\n",volume)
			fmt.Printf("O valor da àrea da pirâmide é: %v\n", area)

		case 6:
			var ladoQuadrado float64
			fmt.Print("\nDigite o lado do quadrado da base do cubo: ")
			fmt.Scanln(&ladoQuadrado)

			volume,err_v := volume.CalcularVolumeCubo(ladoQuadrado)
			area,err_a := area.CalcularCuboArea(ladoQuadrado)

			if err_v != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			if err_a != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			fmt.Printf("O valor do volume do cubo é: %v\n",volume)
			fmt.Printf("O valor da àrea do cubo é: %v\n", area)
		case 7:
			var altura, largura, comprimento float64
			fmt.Print("Digite a altura do paralelepípedo: ")
			fmt.Scanln(&altura)
			fmt.Print("Digite a largura do paralelepípedo: ")
			fmt.Scanln(&largura)
			fmt.Print("Digite o comprimento do paralelepípedo: ")
			fmt.Scanln(&comprimento)

			volume,err_v := volume.CalcularVolumeParalelepipedo(comprimento,largura,alturaBase)
			area,err_a := area.CalcularParalelepipedoArea(comprimento,largura,alturaBase)

			if err_v != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			if err_a != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			fmt.Printf("O valor do volume do paralelepipedo é: %v\n",volume)
			fmt.Printf("O valor da àrea do paralelpidedo é: %v\n", area)

		case 8:
			var raio float32
			fmt.Print("Digite o raio da esfera: ")
			fmt.Scanln(&raio)

			volume,err_v := volume.CalcularVolumeEsfera(raio)
			area,err_a := area.CalcularEsferaArea(raio)

			if err_v != nil {
				fmt.Printf("%v\n", err)
					continue
			}
			if err_a != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			fmt.Printf("O valor do volume da esfera é: %v\n",volume)
			fmt.Printf("O valor da àrea da esfera é: %v\n", area)


		case 0:
			return
		default:
			fmt.Println("Opção inválida")
			continue
		}
	}
}

