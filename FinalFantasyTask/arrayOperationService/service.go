package arrayOperationService

import (
	"bytes"
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
		if len(data.([]interface{})) > 1 {
			current.countStart = data.([]interface{})[1].(float64)
		}
		for _, i := range data.([]interface{})[0].([]interface{}) {
			current.data = append(current.data, i.(float64))
		}
		VinputData = append(VinputData, current)
	}
	//Разобраться потом в этом, сортировка работает - но умом пока не понимаю
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

func requestPost(nameApi string, inputArray []interface{}, resultArray []interface{}) {
	httpPostUrl := urlCheckService + "/tasks/solution"

	message := map[string]interface{}{
		"user_name": "IoganVII",
		"task":      nameApi,
		"results": map[string]interface{}{
			"payload": inputArray,
			"results": resultArray,
		},
	}

	bytesRepresentation, err := json.Marshal(message)

	//fmt.Println(string(bytesRepresentation))

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(httpPostUrl, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}

func (s serviceWorkArray) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/tasks/Циклическая ротация" {
		inputDataTask := getInputData(r.URL.Path)

		var elementForInputArray []interface{}
		var resultArray []interface{}
		var inputArray []interface{}

		for _, data := range inputDataTask {
			elementForInputArray := append(elementForInputArray, data.data)
			elementForInputArray = append(elementForInputArray, data.countStart)
			inputArray = append(inputArray, elementForInputArray)
			resultArray = append(resultArray, ArrayRotation(convertFloat64ToInt(data.data), int(data.countStart)))
		}

		//fmt.Println(inputArray, resultArray)

		requestPost("Циклическая ротация", inputArray, resultArray)

	}

	if r.URL.Path == "/tasks/Чудные вхождения в массив" {
		inputDataTask := getInputData(r.URL.Path)

		var elementForInputArray []interface{}
		var resultArray []interface{}
		var inputArray []interface{}

		for _, data := range inputDataTask {
			elementForInputArray := append(elementForInputArray, data.data)
			if data.countStart > 0 {
				elementForInputArray = append(elementForInputArray, data.countStart)
			}
			inputArray = append(inputArray, elementForInputArray)
			resultArray = append(resultArray, ArrayFindLoner(convertFloat64ToInt(data.data)))
		}

		//fmt.Println(inputArray, resultArray)

		requestPost("Чудные вхождения в массив", inputArray, resultArray)

	}

	if r.URL.Path == "/tasks/Проверка последовательности" {
		inputDataTask := getInputData(r.URL.Path)

		var elementForInputArray []interface{}
		var resultArray []interface{}
		var inputArray []interface{}

		for _, data := range inputDataTask {
			elementForInputArray := append(elementForInputArray, data.data)
			if data.countStart > 0 {
				elementForInputArray = append(elementForInputArray, data.countStart)
			}
			inputArray = append(inputArray, elementForInputArray)
			resultArray = append(resultArray, ArrayCheckSequence(convertFloat64ToInt(data.data)))
		}

		//fmt.Println(inputArray, resultArray)

		requestPost("Проверка последовательности", inputArray, resultArray)

	}

	if r.URL.Path == "/tasks/Поиск отсутствующего элемента" {
		inputDataTask := getInputData(r.URL.Path)

		var elementForInputArray []interface{}
		var resultArray []interface{}
		var inputArray []interface{}

		for _, data := range inputDataTask {
			elementForInputArray := append(elementForInputArray, data.data)
			if data.countStart > 0 {
				elementForInputArray = append(elementForInputArray, data.countStart)
			}
			inputArray = append(inputArray, elementForInputArray)
			resultArray = append(resultArray, ArrayFindSkipEelement(convertFloat64ToInt(data.data)))
		}

		//fmt.Println(inputArray, resultArray)

		requestPost("Поиск отсутствующего элемента", inputArray, resultArray)

	}

}
