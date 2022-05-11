package spin

import (
	"testing"
	"time"
)

func TestFrame1(t *testing.T) {
	s := New("Testing Frame1")
	s.Set(Frame1)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame2(t *testing.T) {
	s := New("Testing Frame2")
	s.Set(Frame2)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame3(t *testing.T) {
	s := New("Testing Frame3")
	s.Set(Frame3)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame4(t *testing.T) {
	s := New("Testing Frame4")
	s.Set(Frame4)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame5(t *testing.T) {
	s := New("Testing Frame5")
	s.Set(Frame5)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame6(t *testing.T) {
	s := New("Testing Frame6")
	s.Set(Frame6)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame7(t *testing.T) {
	s := New("Testing Frame7")
	s.Set(Frame7)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame8(t *testing.T) {
	s := New("Testing Frame8")
	s.Set(Frame8)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame9(t *testing.T) {
	s := New("Testing Frame9")
	s.Set(Frame9)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame10(t *testing.T) {
	s := New("Testing Frame10")
	s.Set(Frame10)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame11(t *testing.T) {
	s := New("Testing Frame11")
	s.Set(Frame11)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame12(t *testing.T) {
	s := New("Testing Frame12")
	s.Set(Frame12)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame13(t *testing.T) {
	s := New("Testing Frame13")
	s.Set(Frame13)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame14(t *testing.T) {
	s := New("Testing Frame14")
	s.Set(Frame14)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame15(t *testing.T) {
	s := New("Testing Frame15")
	s.Set(Frame15)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame16(t *testing.T) {
	s := New("Testing Frame16")
	s.Set(Frame16)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame17(t *testing.T) {
	s := New("Testing Frame17")
	s.Set(Frame17)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame18(t *testing.T) {
	s := New("Testing Frame18")
	s.Set(Frame18)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame19(t *testing.T) {
	s := New("Testing Frame19")
	s.Set(Frame19)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame20(t *testing.T) {
	s := New("Testing Frame20")
	s.Set(Frame20)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame21(t *testing.T) {
	s := New("Testing Frame21")
	s.Set(Frame21)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFrame22(t *testing.T) {
	s := New("Testing Frame22")
	s.Set(Frame22)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestRed(t *testing.T) {
	s := New("Testing Red", WithColor(Red))
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestGreen(t *testing.T) {
	s := New("Testing Green", WithColor(Green))
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestYellow(t *testing.T) {
	s := New("Testing Yellow", WithColor(Yellow))
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestBlue(t *testing.T) {
	s := New("Testing Blue", WithColor(Blue))
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestFuchsia(t *testing.T) {
	s := New("Testing Fuchsia", WithColor(Fuchsia))
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestSkyBlue(t *testing.T) {
	s := New("Testing SkyBlue", WithColor(SkyBlue))
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestTimePerFrame(t *testing.T) {
	s := New("Testing TimePerFrame", WithTimePerFrame(time.Second))
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}
