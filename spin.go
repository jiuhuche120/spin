package spin

import (
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
	Box1    = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	Box2    = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	Box3    = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	Box4    = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	Box5    = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	Box6    = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	Box7    = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	Spin1   = `|/-\`
	Spin2   = `◴◷◶◵`
	Spin3   = `◰◳◲◱`
	Spin4   = `◐◓◑◒`
	Spin5   = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
	Spin6   = `▌▄▐▀`
	Spin7   = `╫╪`
	Spin8   = `■□▪▫`
	Spin9   = `←↑→↓`
	Spin10  = `⦾⦿`
	Spin11  = `⌜⌝⌟⌞`
	Spin12  = `┤┘┴└├┌┬┐`
	Spin13  = `⇑⇗⇒⇘⇓⇙⇐⇖`
	Spin14  = `☰☱☳☷☶☴`
	Spin15  = `䷀䷪䷡䷊䷒䷗䷁䷖䷓䷋䷠䷫`
	Default = Box1
)

type Spinner struct {
	frames []rune
	pos    int
	active bool
	text   string
	tpf    time.Duration
	writer io.Writer
}
type Option func(s *Spinner)

func WithFrames(frames string) Option {
	return func(s *Spinner) {
		s.Set(frames)
	}
}

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

// New creates a Spinner object with the provided
// text. By default, the Default spinner frames are
// used, and new frames are rendered every 100 milliseconds.
// Options can be provided to override these default
// settings.
func New(text string, opts ...Option) *Spinner {
	s := &Spinner{
		frames: []rune(Default),
		text:   ClearLine + text,
		tpf:    DefaultTime,
		writer: os.Stdout,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Start shows the spinner.
func (s *Spinner) Start() {
	if s.active == true {
		return
	}
	s.active = true
	go func() {
		t := time.Now()
		for s.active == true {
			fmt.Fprintf(s.writer, s.text+" %s%s", s.next(), " ("+strconv.FormatFloat(float64(time.Now().UnixNano()-t.UnixNano())/1e9, 'f', 3, 64)+"s)")
			time.Sleep(s.tpf)
		}
		fmt.Fprintf(s.writer, s.text+"  %s", " ("+strconv.FormatFloat(float64(time.Now().UnixNano()-t.UnixNano())/1e9, 'f', 3, 64)+"s)")
	}()
}

// Stop hides the spinner.
func (s *Spinner) Stop() bool {
	if s.active == true {
		s.active = false
		return true
	}
	return false
}

// Set frames to the given string which must not use spaces.
func (s *Spinner) Set(frames string) {
	s.frames = []rune(frames)
}

func (s *Spinner) next() string {
	r := s.frames[s.pos%len(s.frames)]
	s.pos++
	return string(r)
}
