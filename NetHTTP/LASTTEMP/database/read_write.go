package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"temp/entity"
)

/*
	[
		[ApiBinanceResponse, ApiBinanceResponse, ApiBinanceResponse],
		[ApiBinanceResponse, ApiBinanceResponse, ApiBinanceResponse],
		[ApiBinanceResponse, ApiBinanceResponse, ApiBinanceResponse],
	]
*/

func (db *DataBase) ReadAllData() ([][]entity.ApiBinanceResponse, error) {
	var data [][]entity.ApiBinanceResponse

	file, err := os.OpenFile(db.name, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла %s: %v", db.name, err)
	}
	defer file.Close()

	raw_data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %v", db.name, err)
	}

	if len(raw_data) == 0 {
		return make([][]entity.ApiBinanceResponse, 0), nil
	}

	err = json.Unmarshal(raw_data, &data)
	if err != nil {
		return nil, fmt.Errorf("ошибка json.Unmarshal(raw_data, &data): %v", err)
	}

	return data, nil
}

func (db *DataBase) InsertData(data []entity.ApiBinanceResponse) error {

	all_data, err := db.ReadAllData()
	if err != nil {
		return fmt.Errorf("ошибка db.ReadAllData: %v", err)
	}

	new_data := append(all_data, data)

	file, err := os.OpenFile(db.name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла %s: %v", db.name, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // добавили отступы
	err = encoder.Encode(new_data)
	if err != nil {
		return fmt.Errorf("ошибка записи в файл %s: %v", db.name, err)
	}
	return nil
}
