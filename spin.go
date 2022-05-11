package spin

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// ClearLine go to the beginning or the line and clear it
const ClearLine = "\r\033[K"
const DefaultTime = time.Millisecond * 100

// Spinner types
var (
	Frame1  = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	Frame2  = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	Frame3  = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	Frame4  = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	Frame5  = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	Frame6  = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	Frame7  = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	Frame8  = `|/-\`
	Frame9  = `◴◷◶◵`
	Frame10 = `◰◳◲◱`
	Frame11 = `◐◓◑◒`
	Frame12 = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
	Frame13 = `▌▄▐▀`
	Frame14 = `╫╪`
	Frame15 = `■□▪▫`
	Frame16 = `←↑→↓`
	Frame17 = `⦾⦿`
	Frame18 = `⌜⌝⌟⌞`
	Frame19 = `┤┘┴└├┌┬┐`
	Frame20 = `⇑⇗⇒⇘⇓⇙⇐⇖`
	Frame21 = `☰☱☳☷☶☴`
	Frame22 = `䷀䷪䷡䷊䷒䷗䷁䷖䷓䷋䷠䷫`
	OK      = `√`
	Default = Frame1
)

type Color string

// Spinner color
const (
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Fuchsia Color = "\033[35m"
	SkyBlue Color = "\033[36m"
	White   Color = "\033[37m"
)

type Spinner struct {
	frames []rune
	pos    int
	ctx    context.Context
	cancel context.CancelFunc
	text   string
	tpf    time.Duration
	color  Color
	writer io.Writer
}
type Option func(s *Spinner)

// WithTimePerFrame sets how long each frame shall
// be shown.
func WithTimePerFrame(d time.Duration) Option {
	return func(s *Spinner) {
		s.tpf = d
	}
}

// WithWriter sets the writer to use for spinner's text.
func WithWriter(w io.Writer) Option {
	return func(s *Spinner) {
		s.writer = w
	}
}

// WithColor sets the color of text
func WithColor(color Color) Option {
	return func(s *Spinner) {
		s.color = color
	}
}

// New creates a Spinner object with the provided
// text. By default, the Default spinner frames are
// used, and new frames are rendered every 100 milliseconds.
// Options can be provided to override these default
// settings.
func New(text string, opts ...Option) *Spinner {
	s := &Spinner{
		frames: []rune(Default),
		text:   text,
		tpf:    DefaultTime,
		writer: os.Stdout,
		color:  White,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Start shows the spinner.
func (s *Spinner) Start() {
	if s.ctx != nil || s.cancel != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.ctx = ctx
	s.cancel = cancel
	ticker := time.NewTicker(s.tpf)
	t := time.Now()
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				str := fmt.Sprintf("%s %s %s\n", OK, s.text, getTime(t))
				_, err := fmt.Fprintf(s.writer, "%s", ClearLine+getColor(str, s.color))
				if err != nil {
					fmt.Println(err)
					return
				}
				return
			case <-ticker.C:
				str := fmt.Sprintf("%s %s %s", s.next(), s.text, getTime(t))
				_, err := fmt.Fprintf(s.writer, "%s", ClearLine+getColor(str, s.color))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}()
}

// Stop hides the spinner.
func (s *Spinner) Stop() bool {
	if s.ctx == nil || s.cancel == nil {
		return false
	}
	s.cancel()
	time.Sleep(s.tpf)
	s.ctx = nil
	s.cancel = nil
	return true
}

// Set frames to the given string which must not use spaces.
func (s *Spinner) Set(frames string) {
	s.frames = []rune(frames)
}

// next give the next frame
func (s *Spinner) next() string {
	r := s.frames[s.pos%len(s.frames)]
	s.pos++
	return string(r)
}

// getTime get formatting time string
func getTime(t time.Time) string {
	return "(" + strconv.FormatFloat(float64(time.Now().UnixNano()-t.UnixNano())/1e9, 'f', 3, 64) + "s)"
}

// getColor str surround by color
func getColor(str string, color Color) string {
	return string(color) + str + "\033[m"
}
