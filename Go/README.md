# StepOne
Our Go application must be able to download an entire folder from a web server (i.e. download all files via http and save to a local folder).
Программа выполняет следующие шаги:

Получает имена файлов с веб-сервера по указанному URL с помощью функции getFilesName.
Создает необходимые папки для сохранения файлов и лог-файлов с помощью функции createDirectories.
Создает лог-файлы successful.txt и not-successful.txt с помощью функции createLogFiles.
Для каждого файла из списка:
Проверяет, является ли файл MP3 файлом с помощью функции checkMP3File.
Если файл не является MP3 файлом, записывает его в лог-файл not-successful.txt.
Пытается скачать файл с веб-сервера:
Если связь прерывается или возникает ошибка при скачивании, программа ожидает 5 секунд и повторно пытается скачать файл.
Если не удается скачать файл после 12 попыток, программа делает паузу в 1 минуту и повторно пытается скачать файл.
Если с момента начала попыток скачивания прошло более 10 минут, программа прекращает попытки скачивания файла и записывает его в лог-файл not-successful.txt.
Если файл успешно скачан, программа сохраняет его в указанную папку и записывает его имя в лог-файл successful.txt.
По завершении всех операций программа завершает свою работу.

Сервер выполнен на Node.js
http://localhost:3001/api/files - Отправляет JSON объект с названиями файлов
http://localhost:3001/api/files/:name - Отправляет файл по названию файла, где :name - имя файла


Перед запуском сервера выполните установку зависимосте через `npm install`.
Сервер запускается командой `npm run dev` 


<!-- router.get('/files', async (req, res) => {
    const folderPath = 'music'; // Путь к папке с файлами
    fs.readdir(folderPath, (err, files) => {
        if (err) {
            return res.status(500).json({ error: 'Ошибка при чтении папки' });
        }
        res.json({ files: files }); // Отправляем JSON объект с названиями файлов
    });
});

router.get('/files/:name', async (req, res) => {
    const folderPath = 'music'; // Путь к папке с файлами
    const fileName = req.params.name; // Путь к папке с файлами

    if (!fileName || typeof fileName !== 'string') {
        return res.status(400).json({ error: 'Имя файла не указано или указано неверно' });
    }

    const filePath = path.resolve(folderPath, fileName);
    if (!fs.existsSync(filePath)) {
        return res.status(404).json({ error: 'Файл не найден' });
    }

    res.sendFile(filePath); // Отправляем файл по его пути
}); -->
