// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided ‘as is’ and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package txpool

import (
	"encoding/hex"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/iotexproject/iotex-core/blockchain"
	"github.com/iotexproject/iotex-core/config"
	"github.com/iotexproject/iotex-core/proto"
	ta "github.com/iotexproject/iotex-core/test/testaddress"
)

const (
	testingConfigPath = "../config.yaml"
	testDBPath        = "db.test"
)

func decodeHash(in string) []byte {
	hash, _ := hex.DecodeString(in)
	return hash
}

func TestTxPool(t *testing.T) {
	defer os.Remove(testDBPath)
	assert := assert.New(t)

	config, err := config.LoadConfigWithPathWithoutValidation(testingConfigPath)
	assert.Nil(err)
	config.Chain.ChainDBPath = testDBPath
	// Disable block reward to make bookkeeping easier
	Gen.BlockReward = uint64(0)

	// Create a blockchain from scratch
	// bc := CreateBlockchain(Addrinfo["miner"].Address, &config.Config{Chain: config.Chain{ChainDBPath: testDBPath}})
	bc := CreateBlockchain(ta.Addrinfo["miner"].RawAddress, config, Gen)
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()

	//bc := mock_blockchain.NewMockBlockchain(ctrl)
	if assert.NotNil(bc) {
		t.Log("blockchain created")
	}
	defer bc.Close()

	tp := New(bc)
	cbTx := NewCoinbaseTx(ta.Addrinfo["miner"].RawAddress, 50, GenesisCoinbaseData)
	assert.NotNil(cbTx)
	if _, err := tp.ProcessTx(cbTx, true, false, 13245); assert.NotNil(err) {
		t.Logf("Coinbase Tx cannot be processed")
	}
	/*
		if true {
			pool := bc.UtxoPool()
			fmt.Println("utxoPool before tx1:", len(pool))
			for hash, _ := range pool {
				fmt.Printf("hash:%x\n", hash)
			}
			payees := []*Payee{}
			payees = append(payees, &Payee{Addrinfo["alfa"].Address, 10})
			payees = append(payees, &Payee{Addrinfo["bravo"].Address, 1})
			payees = append(payees, &Payee{Addrinfo["charlie"].Address, 1})
			payees = append(payees, &Payee{Addrinfo["delta"].Address, 1})
			payees = append(payees, &Payee{Addrinfo["echo"].Address, 1})
			payees = append(payees, &Payee{Addrinfo["foxtrot"].Address, 5})
			tx1 := bc.CreateTransaction(Addrinfo["miner"], 19, payees)
			fmt.Printf("tx1: %x\n", tx1.Hash())
			fmt.Println("version:", tx1.Version)
			fmt.Println("NumTxIn:", tx1.NumTxIn)
			fmt.Println("TxIn:")
			for idx, txIn := range tx1.TxIn {
				hash := ZeroHash32B
				copy(hash[:], txIn.TxHash)
				fmt.Printf("tx1.TxIn.%d %x %d %x\n", idx, hash, txIn.UnlockScriptSize, txIn.UnlockScript)
			}
			fmt.Println("NumTxOut:", tx1.NumTxOut)
			fmt.Println("TxOut:")
			for idx, txOut := range tx1.TxOut {
				fmt.Printf("tx1.TxOut.%d %d %x %d\n", idx, txOut.LockScriptSize, txOut.LockScript, txOut.Value)
			}
			blk := bc.MintNewBlock([]*Tx{tx1}, Addrinfo["miner"], "")
			fmt.Println(blk)
			err := bc.AddBlockCommit(blk)
			assert.Nil(err)
			fmt.Println("Add 1st block")
			bc.Reset()
			payees = nil
			payees = append(payees, &Payee{Addrinfo["bravo"].Address, 3})
			payees = append(payees, &Payee{Addrinfo["delta"].Address, 2})
			payees = append(payees, &Payee{Addrinfo["echo"].Address, 1})
			tx2 := bc.CreateTransaction(Addrinfo["alfa"], 6, payees)
			fmt.Printf("tx2: %x\n", tx2.Hash())
			fmt.Println("tx2.TxIn:", tx2.NumTxIn)
			for idx, txIn := range tx2.TxIn {
				hash := ZeroHash32B
				copy(hash[:], txIn.TxHash)
				fmt.Printf("tx2.TxIn.%d %x %d %x\n", idx, hash, txIn.UnlockScriptSize, txIn.UnlockScript)
			}
			fmt.Println("TxOut:")
			for idx, txOut := range tx2.TxOut {
				fmt.Printf("tx2.TxOut.%d %d %x %d\n", idx, txOut.LockScriptSize, txOut.LockScript, txOut.Value)
			}
			fmt.Println(tx2.TxIn)
			return
		} //*/
	txIn1_0 := &iproto.TxInputPb{
		TxHash:           decodeHash("9de6306b08158c423330f7a27243a1a5cbe39bfd764f07818437882d21241567"),
		OutIndex:         0,
		UnlockScriptSize: 98,
		UnlockScript:     decodeHash("40f9ea2b1357dde55519246a6ad82c466b9f2b988ff81a7c2fb114c932d44f322ba2edd178c2326739638b536e5f803977c24332b8f5b8ebc5f6683ff2bcaad90720b9b8d7316705dc4ff62bb323e610f3f5072abedc9834e999d6537f6681284ea2"),
	}
	txOut1_0 := NewTxOutput(10, 0)
	txOut1_0.LockScriptSize = 25
	txOut1_0.LockScript = decodeHash("65b014a97ce8e76ade9b3181c63432a62330a5ca83ab9ba1b1")
	txOut1_1 := NewTxOutput(1, 1)
	txOut1_1.LockScriptSize = 25
	txOut1_1.LockScript = decodeHash("65b014af33097c8fd571c6c1efc52b0a802514ea0fbb03a1b1")
	txOut1_2 := NewTxOutput(1, 2)
	txOut1_2.LockScriptSize = 25
	txOut1_2.LockScript = decodeHash("65b0140fb02223c1a78c3f1fb81a1572e8b07adb700bffa1b1")
	txOut1_3 := NewTxOutput(1, 3)
	txOut1_3.LockScriptSize = 25
	txOut1_3.LockScript = decodeHash("65b01443251ba4fd765a2cfa65256aabd64f98c5c00e40a1b1")
	txOut1_4 := NewTxOutput(1, 4)
	txOut1_4.LockScriptSize = 25
	txOut1_4.LockScript = decodeHash("65b01430f1db72a44136e8634121b6730c2b8ef094f1c9a1b1")
	txOut1_5 := NewTxOutput(5, 5)
	txOut1_5.LockScriptSize = 25
	txOut1_5.LockScript = decodeHash("65b014d94ee6c7205e85c3d97c557f08faf8ac41102806a1b1")
	txOut1_6 := NewTxOutput(9999999981, 6)
	txOut1_6.LockScriptSize = 25
	txOut1_6.LockScript = decodeHash("65b014d4f743a24d5386f8d1c2a648da7015f08800cd11a1b1")
	tx1 := &Tx{
		Version:  1,
		NumTxIn:  1,
		TxIn:     []*TxInput{txIn1_0},
		NumTxOut: 7,
		TxOut:    []*TxOutput{txOut1_0, txOut1_1, txOut1_2, txOut1_3, txOut1_4, txOut1_5, txOut1_6},
		LockTime: 0,
	}

	txIn2_0 := &iproto.TxInputPb{
		TxHash:           decodeHash("aeedd06eb44f08abbcc72a2293aff580f13662fa59cc1b0aa4a15ee7c118e4eb"),
		OutIndex:         0,
		UnlockScriptSize: 98,
		UnlockScript:     decodeHash("40535e20b5c5075fa80d6bff220aea755737e3787bfbc7122c0c45015f6c249fbca28d069dc028fad01fda2766ea90411aad38ce9a9de7c59a30e4bebc80b940002009d8c6fc6f5cb0a03df112da90486fad7cdece1501aaab658551f8afbe7f59ee"),
	}
	txOut2_0 := NewTxOutput(3, 0)
	txOut2_0.LockScriptSize = 25
	txOut2_0.LockScript = decodeHash("65b014af33097c8fd571c6c1efc52b0a802514ea0fbb03a1b1")
	txOut2_1 := NewTxOutput(2, 1)
	txOut2_1.LockScriptSize = 25
	txOut2_1.LockScript = decodeHash("65b01443251ba4fd765a2cfa65256aabd64f98c5c00e40a1b1")
	txOut2_2 := NewTxOutput(1, 2)
	txOut2_2.LockScriptSize = 25
	txOut2_2.LockScript = decodeHash("65b01430f1db72a44136e8634121b6730c2b8ef094f1c9a1b1")
	txOut2_3 := NewTxOutput(4, 2)
	txOut2_3.LockScriptSize = 25
	txOut2_3.LockScript = decodeHash("65b014a97ce8e76ade9b3181c63432a62330a5ca83ab9ba1b1")
	tx2 := &Tx{
		Version:  1,
		NumTxIn:  1,
		TxIn:     []*TxInput{txIn2_0},
		NumTxOut: 4,
		TxOut:    []*TxOutput{txOut2_0, txOut2_1, txOut2_2, txOut2_3},
		LockTime: 0,
	}

	t.Logf("tx1 hash: %x", tx1.Hash())
	t.Logf("tx2 hash: %x", tx2.Hash())
	descs, err := tp.ProcessTx(tx2, true, false, 12341234)
	assert.Nil(err)
	assert.Equal(0, len(descs))
	descs, err = tp.ProcessTx(tx1, true, false, 12341234)
	t.Log(tp.TxDescs())
	for hash, desc := range tp.TxDescs() {
		t.Logf("hash: %x desc: %v", hash, desc)
	}
	assert.Nil(err)
	assert.Equal(2, len(tp.TxDescs()))
}
