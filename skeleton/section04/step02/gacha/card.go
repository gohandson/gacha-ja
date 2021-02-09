package gacha

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityXR Rarity = "XR"
)

func (r Rarity) String() string {
	return string(r)
}

// TODO: フィールドをエクスポートする
type Card struct {
	rarity Rarity // レア度
	name   string // 名前
}

func (c *Card) String() string {
	return c.rarity.String() + ":" + c.name
}
