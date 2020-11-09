package main

import (
	"fmt"
	"time"
)

//Process es una Estructura de un proceso
type Process struct{
	ID uint64
	Count uint64
	prin uint64
}

//ProcessList es una estructura con un slice de Procesos
type ProcessList struct{
	ListaProcesos []Process
}

//Mostrar es una Función de mostrar todos los Process
func (PL *ProcessList) Mostrar() {
		for p := range PL.ListaProcesos{
			fmt.Printf("\nid %d: %d", PL.ListaProcesos[p].ID, PL.ListaProcesos[p].Count)
			time.Sleep(time.Millisecond * 500)
		}
}

//idSet asigna el ID al Process
func (p *Process)idSet(id uint64){
	p.ID = id
}

//Proceso es la función principal que aumenta el contador
func (p *Process) Proceso() {
	for {	
		p.Count = p.Count + 1
		time.Sleep(time.Millisecond * 500)
		
	}
}

func main() {

	opcion:=0//Opción para el Switch
	num := uint64(0)//Contador de cuántos procesos hay
	PL := ProcessList{}//Slice de Process

	for opcion != 9{//Menú
		fmt.Println("1.-Agregar Proceso")
		fmt.Println("2.-Mostrar Proceso")
		fmt.Println("3.-Eliminar Proceso")
		fmt.Println("9.-Salir")
		fmt.Scanln(&opcion)
		switch opcion{
		case 1://Agregar
			fmt.Println("Agregar")
			num++//Incrementa el numero de Procesos existentes
			PL.ListaProcesos = append(PL.ListaProcesos, Process{})//Agrega un process a la lista
			PL.ListaProcesos[num-1].idSet(num)//asigna un id al process nuevo
			go func(){//Empieza la rutina
					PL.ListaProcesos[num-1].Proceso()//Llama a la función principal de la rutina para que corra 
			}()
			break
		case 2://Mostrar
			fmt.Println("Mostrar")
			quit := make(chan bool)//Canal para salir de Mostrar
			input := 1 //Auxiliar para salir
			for input != 0{
			go func(){//Rutina de mostrar hasta que se ingrese 0
				for{
					select{
					case <- quit:
						return
					default:
						PL.Mostrar()
					}
					
				}
			}()

			fmt.Println("Ingrese 0 para Salir al menu.")
			fmt.Scanln(&input)
			
			}
			quit <- true //El canal de salida es verdadero, por lo que se sale de Mostrar
			break
		case 3://Eliminar
			if (num == 0){
				fmt.Println("nada que eliminar")
			}else{
				fmt.Println("Eliminar")
				num--
			}	
			break
		case 9://Salir
			return
			break
		default:
			fmt.Println("Seleccione una Opción Valida")
			break
		}
	}


}