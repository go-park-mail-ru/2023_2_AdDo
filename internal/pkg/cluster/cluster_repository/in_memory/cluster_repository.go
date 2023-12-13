package cluster_repository

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"main/internal/pkg/track"
	"math"
	"os"
	"sort"
)

type InMemory struct {
	TrackIdToDataIndex   map[uint64][]float64
	ClusterToDataIndexes map[int][]uint64
	Centroids            [][]float64
	logger               *logrus.Logger
}

// search nearest tracks flow:
// 	1. Get for every track id some nearest tracks in loop
// 	2. Get nearest cluster for track by using euclidian between objectVec and centroids
//  3. Get all tracks in chosen cluster
//  4. for every chosen track find euclidian and sort asc
//  5. Take first N track and return
//  6. Concat all results for every track and return

func calculateEuclid(obj1, obj2 []float64) float64 {
	sum := 0.0

	for i := 0; i < len(obj2); i++ {
		diff := obj1[i] - obj2[i]
		sum += diff * diff
	}

	distance := math.Sqrt(sum)
	return distance
}

func (in InMemory) getNearestClusterData(objectVec []float64) []uint64 {
	in.logger.Infoln("get nearest cluster entered")
	foundIndex := 0
	minDistance := math.MaxFloat64
	for index, centroid := range in.Centroids {
		if calculateEuclid(objectVec, centroid) < minDistance {
			foundIndex = index
		}
	}
	in.logger.Infoln("found cluster for our objs", foundIndex)

	return in.ClusterToDataIndexes[foundIndex]
}

type IdDistance struct {
	Id       uint64
	Distance float64
}

func (in InMemory) getNearestTracks(id track.Id, count int) ([]track.Id, error) {
	in.logger.Infoln("Get Nearest for one track entered with id ", id)

	trackVec := in.TrackIdToDataIndex[id.Id]
	tracksIds := in.getNearestClusterData(trackVec)

	vec := make([]IdDistance, 0)
	for _, trackId := range tracksIds {
		vec = append(vec, IdDistance{
			Id:       trackId,
			Distance: calculateEuclid(trackVec, in.TrackIdToDataIndex[trackId]),
		})
	}

	sort.Slice(vec, func(i, j int) bool {
		return vec[i].Distance < vec[j].Distance
	})
	in.logger.Infoln("sorted vec with dist ", vec)

	result := make([]track.Id, 0)
	for index, t := range vec {
		if index >= count {
			break
		}
		result = append(result, track.Id{Id: t.Id})
	}
	in.logger.Infoln("final result for one track", result)

	return result, nil
}

func (in InMemory) GetNearestTracks(ids []track.Id, countPerTrack int) ([]track.Id, error) {
	in.logger.Infoln("Get Nearest tracks entered with", ids, "for every track count: ", countPerTrack)
	result := make([]track.Id, 0)
	for _, id := range ids {
		nearestForTrack, err := in.getNearestTracks(id, countPerTrack)
		if err != nil {
			return nil, err
		}
		result = append(result, nearestForTrack...)
	}

	in.logger.Infoln("Get Nearest finished with", result)
	return result, nil
}

type ParseFromFile struct {
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
	var dataFromFile ParseFromFile
	err = json.Unmarshal(data, &dataFromFile)
	if err != nil {
		l.Errorln("error parsing json", err, pathToDump)
		return InMemory{}, err
	}
	l.Infoln("json parsed successfully")

	var result InMemory

	result.Centroids = dataFromFile.Centroids
	result.TrackIdToDataIndex = make(map[uint64][]float64)
	result.ClusterToDataIndexes = make(map[int][]uint64)

	for _, row := range dataFromFile.Data {
		result.TrackIdToDataIndex[uint64(row[len(row)-1])] = row[:len(row)-1]
	}

	for index, label := range dataFromFile.Labels {
		result.ClusterToDataIndexes[label] = append(result.ClusterToDataIndexes[label], uint64(dataFromFile.Data[index][len(dataFromFile.Data[index])-1]))
	}
	l.Infoln("InMemory clusters formed", result.Centroids)

	result.logger = l
	return result, nil
}
