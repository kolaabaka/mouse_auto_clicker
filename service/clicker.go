package click_service

import (
	"context"
	"time"

	"github.com/go-vgo/robotgo"
)

type ClickerConfig struct {
	Delimetr int
}

func (p *ClickerConfig) IncreaseTiming() {
	p.Delimetr += 1
}

func (p *ClickerConfig) ReduceTiming() {
	if p.Delimetr-1 > 0 {
		p.Delimetr -= 1
	}
}

func (p *ClickerConfig) ClickingStart(ctx context.Context) {
	if p.Delimetr <= 0 {
		p.Delimetr = 1
	}
	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		default:
			{
				time.Sleep(time.Millisecond * 100 * time.Duration(p.Delimetr))
				robotgo.Click()
			}
		}
	}
}
