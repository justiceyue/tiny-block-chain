package transaction

type Transaction struct {
	// 交易id
	ID []byte
	// 交易输入序列
	Vin []*TxInput
	// 交易输出序列
	Vout []*TxOutput
}

type TxInput struct {
	// 交易的id
	TXid []byte
	// 标识使用该笔交易中哪个输出作为本次输入的代币
	Vout int64
	//	解锁脚本
	ScriptSig string
}

type TxOutput struct {
	// 输出的比特币数量
	Value int64
	// 锁定脚本
	ScriptPubKey string
}
