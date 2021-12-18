package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)



const walletFile = "D:/Studymaterial/CS/golang/tmp/wallets.data"

type Wallets struct {
	Wallets map[string]*Wallet
}
func(ws *Wallets) LoadFile() error{
	if _,err := os.Stat(walletFile);os.IsNotExist(err){
		return err
	}
	var wallets Wallets
	fileContent,err := ioutil.ReadFile(walletFile)
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder((bytes.NewReader(fileContent)))
	err = decoder.Decode(&wallets)
	if err!= nil{
		return err
	}
	ws.Wallets = wallets.Wallets
	return nil

	

}

func (ws *Wallets) SaveFile() {
	var content bytes.Buffer
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil{
		log.Panic(err)
	}
	err = ioutil.WriteFile(walletFile,content.Bytes(),0644)
	if err != nil{
		log.Panic(err)
	}
}