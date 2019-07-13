package text

import (
	model "github.com/project-mahjong/mahjong-backend/core"
)

var tileName = map[string]string{"1m": "一万", "2m": "二万", "3m": "三万", "4m": "四万", "5m": "五万", "6m": "六万", "7m": "七万", "8m": "八万", "9m": "九万",
	"1p": "一筒", "2p": "二筒", "3p": "三筒", "4p": "四筒", "5p": "五筒", "6p": "六筒", "7p": "七筒", "8p": "八筒", "9p:": "九筒",
	"1s": "一条", "2s": "二条", "3s": "三条", "4s": "四条", "5s": "五条", "6s": "六条", "7s": "七条", "8s": "八条", "9s": "九条",
	"1z": "东", "2z": "南", "3z": "西", "4z": "北", "5z": "白", "6z": "发", "7z": "中", "0m": "红五万", "0p": "红五筒", "0s": "红五条"}

var tileNameUnicode = map[string]string{"1m": "🀇", "2m": "🀈", "3m": "🀉", "4m": "🀊", "5m": "🀋", "6m": "🀌", "7m": "🀍", "8m": "🀎", "9m": "🀏",
	"1p": "🀙", "2p": "🀚", "3p": "🀛", "4p": "🀜", "5p": "🀝", "6p": "🀞", "7p": "🀟", "8p": "🀠", "9p": "🀡",
	"1s": "🀐", "2s": "🀑", "3s": "🀒", "4s": "🀓", "5s": "🀔", "6s": "🀕", "7s": "🀖", "8s": "🀗", "9s": "🀘",
	"1z": "🀀", "2z": "🀁", "3z": "🀂", "4z": "🀃", "5z": "🀆", "6z": "🀅", "7z": "🀄", "0m": "红🀋", "0p": "红🀝", "0s": "红🀔"}

func RenderMahjong(title *model.TitleModel, renderType int) (string, error) {
	text := ""
	text += "牌河: \n"
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			text += "东"
		case 1:
			text += "南"
		case 2:
			text += "西"
		case 3:
			text += "北"
		}
		text += ":{"
		for _, j := range title.Player[i].DiscardTile {
			text += tileNameUnicode[string(j)] + ","
		}
		text += "}\n"
	}
	text += "手牌: \n"
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			text += "东"
		case 1:
			text += "南"
		case 2:
			text += "西"
		case 3:
			text += "北"
		}
		text += ":{"
		for _, j := range title.Player[i].HandTile {
			text += tileNameUnicode[string(j)] + ","
		}
		text += "}\n"
	}
	return text, nil
}
