package tickets

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"strconv"
)

type Ticket struct {
	Id          int
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      int
}

type FlyByHour struct {
	Madrugada int
	Mañana    int
	Tarde     int
	Noche     int
}

func Init() {
	var pathGeneral = "./tickets.csv"
	var paisBuscar = "Indonesia"
	var paisProm = "Indonesia"

	sliceTexto := ReadCsvFile(pathGeneral)
	sliceStruct := CrearStruct(sliceTexto)
	//total, err := tickets.GetTotalTickets("Brazil")

	////// Requerimiento 1 ////////////
	fmt.Println("////////Pasajeros segun su destino : ////////")
	fmt.Println("")
	cantidad := GetTotalTickets(paisBuscar, sliceStruct)
	fmt.Println("Pasajeros con destino a ", paisBuscar, ": ", cantidad)
	fmt.Println("")

	////// Requerimiento 2 ////////////
	structByHour, _ := GetTravelByTime(sliceStruct)
	fmt.Println("////////Pasajeros por Hora de Viaje: ////////")
	fmt.Println("")
	fmt.Println("Madrugada:", structByHour.Madrugada)
	fmt.Println("Mañana:", structByHour.Mañana)
	fmt.Println("Tarde:", structByHour.Tarde)
	fmt.Println("Noche:", structByHour.Noche)
	fmt.Println("")

	////// Requerimiento 3 ////////////

	var promedio float64
	fmt.Println("////////Promedio viajes segun pais: ////////")
	fmt.Println("")
	cant := GetTotalTickets(paisProm, sliceStruct)
	promedio = PromedioDia(cant, sliceStruct)
	fmt.Println("El promedio viajes de ", paisProm, "es: ", promedio, "%")

}

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("No pudo leer archivo "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func CrearStruct(matrizLeida [][]string) (sliceStruct []Ticket) { //crear Slice de Struct

	filas := (len(matrizLeida))
	colum := (len(matrizLeida[0]))

	var err error

	for i := 0; i < filas; i++ { //filas

		var structOnlyTicket Ticket
		for j := 0; j < colum; j++ { //Columnas

			switch j {

			case 0:
				structOnlyTicket.Id, err = strconv.Atoi(matrizLeida[i][j])
				if err != nil {
					panic("Error a covertir entero")
				}
			case 1:

				structOnlyTicket.Nombre = (matrizLeida[i][j])
			case 2:

				structOnlyTicket.Email = (matrizLeida[i][j])
			case 3:

				structOnlyTicket.PaisDestino = (matrizLeida[i][j])
			case 4:

				structOnlyTicket.HoraVuelo = (matrizLeida[i][j])
			case 5:

				structOnlyTicket.Precio, err = strconv.Atoi(matrizLeida[i][j])
				if err != nil {
					panic("Error a covertir entero")
				}
				//default:

			}
		}
		sliceStruct = append(sliceStruct, structOnlyTicket)

	}

	return

}

func Atoi(s string) {
	panic("unimplemented")
}

// Requerimiento 1

func GetTotalTickets(destination string, sliceSt []Ticket) (cant int) {

	cant = 0

	for _, data := range sliceSt {

		if data.PaisDestino == destination {
			cant += 1
		}

	}

	return

}

// Requerimiento 2
func GetTravelByTime(listaTicket []Ticket) (structByHour FlyByHour, err error) {
	for _, tick := range listaTicket {
		hora := strings.Split(tick.HoraVuelo, ":")
		horaInt, err := strconv.Atoi(hora[0])
		if err != nil {
			panic(err)
		}
		switch {
		case horaInt == 0 || horaInt <= 6:
			structByHour.Madrugada += 1
		case horaInt == 7 || horaInt <= 12:
			structByHour.Mañana += 1
		case horaInt == 13 || horaInt <= 19:
			structByHour.Tarde += 1
		default:
			structByHour.Noche += 1
		}
	}

	return structByHour, nil

}

// Requerimiento 3
func PromedioDia(totalDestino int, sliceSt []Ticket) (prom float64) {

	total := len(sliceSt)
	prom = (float64(totalDestino) / float64(total)) * 100

	return

}
