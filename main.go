package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

func NewWorld(height, width int) *World {
	// Создаём тип World с количеством слайсов hight (количество строк)
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width) // Создаём новый слайс в каждой строке
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}
}

var brownSquare = "\xF0\x9F\x9F\xAB"
var greenSquare = "\xF0\x9F\x9F\xA9"

//	type Color struct{
//		color string
//	}
//
//	func (w *World) String() [][]string {
//		array := make([][]string, w.Height)
//		for n, _ := range array {
//			array[n] = make([]string, w.Width)
//		}
//		for i := 0; i < len(w.Cells); i++ {
//			for k := 0; k < len(w.Cells[i]); k++ {
//				if w.Cells[i][k] {
//					array[i][k] = greenSquare // Живая клетка
//				} else {
//					array[i][k] = brownSquare // Мертвая клетка
//				}
//			}
//		}
//		return array
//	}
func (w *World) String() string {
	res := ""
	for i := 0; i < len(w.Cells); i++ {
		for k := 0; k < len(w.Cells[i]); k++ {
			if w.Cells[i][k] {
				res += greenSquare // Живая клетка
			} else {
				res += brownSquare // Мертвая клетка
			}
		}
		res += "\n"
	}
	return res
}

type Position struct {
	X int
	Y int
}

func (w *World) Neighbors(x, y int) int {
	counter := 0
	pos := []Position{{x - 1, y}, {x + 1, y}, {x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x - 1, y + 1}, {x, y - 1}, {x + 1, y + 1}}
	for _, p := range pos {
		if (0 <= p.Y && p.Y < w.Height) && (0 <= p.X && p.X < w.Width) {
			if w.Cells[p.Y][p.X] {
				counter++
			}
		}
	}
	return counter
}
func (w *World) Next(x, y int) bool {
	n := w.Neighbors(x, y)       // Получим количество живых соседей
	alive := w.Cells[y][x]       // Текущее состояние клетки
	if n < 4 && n > 1 && alive { // Если соседей двое или трое, а клетка жива,
		return true // то следующее её состояние — жива
	}
	if n == 3 && !alive { // Если клетка мертва, но у неё трое соседей,
		return true // клетка оживает
	}

	return false // В любых других случаях — клетка мертва
}

func (w *World) LoadState(filename string) error {
	file, _ := os.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	width := len(lines[0])
	height := len(lines)
	for i := range lines {
		if len(lines[i]) != width {
			return fmt.Errorf("error")
		}
	}
	w.Height = height
	w.Width = width
	return nil
}
func (w *World) SaveState(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] {
				_, err := file.WriteString("1")
				if err != nil {
					return err
				}
			} else {
				_, err := file.WriteString("0")
				if err != nil {
					return err
				}
			}
		}
		if i != w.Height-1 {
			_, err := file.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func NextState(oldWorld, newWorld *World) {
	// Переберём все клетки, чтобы понять, в каком они состоянии
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// Для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.Next(j, i)
		}
	}
}
func (w *World) Seed() {
	// Снова переберём все клетки
	for _, row := range w.Cells {
		for i := range row {
			//rand.Intn(10) возвращает случайное число из диапазона	от 0 до 9
			if rand.Intn(2) == 1 {
				row[i] = true
			}
		}
	}
}
func PrintField(field [][]bool) {

	for _, row := range field {
		for _, cell := range row {
			if cell {
				color.Set(color.FgGreen)
				fmt.Print("█") // Живая клетка
			} else {
				color.Set(color.FgRed) // Сброс цвета
				fmt.Print("+")         // Мертвая клетка
			}
		}
		color.Unset() // Сброс цвета после строки
		fmt.Println() // Переход на новую строку
	}
}
func main() {
	// Зададим размеры сетки
	height := 10
	width := 10
	// Объект для хранения текущего состояния сетки
	currentWorld := NewWorld(height, width)
	// Объект для хранения следующего состояния сетки
	nextWorld := NewWorld(height, width)
	// Установим начальное состояние
	currentWorld.Seed()
	for { // Цикл для вывода каждого состояния
		// Выведем текущее состояние на экран
		fmt.Println(currentWorld.String())
		// Рассчитываем с	ледующее состояние
		NextState(currentWorld, nextWorld)
		// Изменяем текущее состояние
		currentWorld = nextWorld
		// Делаем паузу
		time.Sleep(500 * time.Millisecond)
		// Специальная последовательность для очистки экрана после каждого шага
		fmt.Print("\033[H\033[2J")
	}
}
