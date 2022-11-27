package arrayOperationService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

const urlCheckService = "https://kuvaev-ituniversity.vps.elewise.com"

type input2 interface {
}

type inputDataTask struct {
	data       []float64
	countStart float64
}

type serviceWorkArray struct {
}

func NewService() serviceWorkArray {
	return serviceWorkArray{}
}

// Функция получает данные от сервиса по конкретной задаче
func getInputData(urlPath string) []inputDataTask {
	resp, err := http.Get(urlCheckService + urlPath)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var inpudData []interface{}

	err = json.Unmarshal(body, &inpudData)
	if err != nil {
		log.Fatalln(err)
	}

	var VinputData []inputDataTask

	for _, data := range inpudData {
		var current inputDataTask
		current.countStart = data.([]interface{})[1].(float64)
		for _, data2 := range data.([]interface{})[0].([]interface{}) {
			current.data = append(current.data, data2.(float64))
		}
		VinputData = append(VinputData, current)
	}
	// Разобраться потом в этом, сортировка работает - но умом пока не понимаю
	sort.Slice(VinputData[:], func(i, j int) bool {
		return VinputData[i].countStart < VinputData[j].countStart
	})

	return VinputData

}

func convertFloat64ToInt(ar []float64) []int {
	newar := make([]int, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int(v)
	}
	return newar
}

func (s serviceWorkArray) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/tasks/Циклическая ротация" {
		inputDataTask := getInputData(r.URL.Path)

		var resultArray [][]int

		for _, data := range inputDataTask {
			result := ArrayRotation(convertFloat64ToInt(data.data), int(data.countStart))
			resultArray = append(resultArray, result)
		}

		fmt.Println(resultArray)

	}
}
