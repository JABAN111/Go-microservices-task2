package l

import (
	"context"

	generator "github.com/dustinkirkland/golang-petname"
	"yadro.com/course/internal/logger"
)

var log = logger.GetInstance()

type Petnamer interface {
	SingleGenerator
	StreamGenerator
}

type SingleGenerator interface {
	GenerateName(words int, separator string) string
}

type StreamGenerator interface {
	GenerateMany(ctx context.Context, words int, separator string, names int) <-chan string
}

type petname struct{}

func NewNamer() Petnamer {
	return &petname{}
}

func (p *petname) GenerateName(words int, separator string) string {
	return generator.Generate(words, separator)
}

func (p *petname) GenerateMany(ctx context.Context, words int, separator string, names int) <-chan string {
	nameChan := make(chan string)

	go func() {
		defer close(nameChan)
		for range names {
			select {
			case <-ctx.Done():
				return
			case nameChan <- p.GenerateName(words, separator):
			}
		}
		log.Debug("End of GenerateMany")
	}()

	return nameChan
}
