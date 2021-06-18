package main

import "fmt"

// HOUSE OMIT
type house struct {
	windowType string
	doorType   string
	floor      int
}

// END OMIT

// iBUILDER OMIT
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}

// END OMIT

// NORMAL_BUILDER OMIT
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}
func (n *normalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}
func (n *normalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}
func (n *normalBuilder) setNumFloor() {
	n.floor = 2
}
func (n *normalBuilder) getHouse() house {
	return house{
		windowType: n.windowType,
		doorType:   n.doorType,
		floor:      n.floor,
	}
}

// END OMIT

// IGLOO_BUILDER OMIT
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}
func (i *iglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}
func (i *iglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}
func (i *iglooBuilder) setNumFloor() {
	i.floor = 1
}
func (i *iglooBuilder) getHouse() house {
	return house{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}

// END OMIT

// DIRECTOR OMIT
type director struct {
	builder iBuilder
}

func newDirector(builder iBuilder) *director {
	return &director{builder: builder}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// END OMIT

// MAIN OMIT
func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}

// END OMIT
