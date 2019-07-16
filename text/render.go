package text

import (
	"bytes"
	model "github.com/project-mahjong/mahjong-backend/core"
	"log"
	"text/template"
)

var tileName = map[string]string{"1m": "一万", "2m": "二万", "3m": "三万", "4m": "四万", "5m": "五万", "6m": "六万", "7m": "七万", "8m": "八万", "9m": "九万",
	"1p": "一筒", "2p": "二筒", "3p": "三筒", "4p": "四筒", "5p": "五筒", "6p": "六筒", "7p": "七筒", "8p": "八筒", "9p:": "九筒",
	"1s": "一条", "2s": "二条", "3s": "三条", "4s": "四条", "5s": "五条", "6s": "六条", "7s": "七条", "8s": "八条", "9s": "九条",
	"1z": "东", "2z": "南", "3z": "西", "4z": "北", "5z": "白", "6z": "发", "7z": "中", "0m": "红五万", "0p": "红五筒", "0s": "红五条"}

var tileNameUnicode = map[string]string{"": "🀫", "1m": "🀇", "2m": "🀈", "3m": "🀉", "4m": "🀊", "5m": "🀋", "6m": "🀌", "7m": "🀍", "8m": "🀎", "9m": "🀏",
	"1p": "🀙", "2p": "🀚", "3p": "🀛", "4p": "🀜", "5p": "🀝", "6p": "🀞", "7p": "🀟", "8p": "🀠", "9p": "🀡",
	"1s": "🀐", "2s": "🀑", "3s": "🀒", "4s": "🀓", "5s": "🀔", "6s": "🀕", "7s": "🀖", "8s": "🀗", "9s": "🀘",
	"1z": "🀀", "2z": "🀁", "3z": "🀂", "4z": "🀃", "5z": "🀆", "6z": "🀅", "7z": "🀄", "0m": "红🀋", "0p": "红🀝", "0s": "红🀔"}

func RenderMahjong(title *model.TitleModel, renderType int) (string, error) {
	vars := make(map[string]interface{})
	feng := []string{"东", "南", "西", "北"}
	vars["feng"] = feng
	dora := make([]string, 5)
	for i := 0; i < 5; i++ {
		if i < title.DoraIndicatorCount {
			dora[i] = tileNameUnicode[string(title.Wall[124+i*2])]
		} else {
			dora[i] = tileNameUnicode[""]
		}
	}
	vars["dora"] = dora
	discard := make([][]string, 0)
	for i := 0; i < 4; i++ {
		d := make([]string, 0)
		for _, j := range title.Player[i].DiscardTile {
			d = append(d, tileNameUnicode[string(j)])
		}
		discard = append(discard, d)
	}
	vars["discard"] = discard
	hand := make([][]string, 0)
	group := make([][][]string, 0)
	for i := 0; i < 4; i++ {
		h := make([]string, 0)
		for _, j := range title.Player[i].HandTile {
			h = append(h, tileNameUnicode[string(j)])
		}
		hand = append(hand, h)
		gr := make([][]string, 0)
		for _, j := range title.Player[i].Groups {
			g := make([]string, 0)
			for i, k := range j.Tiles {
				if j.Type == 2 && (i == 1 || i == 2) {
					g = append(g, tileNameUnicode[""])
				} else {
					g = append(g, tileNameUnicode[string(k)])
				}
			}
			gr = append(gr, g)
		}
		group = append(group, gr)
	}
	vars["hand"] = hand
	vars["group"] = group
	tpl := template.New("")
	tpl, err := tpl.Parse(`{{$discard:=.discard}}{{$hand:=.hand}}{{$group:=.group}}
宝牌:{{range .dora}} {{.}}{{end}}
{{range $p,$F := .feng}}{{$F}}:
  牌河:{{range index $discard $p}} {{.}}{{end}}
  手牌:{{range index $hand $p}} {{.}}{{end}}{{range $i:=index $group $p}}   {{range $i}} {{.}}{{end}}{{end}}
{{end}}
`)
	if err != nil {
		log.Panicln(err)
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, vars)
	if err != nil {
		log.Panicln(err)
	}
	return buf.String(), nil
}
