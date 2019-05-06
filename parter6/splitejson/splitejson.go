package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size float32
	ResX, ResY int
}

type Battery struct {
	Cpacity int
}

func getJsonDate() []byte {

	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen:Screen{
			Size:5.5,
			ResX:1920,
			ResY:1080,
		},
		Battery: Battery{
			2910,
		},
		HasTouchID:true,
	}
	jsonData,_ := json.Marshal(raw)
	return jsonData
}

func main()  {
	jsonData := getJsonDate()
	fmt.Printf(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData,&screenAndTouch) 

	fmt.Printf("%+v\n", screenAndTouch)

	BatteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, &BatteryAndTouch)

	fmt.Printf("%+v\n", BatteryAndTouch)
}
