package canvas

import (
	"testing"
)

func TestNewLayerSystem(t *testing.T) {
	ls := NewLayerSystem()
	
	if ls.GetLayerCount() != 0 {
		t.Errorf("Expected 0 layers, got %d", ls.GetLayerCount())
	}
}

func TestAddLayer(t *testing.T) {
	ls := NewLayerSystem()
	
	id1 := ls.AddLayer()
	if id1 != 1 {
		t.Errorf("Expected first layer ID to be 1, got %d", id1)
	}
	
	id2 := ls.AddLayer()
	if id2 != 2 {
		t.Errorf("Expected second layer ID to be 2, got %d", id2)
	}
	
	if ls.GetLayerCount() != 2 {
		t.Errorf("Expected 2 layers, got %d", ls.GetLayerCount())
	}
}

func TestGetLayer(t *testing.T) {
	ls := NewLayerSystem()
	id := ls.AddLayer()
	
	layer, err := ls.GetLayer(id)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if layer.ID != id {
		t.Errorf("Expected layer ID %d, got %d", id, layer.ID)
	}
	
	if layer.Canvas == nil {
		t.Error("Expected layer to have a canvas")
	}
}

func TestGetLayerNotFound(t *testing.T) {
	ls := NewLayerSystem()
	
	_, err := ls.GetLayer(999)
	if err == nil {
		t.Error("Expected error for non-existent layer")
	}
}

func TestRemoveLayer(t *testing.T) {
	ls := NewLayerSystem()
	id := ls.AddLayer()
	
	err := ls.RemoveLayer(id)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if ls.GetLayerCount() != 0 {
		t.Errorf("Expected 0 layers after removal, got %d", ls.GetLayerCount())
	}
	
	_, err = ls.GetLayer(id)
	if err == nil {
		t.Error("Expected error when getting removed layer")
	}
}

func TestSetZOrder(t *testing.T) {
	ls := NewLayerSystem()
	id := ls.AddLayer()
	
	err := ls.SetZOrder(id, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	layer, _ := ls.GetLayer(id)
	if layer.ZOrder != 10 {
		t.Errorf("Expected ZOrder 10, got %d", layer.ZOrder)
	}
}

func TestComposite(t *testing.T) {
	ls := NewLayerSystem()
	
	id1 := ls.AddLayer()
	layer1, _ := ls.GetLayer(id1)
	layer1.Canvas.ReplaceChar(0, 0, 'A')
	
	id2 := ls.AddLayer()
	layer2, _ := ls.GetLayer(id2)
	layer2.Canvas.ReplaceChar(1, 0, 'B')
	layer2.Canvas.ReplaceChar(0, 0, 'C')
	ls.SetZOrder(id2, 1)
	
	result := ls.Composite()
	
	char, _ := result.GetChar(0, 0)
	if char != 'C' {
		t.Errorf("Expected 'C' at (0,0) due to higher Z-order, got %c", char)
	}
	
	char, _ = result.GetChar(1, 0)
	if char != 'B' {
		t.Errorf("Expected 'B' at (1,0), got %c", char)
	}
}

func TestCompositeEmptyLayers(t *testing.T) {
	ls := NewLayerSystem()
	
	result := ls.Composite()
	if result == nil {
		t.Error("Expected composite to return a canvas even with no layers")
	}
	
	if result.String() != "" {
		t.Error("Expected empty composite result")
	}
}

func TestGetLayerIDs(t *testing.T) {
	ls := NewLayerSystem()
	
	id1 := ls.AddLayer()
	id2 := ls.AddLayer()
	
	ids := ls.GetLayerIDs()
	if len(ids) != 2 {
		t.Errorf("Expected 2 IDs, got %d", len(ids))
	}
	
	found1, found2 := false, false
	for _, id := range ids {
		if id == id1 {
			found1 = true
		}
		if id == id2 {
			found2 = true
		}
	}
	
	if !found1 || !found2 {
		t.Error("Expected to find both layer IDs in the result")
	}
}