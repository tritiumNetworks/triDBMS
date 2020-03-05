const PORT = process.env.triDBMSPort || 8080

const cors = require('cors')
const path = require('path').resolve()
const express = require('express')
const apiRouter = require('./router/apiRouter')

const app = express()
app.use(cors())
app.use('/static', express.static(path + '/datas'))

app.get('/', (_req, res) => res.redirect('/api'))
app.get('/api', apiRouter)
app.get('/api/:apiv', apiRouter)

app.listen(PORT, () => console.log('Server is now on http://localhost:' + PORT))
