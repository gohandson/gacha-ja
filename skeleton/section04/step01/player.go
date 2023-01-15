package main

import "fmt"

// TODO: playerに関する処理をここに移す
type player struct {
	tickets int // ガチャ券の枚数
	coin    int // コイン
}

// プレイヤーが行えるガチャの回数
func (p *player) drawableNum() int {
	// TODO: ガチャが行える回数を返す
	// ガチャ券は1枚で1回、コインは10枚で1回ガチャが行える
	return p.tickets + p.coin/10
}

func (p *player) draw(n int) {

	if p.drawableNum() < n {
		fmt.Println("ガチャ券またはコインが不足しています")
		return
	}

	// TODO: ガチャ券で足りる場合はガチャ券だけ使う
	// ガチャ券から優先的に使う
	if p.tickets > n {
		p.tickets -= n
		return
	}

	p.tickets = 0
	p.coin -= n * 10 // 1回あたり10枚消費する
}
