package fighters

type BaseFighter struct {
	Life int // 0..100
}

func (bf BaseFighter) IsAlive() bool {
	return bf.Life > 0
}

func (bf BaseFighter) GetLife() int {
	return bf.Life
}
