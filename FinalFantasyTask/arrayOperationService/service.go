package arrayOperationService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

func (s serviceWorkArray) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/tasks/Циклическая ротация" {
		resp, err := http.Get("https://kuvaev-ituniversity.vps.elewise.com" + r.URL.Path)
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

		fmt.Println(string(body))

	}
}
