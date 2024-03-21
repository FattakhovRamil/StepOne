import express from 'express';

import fs from 'fs'
import path from 'path'

const { Router } = express;
const router = Router();

router.get('/files', async (req, res) => {
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
});


export default router;

