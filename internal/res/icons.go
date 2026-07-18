package res

import (
	
	"fyne.io/fyne/v2"
	_ "embed"
)
//go:embed images/Logo.jpg
var resourceLogoJpgData []byte
var ResourceLogoJpg = &fyne.StaticResource{
	StaticName:    "images/Logo.jpg",
	StaticContent: resourceLogoJpgData,
}

//go:embed images/goldenStar.png
var resourceGoldenStarPngData []byte
var ResourceGoldenStarPng = &fyne.StaticResource{
	StaticName:    "images/goldenStar.png",
	StaticContent: resourceGoldenStarPngData,
}

//go:embed images/blackStar.png
var resourceBlackStarPngData []byte
var ResourceBlackStarPng = &fyne.StaticResource{
	StaticName:    "images/blackStar.png",
	StaticContent: resourceBlackStarPngData,
}

//go:embed images/onOffButton.png
var resourceOnOffButtonPngData []byte
var ResourceOnOffButtonPng = &fyne.StaticResource{
	StaticName:    "images/onOffButton.png",
	StaticContent: resourceOnOffButtonPngData,
}

//go:embed images/floralCircle.png
var resourceFloralCirclePngData []byte
var ResourceFloralCirclePng = &fyne.StaticResource{
	StaticName:    "images/floralCircle.png",
	StaticContent: resourceFloralCirclePngData,
}

//go:embed images/checkMark.png
var resourceCheckMarkPngData []byte
var ResourceCheckMarkPng = &fyne.StaticResource{
	StaticName:    "images/checkMark.png",
	StaticContent: resourceCheckMarkPngData,
}
