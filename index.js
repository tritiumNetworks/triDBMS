const PORT = process.env.triDBMSPort || 8443

const cors = require('cors')
const path = require('path').resolve()
const https = require('https')
const express = require('express')
const apiRouter = require('./router/apiRouter')
const { readFileSync } = require('fs')

const app = express()
const ssl = { cert: readFileSync(path + '/cert/trinets-cert.pem'), key: readFileSync(path + '/cert/trinets-key.pem') }

app.use(cors())
app.use('/static', express.static(path + '/datas'))
app.use('/docs', express.static(path + '/docs'))

app.get('/', (_req, res) => res.redirect('/docs'))
app.get('/api', apiRouter)
app.get('/api/:apiv', apiRouter)

https.createServer(ssl, app).listen(PORT, () => { console.log('SSL Server is now on https://localhost:' + PORT) })
