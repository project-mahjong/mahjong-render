package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	model "github.com/project-mahjong/mahjong-backend/core"
	"github.com/project-mahjong/mahjong-render/text"
	"log"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)
	dataString, _, err := cin.ReadLine()
	if err != nil || len(dataString) == 0 {
		dataString = []byte(`{"Wall":["6z","2m","7z","4p","1z","7z","5z","6s","9m","5z","2p","6z","3s","4p","2m","6s","9p","6m","4p","4s","2s","8m","8p","8s","1p","5p","9s","6s","5s","6p","3m","4m","2p","0m","2z","6p","6p","2m","2z","1m","5p","7s","1z","2s","8p","0p","8m","9s","4m","9s","4s","8p","7s","9p","3z","6p","2s","9p","3m","7s","7p","6z","9m","4z","8m","6m","5z","7m","3z","7z","7p","6m","1s","3s","1p","1s","1s","0s","5z","7m","1m","3m","5s","4p","5m","2z","2z","1m","2p","3m","5m","8s","9p","1p","5p","2s","1z","6m","3z","7m","4m","4m","4z","8m","4z","1z","1s","8s","3p","3s","2m","7p","9m","5m","8p","1p","2p","7z","7s","3p","9s","3s","3z","6z","7p","5s","4s","3p","8s","4s","6s","3p","7m","1m","4z","9m"],"MD5":"cb48696c7e96d599c37c880f708cffc8","DoraIndicatorCount":1,"ReplacementTileCount":0,"WallCount":136,"Player":[{"HandTile":["2m","7m","9m","2p","3p","4p","7p","8s","9s","1z","4z","6z","7z"],"NowTile":"","DiscardTile":["7z","5z","6s","9m","5z","2p","6z","3s","7s","2s","7p","8m","3z","1s","1s","1m","5m","2p","9p","1z","4m"],"ReadHand":0,"Riichi":-1,"Group":null},{"HandTile":["1m","2m","5m","6m","4p","9p","3s","3s","4s","5s","6s","1z","7z"],"NowTile":"","DiscardTile":["4p","4s","2s","8m","8p","8s","1p","5p","9p","9p","6z","6m","7z","3s","0s","3m","2z","3m","1p","6m","4m"],"ReadHand":0,"Riichi":-1,"Group":null},{"HandTile":["2m","3m","6p","8p","1s","4s","5s","6s","6s","7s","9s","3z","4z"],"NowTile":"","DiscardTile":["4m","2p","0m","2z","6p","6p","2m","2z","3z","3m","9m","5z","7p","1p","5z","5s","2z","5m","5p","3z","4z"],"ReadHand":0,"Riichi":-1,"Group":null},{"HandTile":["1m","9m","1p","3p","3p","3p","5p","7p","2s","7s","8s","1z","6z"],"NowTile":"","DiscardTile":["8p","0p","8m","9s","4m","9s","4s","8p","6p","7s","4z","7m","6m","1s","7m","4p","1m","8s","2s","7m","8m"],"ReadHand":0,"Riichi":-1,"Group":null}]}`)
		err = nil
	}
	if err != nil {
		log.Panicln("unable to read stdin")
	}
	data := &model.TitleModel{}
	err = json.Unmarshal(dataString, data)
	if err != nil {
		log.Panicln("json error")
	}
	txt, err := text.RenderMahjong(data, 0)
	if err != nil {
		log.Panicln("render error")
	}
	fmt.Println(txt)
	/*file,err:=os.OpenFile("out.png",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0644)
	if err!=nil {
		log.Panicln("open file error")
	}
	err=png.Encode(file,img)
	if err!=nil {
		log.Panicln("write file error")
	}*/
}
