package test

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Запустить программу и перенаправить ее вывод в файл
	f, err := os.Create("output.txt")
	if err != nil {
		t.Fatalf("Error creating output file: %v", err)
	}
	cmd := exec.Command("../../task.exe", "test_file2.txt")
	cmd.Stdout = f

	// Запустить программу
	if err := cmd.Run(); err != nil {
		t.Fatalf("Error running program: %v", err)
	}

	// Считать содержимое файла с выводом программы
	output, err := os.ReadFile("output.txt")
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	// Считать содержимое ожидаемого файла
	expectedOutput, err := os.ReadFile("expected_output.txt")
	if err != nil {
		t.Fatalf("Error reading expected output file: %v", err)
	}

	// Найти первую разницу между двумя файлами
	diffIndex := strings.IndexByte(string(output), byte(output[0])^byte(expectedOutput[0]))

	// Если разница найдена, вывести номер строки и столбца
	if diffIndex != -1 {
		line := strings.Count(string(output)[:diffIndex], "\n") + 1
		col := diffIndex - strings.LastIndex(string(output)[:diffIndex], "\n") + 1
		fmt.Printf("Difference found at line %d, column %d", line, col)
		t.Errorf("Expected output '%s', got '%s'", expectedOutput, output)
	} else {
		fmt.Println("No difference found")
	}
}
