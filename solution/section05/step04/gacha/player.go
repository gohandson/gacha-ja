package gacha

import "fmt"

type Player struct {
	tickets int // ガチャ券の枚数
	coin    int // コイン
}

func NewPlayer(tickets, coin int) *Player {
	return &Player{tickets: tickets, coin: coin}
}

// プレイヤーが行えるガチャの回数
func (p *Player) DrawableNum() int {
	// ガチャ券は1枚で1回、コインは10枚で1回ガチャが行える
	return p.tickets + p.coin/10
}

func (p *Player) draw(n int) {

	if p.DrawableNum() < n {
		fmt.Println("ガチャ券またはコインが不足しています")
		return
	}

	// ガチャ券から優先的に使う
	if p.tickets > n {
		p.tickets -= n
		return
	}

	p.tickets = 0
	p.coin -= n * 10 // 1回あたり10枚消費する
}
