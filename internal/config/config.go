package config

type Config struct {
    DBPath     string // Путь к базе данных
    APIKey     string // Ключ для API
    GameTimeout int   // Таймер игры
}
import "os"

func LoadConfig() Config {
    return Config{
        DBPath:     os.Getenv("DB_PATH"),      // Читаем путь из переменной окружения
        APIKey:     os.Getenv("API_KEY"),     // Ключ API
        GameTimeout: 300,                      // Таймер по умолчанию (в секундах)
    }
}
var AppConfig = LoadConfig() // Глобальная конфигурация


### Ниже подключение других файлов ко мне
### import "project/internal/config"

###func ConnectDB() {
###    dbPath := config.AppConfig.DBPath
###// Логика подключения к базе данных
###}
### Завершить пример скрипта выше

