package tickets

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

var sliceTexto [][]string
var pathGeneral = "./tickets.csv"

/*func ReadFiles() {
	lec, err := os.ReadFile("./tickets.csv")
	sliceTexto = strings.Split(texto, ",")

	if err != nil {
		panic("No encontré archivo")
	}
	//fmt.Printf("%s Leyó: ", lec)
	fmt.Println("Array: ", sliceTexto[1])
	//fmt.Println("tam: ", (len(sliceTexto) / 5))

}*/

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func CrearStruct() {

	sliceTexto = ReadCsvFile(pathGeneral)
	filas := (len(sliceTexto))
	fmt.Println("filas:", filas)
	var sliceStruct []Ticket
	//var err error

	//

	/*for i := 0; i < filas; i++ { //filas
		for j := 0; j < 5; j++ { //Cloumnas

			sliceStruct[i].Id, err = strconv.Atoi(sliceTexto[i][j])
			if err != nil {
				panic("Eror a covertir entero")
			}
			sliceStruct[i].Nombre = (sliceTexto[i][j])
			sliceStruct[i].Email = (sliceTexto[i][j])
			sliceStruct[i].PaisDestino = (sliceTexto[i][j])
			sliceStruct[i].HoraVuelo = (sliceTexto[i][j])
			sliceStruct[i].Precio, err = strconv.Atoi(sliceTexto[i][j])
			if err != nil {
				panic("Eror a covertir entero")
			}
		}
	}*/

	sliceStruct[0].Nombre = (sliceTexto[0][1])
	fmt.Println(sliceStruct[0].Nombre)

}

func Atoi(s string) {
	panic("unimplemented")
}

// ejemplo 1
/*func GetTotalTickets(destination string) (int error) {

	var cantidad int = 0

}

// ejemplo 2
func GetMornings(time string) (int error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int error) {}
*/
