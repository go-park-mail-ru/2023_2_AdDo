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

type IdDistance struct {
	Id       uint64
	Distance float64
}

func (in InMemory) getNearestClusterData(objectVec []float64) []IdDistance {
	in.logger.Infoln("get nearest cluster entered")
	centroidsSortedVec := make([]IdDistance, 0)
	for index, centroid := range in.Centroids {
		centroidsSortedVec = append(centroidsSortedVec, IdDistance{Id: uint64(index), Distance: calculateEuclid(objectVec, centroid)})
	}

	sort.Slice(centroidsSortedVec, func(i, j int) bool {
		return centroidsSortedVec[i].Distance < centroidsSortedVec[j].Distance
	})
	in.logger.Infoln("sorted clusters")

	return centroidsSortedVec
}

func (in InMemory) getNearestTracks(uniqTracks map[uint64]bool, id track.Id, count int) ([]track.Id, error) {
	in.logger.Infoln("Get Nearest for one track entered with id ", id)

	trackVec := in.TrackIdToDataIndex[id.Id]
	sortedClusters := in.getNearestClusterData(trackVec)

	vec := make([]IdDistance, 0)
	for _, cluster := range sortedClusters {
		for _, trackId := range in.ClusterToDataIndexes[int(cluster.Id)] {
			if len(vec) > 2*count {
				break
			}

			if _, ok := uniqTracks[trackId]; ok {
				continue
			}

			uniqTracks[trackId] = true
			vec = append(vec, IdDistance{
				Id:       trackId,
				Distance: calculateEuclid(trackVec, in.TrackIdToDataIndex[trackId]),
			})
		}
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

func (in InMemory) findNearestInSet(tracks []track.Id, skip track.Id, count int) []track.Id {
	vec := make([]IdDistance, 0)
	for _, t := range tracks {
		vec = append(vec, IdDistance{
			Id:       t.Id,
			Distance: calculateEuclid(in.TrackIdToDataIndex[t.Id], in.TrackIdToDataIndex[skip.Id]),
		})
	}
	sort.Slice(vec, func(i, j int) bool {
		return vec[i].Distance < vec[j].Distance
	})

	result := make([]track.Id, 0)
	for i := 0; i < count; i++ {
		result = append(result, track.Id{Id: vec[i].Id})
	}

	return result
}

func (in InMemory) eraseVecFromVec(tracks []track.Id, skips []track.Id) []track.Id {
	redundantElems := make(map[uint64]bool)
	for _, id := range skips {
		redundantElems[id.Id] = true
	}

	result := make([]track.Id, 0)
	for _, id := range tracks {
		if _, ok := redundantElems[id.Id]; ok {
			continue
		}
		result = append(result, id)
	}

	return result
}

func (in InMemory) FilterSkips(tracks []track.Id, skips []track.Id) []track.Id {
	in.logger.Infoln("Filter skips entered", tracks, skips)

	if len(skips) == 0 {
		return tracks
	}

	skipPerTrack := len(tracks)/2/len(skips) + 1
	redundant := make([]track.Id, 0)
	for _, skip := range skips {
		temp := in.findNearestInSet(tracks, skip, skipPerTrack)
		redundant = append(redundant, temp...)
	}

	result := in.eraseVecFromVec(tracks, redundant)
	in.logger.Infoln("Result", result)
	return result
}

func (in InMemory) GetNearestTracks(ids []track.Id, countPerTrack int) ([]track.Id, error) {
	in.logger.Infoln("Get Nearest tracks entered with", ids, "for every track count: ", countPerTrack)

	uniqueTracks := make(map[uint64]bool)
	in.logger.Infoln("Uniq tracks map created")

	result := make([]track.Id, 0)
	for _, id := range ids {
		nearestForTrack, err := in.getNearestTracks(uniqueTracks, id, countPerTrack)
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
