package arrayOperationService

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const urlCheckService = "https://kuvaev-ituniversity.vps.elewise.com"
const apiNameRotationArray = "Циклическая ротация"
const apiNameFindLoner = "Чудные вхождения в массив"
const apiNameCheckSequence = "Проверка последовательности"
const apiNameSkipElement = "Поиск отсутствующего элемента"

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
	resp, err := http.Get(urlCheckService + "/tasks/" + urlPath)
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

	var arrayInputDataTask []inputDataTask

	for _, data := range inpudData {
		var currentInputDataTask inputDataTask
		if len(data.([]interface{})) > 1 {
			currentInputDataTask.countStart = data.([]interface{})[1].(float64)
		}
		for _, i := range data.([]interface{})[0].([]interface{}) {
			currentInputDataTask.data = append(currentInputDataTask.data, i.(float64))
		}
		arrayInputDataTask = append(arrayInputDataTask, currentInputDataTask)
	}

	return arrayInputDataTask

}

func requestPost(nameApi string, inputArray []interface{}, resultArray []interface{}) []byte {
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
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(httpPostUrl, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func doTask(nameApi string, w http.ResponseWriter, wg ...*sync.WaitGroup) {

	// Как же тут больно передавать необязательные параметры(
	defer func() {
		if len(wg) > 0 {
			wg[0].Done()
		}
	}()

	inputDataTask := getInputData(nameApi)

	var elementForInputArray []interface{}
	var resultArray []interface{}
	var inputArray []interface{}

	for _, data := range inputDataTask {
		elementForInputArray := append(elementForInputArray, data.data)
		if data.countStart > 0 {
			elementForInputArray = append(elementForInputArray, data.countStart)
		}
		inputArray = append(inputArray, elementForInputArray)
		switch nameApi {
		case apiNameRotationArray:
			resultArray = append(resultArray, arrayRotation(convertFloat64ToInt(data.data), int(data.countStart)))
		case apiNameFindLoner:
			resultArray = append(resultArray, arrayFindLoner(convertFloat64ToInt(data.data)))
		case apiNameCheckSequence:
			resultArray = append(resultArray, arrayCheckSequence(convertFloat64ToInt(data.data)))
		case apiNameSkipElement:
			resultArray = append(resultArray, arrayFindSkipEelement(convertFloat64ToInt(data.data)))
		}
	}
	response := requestPost(nameApi, inputArray, resultArray)
	w.Write([]byte("\n" + "Ответ по задаче: " + nameApi + "\n" + string(response)))
}

func (s serviceWorkArray) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/task/" + apiNameRotationArray:
		doTask(apiNameRotationArray, w)
	case "/task/" + apiNameFindLoner:
		doTask(apiNameFindLoner, w)
	case "/task/" + apiNameCheckSequence:
		doTask(apiNameCheckSequence, w)
	case "/task/" + apiNameSkipElement:
		doTask(apiNameSkipElement, w)
	case "/tasks":
		var wg sync.WaitGroup
		wg.Add(4)
		go doTask(apiNameRotationArray, w, &wg)
		go doTask(apiNameFindLoner, w, &wg)
		go doTask(apiNameCheckSequence, w, &wg)
		go doTask(apiNameSkipElement, w, &wg)
		wg.Wait()
	}

}
