package cluster_repository

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"main/internal/pkg/track"
	"testing"
)

func TestGettingNearestTracks(t *testing.T) {
	l := logrus.New()
	mockCluster := InMemory{
		TrackIdToDataIndex:   map[uint64][]float64{0: []float64{1, 1, 1}, 1: []float64{0, 1, 0}},
		ClusterToDataIndexes: map[int][]uint64{0: []uint64{0}, 1: []uint64{1}},
		Centroids:            [][]float64{{1, 1, 1}, {0, 1, 0}},
		logger:               l,
	}

	result, err := mockCluster.GetNearestTracks([]track.Id{track.Id{Id: 0}}, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, []track.Id{{Id: 0}}, result)
}

func TestFilteringSkips(t *testing.T) {
	l := logrus.New()
	mockCluster := InMemory{
		TrackIdToDataIndex:   map[uint64][]float64{0: []float64{1, 1, 1}, 1: []float64{0, 1, 0}},
		ClusterToDataIndexes: map[int][]uint64{0: []uint64{0}, 1: []uint64{1}},
		Centroids:            [][]float64{{1, 1, 1}, {0, 1, 0}},
		logger:               l,
	}

	result, err := mockCluster.GetNearestTracks([]track.Id{track.Id{Id: 0}}, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, []track.Id{{Id: 0}}, result)

	result = mockCluster.FilterSkips(result, result)
	assert.Equal(t, []track.Id{}, result)
}
