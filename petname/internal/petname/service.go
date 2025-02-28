package l

import (
	"context"
	"fmt"

	generator "github.com/dustinkirkland/golang-petname"
	"yadro.com/course/internal/logger"
)

type InvalidArgument struct {
	Err error
}

func (r *InvalidArgument) Error() string {
	return fmt.Sprintf("Error has occured: %v", r.Err)
}

var log = logger.GetInstance()

type Petnamer interface {
	SingleGenerator
	StreamGenerator
}

type SingleGenerator interface {
	GenerateName(words int, separator string) (string, error)
}

type StreamGenerator interface {
	GenerateMany(ctx context.Context, words int, separator string, names int) (<-chan string, error)
}

type petname struct{}

func NewNamer() Petnamer {
	return &petname{}
}

func (p *petname) GenerateName(words int, separator string) (string, error) {
	if !isValid(words) {
		return "", &InvalidArgument{Err: fmt.Errorf("all parameters must be positive")}
	}
	return generator.Generate(words, separator), nil
}

func (p *petname) GenerateMany(ctx context.Context, words int, separator string, names int) (<-chan string, error) {
	nameChan := make(chan string)
	if !isValid(words, names) {
		return nil, &InvalidArgument{Err: fmt.Errorf("all parameters must be positive")}
	}

	go func() {
		defer close(nameChan)
		for range names {
			select {
			case <-ctx.Done():
				return
			case nameChan <- generator.Generate(words, separator):
			}
		}
	}()
	log.Debug("End of GenerateMany")
	return nameChan, nil
}

func isValid(args ...int) bool {
	for _, val := range args {
		if val < 1 {
			return false
		}
	}
	return true
}
