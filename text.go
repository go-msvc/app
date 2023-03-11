package app

import (
	"context"
)

type TextFunc func(ctx context.Context) string

func StaticText(s string) TextFunc {
	return func(ctx context.Context) string {
		return s
	}
}

// func P(tf TextFunc) App {
// 	return p{
// 		tf: tf,
// 	}
// }

// type p struct {
// 	tf TextFunc
// }

// func (p p) Render(scr Screen, ctx context.Context) {
// 	scr.P(p.tf(ctx))
// }
