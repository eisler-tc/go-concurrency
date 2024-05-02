package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	watcher := make(map[<-chan struct{}]struct{})

	watcher[make(chan struct{})] = struct{}{}
	watcher[make(chan struct{})] = struct{}{}
	watcher[make(chan struct{})] = struct{}{}
	watcher[make(chan struct{})] = struct{}{}
	watcher[make(chan struct{})] = struct{}{}
	watcher[make(chan struct{})] = struct{}{}

	fmt.Printf("wathcer = %+v", watcher)

}

type FeatureData struct {
}

type DataView2 struct {
}

type FeaturesCache struct {
	mu             sync.RWMutex
	features       map[int8]map[string]FeatureData
	anotherView    map[int8]map[string]DataView2
	oldSearchIndex int8
	newSearchIndex int8 //Just allow to retrive cache data from new Index of the features and anotherView maps
}

func (f *FeaturesCache) CreateNewIndex() int8 {
	f.mu.Lock()
	defer f.mu.Unlock()
	return (f.newSearchIndex + 1) % 16 // mod 16 could be change per your refill rate
}

func (f *FeaturesCache) SetNewIndex(newIndex int8) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.oldSearchIndex = f.newSearchIndex
	f.newSearchIndex = newIndex
}

func (f *FeaturesCache) refill(ctx context.Context) {
	var newSources map[string]FeatureData
	var anotherView map[string]DataView2
	// some querying and processing logic

	// save new data for future queries
	newSearchIndex := f.CreateNewIndex()
	f.features[newSearchIndex] = newSources
	f.anotherView[newSearchIndex] = anotherView
	f.SetNewIndex(newSearchIndex) //Just let the queries to new cached datas after updating search Index
	f.features[f.oldSearchIndex] = nil
	f.anotherView[f.oldSearchIndex] = nil
}
