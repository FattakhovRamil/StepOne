package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	maxAttempts     = 12      // максимальное количество попыток
	maxTotalTimeout = 10 * 60 // максимальное время выполнения попыток (10 минут)
)

func main() {
	filesURL := "http://localhost:3001/api/files" // получаем массив имен
	files, err := getFilesName(filesURL)
	if err != nil {
		fmt.Println("Ошибка при получении списка файлов:", err)
		return
	}

	createDirectories()
	successfulFile, notSuccessfulFile := createLogFiles()

	downloadFiles(files, filesURL, successfulFile, notSuccessfulFile)
}

func downloadFiles(names []string, filesURL string, successfulFile, notSuccessfulFile *os.File) {
	folderPath := "downloaded_files"

	for _, filename := range names {
		if !checkMP3File(filename) {
			fmt.Printf("Файл %s не является mp3 файлом\n", filename)
			fmt.Fprintf(notSuccessfulFile, "%s\n", filename)
			continue
		}

		startTime := time.Now()
		attempt := 0

		for time.Since(startTime).Seconds() < maxTotalTimeout {
			attempt++
			if attempt > maxAttempts {
				fmt.Printf("Превышено максимальное количество попыток (%d) для файла %s\n", maxAttempts, filename)
				fmt.Fprintf(notSuccessfulFile, "%s\n", filename)
				break
			}

			fileURL := filesURL + "/" + filename
			fmt.Println("Попытка скачивания файла:", fileURL)

			resp, err := http.Get(fileURL)
			if err != nil {
				fmt.Printf("Ошибка при получении файла %s: %v\n", filename, err)
				time.Sleep(5 * time.Second)
				continue
			}

			defer resp.Body.Close()
			filePath := filepath.Join(folderPath, filename) // Путь к файлу
			file, err := os.Create(filePath)

			if err != nil {
				fmt.Printf("Ошибка при создании файла %s: %v\n", filename, err)
				time.Sleep(5 * time.Second)
				continue
			}

			defer file.Close()

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				fmt.Printf("Ошибка при записи файла %s: %v\n", filename, err)
				time.Sleep(5 * time.Second)
			}
			fmt.Printf("Файл %s успешно скачан и сохранен в %s\n", filename, filePath)
			fmt.Fprintf(successfulFile, "%s\n", filename)
			break // Выход из цикла, если файл успешно скачан и сохранен
		}

		if time.Since(startTime).Seconds() >= maxTotalTimeout {
			fmt.Printf("Превышено максимальное время выполнения попыток (%d секунд) для файла %s\n", maxTotalTimeout, filename)
			fmt.Fprintf(notSuccessfulFile, "%s\n", filename)
		}
	}
}

func createDirectories() {
	folderPath := "downloaded_files"
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		fmt.Println("Ошибка при создании папки:", err)
		return
	}
}

func createLogFiles() (*os.File, *os.File) {
	successfulFile, err := os.Create("successful.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла successful.txt:", err)
		return nil, nil
	}

	notSuccessfulFile, err := os.Create("not-successful.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла not-successful.txt:", err)
		return nil, nil
	}

	return successfulFile, notSuccessfulFile
}

func checkMP3File(filename string) bool {
	return strings.HasSuffix(filename, ".mp3")
}

// Структура для представления возвращаемого JSON-объекта
type FileList struct {
	Files []string `json:"files"`
}

func getFilesName(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fileList FileList
	err = json.NewDecoder(resp.Body).Decode(&fileList)
	if err != nil {
		return nil, err
	}

	return fileList.Files, nil
}
