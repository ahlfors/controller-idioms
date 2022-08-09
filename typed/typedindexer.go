package typed

// NOTE: this is not used yet, API and implementation may change

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

type TypedIndexer[K runtime.Object] struct {
	indexer cache.Indexer
}

func NewTypedIndexer[K runtime.Object](indexer cache.Indexer) *TypedIndexer[K] {
	return &TypedIndexer[K]{indexer: indexer}
}

func (t TypedIndexer[K]) Add(obj K) error {
	return t.indexer.Add(obj)
}

func (t TypedIndexer[K]) Update(obj K) error {
	return t.indexer.Update(obj)
}

func (t TypedIndexer[K]) Delete(obj K) error {
	return t.indexer.Delete(obj)
}

func (t TypedIndexer[K]) List() []K {
	return IndexerListToTypedList[K](t.indexer.List())
}

func (t TypedIndexer[K]) ListKeys() []string {
	return t.indexer.ListKeys()
}

func (t TypedIndexer[K]) Get(obj K) (item K, exists bool, err error) {
	var typedObj *K
	gotItem, gotExists, gotErr := t.indexer.Get(obj)
	if err != nil || !gotExists {
		return *typedObj, gotExists, gotErr
	}
	gotRObj, ok := gotItem.(runtime.Object)
	if !ok {
		return *typedObj, gotExists, fmt.Errorf("%v is not a runtime.Object", gotItem)
	}

	gotTypedObj, err := UnstructuredObjToTypedObj[K](gotRObj)
	if err != nil {
		return *typedObj, gotExists, fmt.Errorf("could not convert %s to %T", gotItem, *typedObj)
	}
	return gotTypedObj, gotExists, gotErr
}

func (t TypedIndexer[K]) GetByKey(key string) (item interface{}, exists bool, err error) {
	var typedObj *K
	gotItem, gotExists, gotErr := t.indexer.GetByKey(key)
	if err != nil || !gotExists {
		return *typedObj, gotExists, gotErr
	}
	gotRObj, ok := gotItem.(runtime.Object)
	if !ok {
		return *typedObj, gotExists, fmt.Errorf("%v is not a runtime.Object", gotItem)
	}

	gotTypedObj, err := UnstructuredObjToTypedObj[K](gotRObj)
	if err != nil {
		return *typedObj, gotExists, fmt.Errorf("could not convert %s to %T", gotItem, *typedObj)
	}
	return gotTypedObj, gotExists, gotErr
}

func (t TypedIndexer[K]) IndexKeys(indexName, indexedValue string) ([]string, error) {
	return t.indexer.IndexKeys(indexName, indexedValue)
}

func (t TypedIndexer[K]) ListIndexFuncValues(indexName string) []string {
	return t.indexer.ListIndexFuncValues(indexName)
}

func (t TypedIndexer[K]) ByIndex(indexName, indexedValue string) ([]K, error) {
	objs, err := t.indexer.ByIndex(indexName, indexedValue)
	if err != nil {
		return nil, err
	}
	return IndexerListToTypedList[K](objs), nil
}

func (t TypedIndexer[K]) GetIndexers() cache.Indexers {
	return t.indexer.GetIndexers()
}

func (t TypedIndexer[K]) AddIndexers(newIndexers cache.Indexers) error {
	return t.indexer.AddIndexers(newIndexers)
}

func IndexerListToTypedList[K runtime.Object](objs []any) []K {
	typedObjs := make([]K, 0, len(objs))
	for _, obj := range objs {
		rObj, ok := obj.(runtime.Object)
		if !ok {
			continue
		}
		typedObj, err := UnstructuredObjToTypedObj[K](rObj)
		if err != nil {
			continue
		}
		typedObjs = append(typedObjs, typedObj)
	}
	return typedObjs
}