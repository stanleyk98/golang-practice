package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Arafatk/glot"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("데이터 파일이 필요합니다.")
		return
	}
	// file 이 기 존재하는지 os.Stat()함수로 체크
	file := os.Args[1]
	_, err := os.Stat(file)

	if err != nil {
		fmt.Println("Can NOT stat", file)
		return
	}
	// file 열기
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("파일을 열수 없음", file)
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	xP := []float64{}
	yP := []float64{}

	for _, rec := range allRecords {
		//string을 float 타입으로 전환
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		yP = append(yP, y)
	}

	// 2차원 슬라이스 생성
	points := [][]float64{}
	points = append(points, xP)
	points = append(points, yP)
	fmt.Println(points)

	dimensions := 2
	persist := true
	debug := false

	plot, _ := glot.NewPlot(dimensions, persist, debug)
	plot.SetTitle("Glot 플러그인사용")
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")

	style := "원형"
	plot.AddPointGroup("Circle: ", style, points)
	plot.SavePlot("csv-output.png")
}
