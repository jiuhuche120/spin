# spin

## Installation

To install spin, simply run

~~~bash
go get github.com/jiuhuche120/spin@latest
~~~

## Example usage

~~~go
package main

import (
	"time"

	"github.com/jiuhuche120/spin"
)

func main() {
	s := spin.New("Testing", spin.WithColor(spin.Red))
	s.Start()
	time.Sleep(time.Second*5)
	s.Stop()
}

~~~

## Func

~~~go
func New(text string, opts ...Option) *Spinner
~~~

create spinner with text by default options

~~~go
func WithTimePerFrame(d time.Duration) Option
~~~

the option of speed

~~~go
func WithColor(color Color) Option 
~~~

the option of text's color

~~~go
func (s *Spinner) Set(frames string)
~~~

set the spinner style

~~~go
func (s *Spinner) Start() 
~~~

start spinner

~~~go
func (s *Spinner) Stop() bool 
~~~

stop spinner
