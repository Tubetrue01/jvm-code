package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Instruction struct {
	Code     string `json:"code"`
	Command  string `json:"command"`
	Describe string `json:"describe"`
}

const filePath = "./file/code.json"

var instructions []Instruction
var instructionMap map[string]Instruction

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		code, _ := reader.ReadString('\n')
		code = strings.Replace(code, "\n", "", -1)

		if strings.Compare("q", code) == 0 {
			break
		}
		if instruction, ok := instructionMap[code]; ok {
			fmt.Printf("code : %s\t command: %s\t describe: %s\t\n", instruction.Code, instruction.Command, instruction.Describe)
		} else {
			fmt.Println("暂无该指令，请重新输入")
		}
	}
}

func init() {
	file, err := os.Open(filePath)

	if err != nil {
		log.Panic("文件打开失败")
	}

	defer file.Close()

	instructions = make([]Instruction, 202)
	err = json.NewDecoder(file).Decode(&instructions)

	if err != nil {
		log.Panic("转换 Json 数据失败")
	}
	instructionMap = make(map[string]Instruction)

	for _, instruction := range instructions {
		instructionMap[instruction.Command] = instruction
	}

}
