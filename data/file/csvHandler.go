package file

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// will append/add new record in existing records
func AppendRow(fileName string, rowData []string) error {
	// open file in create, append mode
	csvfile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicf("failed in appending row: %s", err)
		return err
	}
	// new file writer
	csvwriter := csv.NewWriter(csvfile)
	// append new record
	_ = csvwriter.Write(rowData)
	// flushing csvwriter
	csvwriter.Flush()
	csvfile.Close()
	return nil
}

// create a new file or overwrite existing file
func WriteMultiRow(fileName string, rowDataArr [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("failed in inserting file: %s", err)
		return err
	}
	// new file writer
	writer := csv.NewWriter(file)
	// Write all the records
	err = writer.WriteAll(rowDataArr)
	if err != nil {
		log.Println("Error while writing to the file ::", err)
		return err
	}
	//flushed writer
	writer.Flush()
	err = file.Close()
	if err != nil {
		log.Println("Error while closing the file ::", err)
	}
	return nil
}

// read all content of file
func Fetch(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// file scanner
	scanner := bufio.NewScanner(file)
	// split file data for new line
	scanner.Split(bufio.ScanLines)
	var rows [][]string
	for scanner.Scan() {
		rowStr := scanner.Text()
		// split row for columns
		cols := strings.Split(rowStr, ",")
		rows = append(rows, cols)
	}
	// file close
	file.Close()
	return rows, nil
}
