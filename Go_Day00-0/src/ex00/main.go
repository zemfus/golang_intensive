package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Numbers interface {
	ReadNumbers()
	Mean()
	Median()
	Mode()
	Sd()
}

type numbers struct {
	mass []int32
}

func (a *numbers) Median() {
	mNumber := len(a.mass) / 2

	if len(a.mass)%2 == 1 {
		fmt.Printf("Median: %.2f\n", float64(a.mass[mNumber]))
	} else {
		fmt.Printf("Median: %.2f\n", (float64(a.mass[mNumber-1])+float64(a.mass[mNumber]))/2)
	}
}

func (a *numbers) Mode() {
	countMap := make(map[int32]int)
	var mode int32
	for _, value := range a.mass {
		countMap[value] += 1
	}

	max := 0

	for _, key := range a.mass {
		freq := countMap[key]
		if freq > max {
			mode = key
			max = freq
		}
		if freq == max && key < mode {
			mode = key
		}
	}
	fmt.Printf("Mode: %.2f\n", float64(mode))
}

func (a *numbers) Sd() {
	var total float64
	var sd float64

	for _, v := range a.mass {
		total += float64(v)
	}
	total /= float64(len(a.mass))
	for _, numb := range a.mass {
		sd += math.Pow(float64(numb)-total, 2)
	}
	sd = math.Sqrt(sd / float64(len(a.mass)))

	fmt.Printf("SD: %.2f\n", sd)
}

func (a *numbers) Mean() {
	var total int32

	for _, v := range a.mass {
		total += v
	}
	fmt.Printf("Mean: %.2f\n", float64(total)/float64(len(a.mass)))
}

func (a *numbers) ReadNumbers() {
	var buff int32
	for {
		if _, err := fmt.Scan(&buff); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Print(err)
				os.Exit(1)
			}
			break
		}
		a.mass = append(a.mass, buff)
	}

}

func main() {

	var number numbers
	var ans Numbers
	ans = &number

	ans.ReadNumbers()

	data := map[string]int{
		"mean":   0,
		"median": 0,
		"mode":   0,
		"sd":     0,
	}

	if len(os.Args) > 1 {
		for _, param := range os.Args[1:] {
			if _, ok := data[strings.ToLower(param)]; ok {
				data[strings.ToLower(param)] = 1
			} else {
				fmt.Printf("Wrong argument: \"%s\"", param)
				return
			}
		}
	} else {
		data["mean"] = 1
		data["median"] = 1
		data["mode"] = 1
		data["sd"] = 1
	}

	sort.Slice(number.mass, func(i, j int) bool { return number.mass[i] < number.mass[j] })
	
	if i, _ := data["mean"]; i == 1 {
		ans.Mean()
	}
	if i, _ := data["median"]; i == 1 {
		ans.Median()
	}
	if i, _ := data["mode"]; i == 1 {
		ans.Mode()
	}
	if i, _ := data["sd"]; i == 1 {
		ans.Sd()
	}

}
