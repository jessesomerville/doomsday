package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/guptarohit/asciigraph"
	"golang.org/x/term"
)

const (
	timesFile    = "answer_times.log"
	wrongCntFile = "incorrect_answers_count"
)

func logStats(t time.Duration, wrong int) error {
	fmt.Println("Answered in " + t.Truncate(time.Second).String())
	p, err := cfgDir()
	if err != nil {
		return err
	}
	if err := appendTime(t, path.Join(p, timesFile)); err != nil {
		return fmt.Errorf("failed to add answer time: %v", err)
	}
	if err := addWrongCount(wrong, path.Join(p, wrongCntFile)); err != nil {
		return fmt.Errorf("failed to add wrong guess count: %v", err)
	}
	return nil
}

func cfgDir() (string, error) {
	uCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("could not get config directory: %v", err)
	}
	cfgDir := path.Join(uCfgDir, "doomsday")
	err = os.Mkdir(cfgDir, 0o750)
	if err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("could not get config directory: %v", err)
	}
	return cfgDir, nil
}

func appendTime(t time.Duration, path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	ts := fmt.Sprintf("%dms", t.Milliseconds())
	stat, err := f.Stat()
	if err != nil {
		return err
	}
	if stat.Size() > 0 {
		// Prepend newline unless this is the first entry.
		ts = "\n" + ts
	}
	if _, err := f.WriteString(ts); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func addWrongCount(cnt int, path string) error {
	allCnt, err := readWrongCount(path)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.WriteString(fmt.Sprintf("%d", allCnt+cnt))
	return err
}

func readWrongCount(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}
	cnt, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func displayStats() error {
	p, err := cfgDir()
	if err != nil {
		return err
	}
	wrongCnt, err := readWrongCount(path.Join(p, wrongCntFile))
	if err != nil {
		return err
	}
	timeData, err := os.ReadFile(path.Join(p, timesFile))
	if err != nil {
		return err
	}
	if len(timeData) == 0 {
		fmt.Println("No entries have been recorded")
		return nil
	}
	lines := strings.Split(string(timeData), "\n")
	times := make([]float64, 0, len(lines))
	statSum := 0
	var fastest time.Duration
	for i, l := range lines {
		ms, err := time.ParseDuration(l)
		if err != nil {
			return fmt.Errorf("failed to parse stat from %q: %v", p, err)
		}
		times = append(times, ms.Seconds())
		statSum += int(ms.Milliseconds())
		if i == 0 || ms < fastest {
			fastest = ms
		}
	}
	avg, _ := time.ParseDuration(fmt.Sprintf("%dms", statSum/len(times)))
	fmt.Printf("Wrong guesses: %d\n", wrongCnt)
	fmt.Println("Average time:  " + avg.String())
	fmt.Println("Fastest time:  " + fastest.String())
	showGraph(times)
	return nil
}

func showGraph(times []float64) {
	width := len(times)
	if width < 10 {
		return
	}

	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return
	}
	if width > w {
		width = w
	} else if width*4 < w {
		width *= 4
	} else if width*2 < w {
		width *= 2
	}
	offset := 3
	for _, t := range times {
		ts := fmt.Sprintf("%d", int(t))
		if len(ts) > offset {
			offset = len(ts)
		}
	}
	g := asciigraph.Plot(times,
		asciigraph.SeriesColors(asciigraph.Blue),
		asciigraph.Caption(fmt.Sprintf("%d answers / time", len(times))),
		asciigraph.Precision(0),
		asciigraph.Height(15),
		asciigraph.Width(width-offset-2),
		asciigraph.Offset(offset),
	)
	fmt.Printf("\n%s\n", g)
}
