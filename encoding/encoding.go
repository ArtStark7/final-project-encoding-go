package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose // данные для сериализации и десериализации
	FileInput     string                // имя файла, который нужно перекодировать
	FileOutput    string                // имя файла с результатом перекодирования
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose // данные для сериализации и десериализации
	FileInput     string                // имя файла, который нужно перекодировать
	FileOutput    string                // имя файла с результатом перекодирования
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	//Первым делом прочитали входящий файл (Input)
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	// Десериализировалии полученные JSON данные
	err = json.Unmarshal(jsonFile, &j.DockerCompose)
	if err != nil {
		return err
	}
	// Сериализировали данные в YAML
	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		return err
	}
	//os.Create(j.FileOutput): Эта функция пытается создать новый файл с именем, указанным в j.FileOutput. Если файл уже существует, его содержимое будет перезаписано.
	//- yamlFile: Переменная yamlFile будет содержать указатель на созданный файл.
	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}
	//- defer: Эта инструкция откладывает выполнение функции yamlFile.Close() до момента, когда текущая функция завершит свое выполнение. Это гарантирует, что файл будет закрыт, даже если в процессе выполнения произойдет ошибка.
	//- Закрытие файла освобождает ресурсы, связанные с ним, и корректно завершает работу с файлом.
	defer yamlFile.Close()

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
