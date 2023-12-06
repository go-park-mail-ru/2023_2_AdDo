package in_memory

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type InMemory struct {
	Data      [][]float64 `json:"data"`
	Centroids [][]float64 `json:"centroids"`
	Labels    []int       `json:"labels"`
}

func NewInMemory(pathToDump string, l *logrus.Logger) (InMemory, error) {
	data, err := os.ReadFile(pathToDump)
	if err != nil {
		l.Errorln("error opening file", err, pathToDump)
		return InMemory{}, err
	}

	var clusteringData InMemory
	err = json.Unmarshal(data, &clusteringData)
	if err != nil {
		l.Errorln("error parsing json", err, pathToDump)
		return InMemory{}, err
	}
	l.Infoln("json parsed successfully")

	return clusteringData, nil
}
