import express from 'express';
import router from './router.js'
import cors from 'cors'

const PORT = 3001;



const app = express();
app.use(cors());
app.use(express.json());


app.use('/api', router);


async function startApp() {
    try {
        
        app.listen(PORT, () => console.log('SERVER STARTED ON PORT ' + PORT));
    } catch (e) {
        console.error('Error connecting to MongoDB:', e.message);
    }
}




startApp();