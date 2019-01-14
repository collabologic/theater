package foundation

import (
	"time"
)

//Pulser は定期的な振る舞いをエレメントに実施させるstructです
type Pulsar struct {
	interval  time.Duration       // 1Tickのマイクロ秒数
	behaviors []Behavior          // 振る舞いを行うエレメントの配列
	Running   bool                // 動作中か否か
	BehaveRun func(bs []Behavior) // 振る舞い実行の処理
	stopCh    chan bool           // 停止用フラグ
}

//newPulsar は新しいPalsarを生成します
func NewPulsar(interval time.Duration) Pulsar {
	return Pulsar{
		interval:  interval,
		behaviors: make([]Behavior, 0),
		Running:   true,
		BehaveRun: AllRun,
	}
}

//RegistBehavior はPulsarに対して実行する処理を追加できます
func (p *Pulsar) RegistBehavior(b Behavior) {
	p.behaviors = append(p.behaviors, b)
}

//AllRun は標準のBahaveRun。全て実行します
func AllRun(bs []Behavior) {
	for _, b := range bs {
		b.Behave()
	}
}

//Stop はPulsarの動作を停止します
func (p *Pulsar) Stop() {
	p.stopCh <- true
}

//Loop は定期的な振る舞いの呼び出しを実行します
func (p *Pulsar) Loop() {
	ticker := time.NewTicker(p.interval * time.Microsecond)
	p.stopCh = make(chan bool)
	go func(stopCh chan bool) {
	loop:
		for p.Running {
			select {
			case <-ticker.C:
				p.BehaveRun(p.behaviors)
			case <-stopCh:
				p.Running = false
				break loop
			}
		}
		close(stopCh)
	}(p.stopCh)
}
