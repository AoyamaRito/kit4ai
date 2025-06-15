package canvas

import (
	"errors"
	"fmt"
	"sort"
)

type Layer struct {
	ID     int
	Name   string
	Canvas *Canvas
	ZOrder int
	OffsetX int
	OffsetY int
}

type LayerSystem struct {
	layers   []*Layer
	nextID   int
	maxOrder int
}

func NewLayerSystem() *LayerSystem {
	return &LayerSystem{
		layers:   make([]*Layer, 0),
		nextID:   1,
		maxOrder: 0,
	}
}

func (ls *LayerSystem) AddLayer() int {
	return ls.AddLayerWithName(fmt.Sprintf("Layer %d", ls.nextID))
}

func (ls *LayerSystem) AddLayerWithName(name string) int {
	layer := &Layer{
		ID:      ls.nextID,
		Name:    name,
		Canvas:  NewCanvas(),
		ZOrder:  ls.maxOrder,
		OffsetX: 0,
		OffsetY: 0,
	}
	
	ls.layers = append(ls.layers, layer)
	ls.nextID++
	ls.maxOrder++
	
	return layer.ID
}

// AddTextLayer 全角文字対応のテキストレイヤーを追加
func (ls *LayerSystem) AddTextLayer(name string) int {
	textLayer := NewTextLayer()
	layer := &Layer{
		ID:      ls.nextID,
		Name:    name,
		Canvas:  textLayer.Canvas,
		ZOrder:  ls.maxOrder,
		OffsetX: 0,
		OffsetY: 0,
	}
	
	ls.layers = append(ls.layers, layer)
	ls.nextID++
	ls.maxOrder++
	
	return layer.ID
}

// GetTextLayer テキストレイヤーとして取得
func (ls *LayerSystem) GetTextLayer(id int) (*TextLayer, error) {
	layer, err := ls.GetLayer(id)
	if err != nil {
		return nil, err
	}
	
	return &TextLayer{Canvas: layer.Canvas}, nil
}

func (ls *LayerSystem) RemoveLayer(id int) error {
	for i, layer := range ls.layers {
		if layer.ID == id {
			ls.layers = append(ls.layers[:i], ls.layers[i+1:]...)
			return nil
		}
	}
	return errors.New("layer not found")
}

func (ls *LayerSystem) GetLayer(id int) (*Layer, error) {
	for _, layer := range ls.layers {
		if layer.ID == id {
			return layer, nil
		}
	}
	return nil, errors.New("layer not found")
}

func (ls *LayerSystem) SetZOrder(id int, order int) error {
	layer, err := ls.GetLayer(id)
	if err != nil {
		return err
	}
	
	layer.ZOrder = order
	if order > ls.maxOrder {
		ls.maxOrder = order
	}
	
	return nil
}

func (ls *LayerSystem) Composite() *Canvas {
	if len(ls.layers) == 0 {
		return NewCanvas()
	}
	
	// Z-Orderでソート（小さい順 = 下層から上層へ）
	sortedLayers := make([]*Layer, len(ls.layers))
	copy(sortedLayers, ls.layers)
	
	sort.Slice(sortedLayers, func(i, j int) bool {
		return sortedLayers[i].ZOrder < sortedLayers[j].ZOrder
	})
	
	// 80文字 x 100行の配列を作成
	result := NewCanvas()
	
	// 各レイヤーを下層から上層へ順番に合成
	for _, layer := range sortedLayers {
		for y := 0; y < DefaultHeight; y++ {
			for x := 0; x < DefaultWidth; x++ {
				sourceChar, err := layer.Canvas.GetChar(x, y)
				if err != nil {
					continue
				}
				
				// オフセットを適用した配置位置
				targetX := x + layer.OffsetX
				targetY := y + layer.OffsetY
				
				// 境界チェック
				if targetX < 0 || targetX >= DefaultWidth || targetY < 0 || targetY >= DefaultHeight {
					continue
				}
				
				// 現在の配置位置にある文字を取得
				currentChar, err := result.GetChar(targetX, targetY)
				if err != nil {
					continue
				}
				
				// 合成ルール：
				// 1. 現在位置がスペース -> 新しい文字を配置
				// 2. 現在位置に文字がある -> 上位レイヤーの文字を優先（上書き）
				// 3. 新しい文字がスペース -> 既存文字を保持
				if sourceChar != ' ' || currentChar == ' ' {
					result.ReplaceChar(targetX, targetY, sourceChar)
				}
			}
		}
	}
	
	return result
}

func (ls *LayerSystem) GetLayerCount() int {
	return len(ls.layers)
}

func (ls *LayerSystem) GetLayerIDs() []int {
	ids := make([]int, len(ls.layers))
	for i, layer := range ls.layers {
		ids[i] = layer.ID
	}
	return ids
}

func (ls *LayerSystem) SetLayerName(id int, name string) error {
	layer, err := ls.GetLayer(id)
	if err != nil {
		return err
	}
	
	layer.Name = name
	return nil
}

func (ls *LayerSystem) MoveLayer(id int, offsetX, offsetY int) error {
	layer, err := ls.GetLayer(id)
	if err != nil {
		return err
	}
	
	layer.OffsetX = offsetX
	layer.OffsetY = offsetY
	return nil
}

func (ls *LayerSystem) GetLayerInfo(id int) (string, int, int, int, error) {
	layer, err := ls.GetLayer(id)
	if err != nil {
		return "", 0, 0, 0, err
	}
	
	return layer.Name, layer.ZOrder, layer.OffsetX, layer.OffsetY, nil
}

func (ls *LayerSystem) ListLayers() {
	fmt.Println("レイヤー一覧:")
	fmt.Println("ID | Name | Z-Order | Offset")
	fmt.Println("---|------|---------|--------")
	
	for _, layer := range ls.layers {
		fmt.Printf("%d  | %s | %d       | (%d,%d)\n", 
			layer.ID, layer.Name, layer.ZOrder, layer.OffsetX, layer.OffsetY)
	}
}