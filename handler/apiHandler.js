const apiHandler = {}
const fileHandler = require('./fileHandler')

apiHandler.v1 = v1Handler
apiHandler['v1.0'] = v1Handler
apiHandler['v1.0.0'] = v1Handler

function v1Handler (req, res) {
  const type = req.query.type || 'page'
  const target = req.query.target || 'neko'

  switch (type) {
    case 'redirect': {
      fileHandler(target, 'url', (result) => {
        if (result.success) res.redirect('/static/' + target + '/' + result.data)
        else res.sendStatus(result.data)
      })
      break
    }

    case 'buffer': {
      fileHandler(target, 'byte', (result) => {
        if (result.success) res.send(result.data)
        else res.sendStatus(result.data)
      })
      break
    }

    case 'page': {
      fileHandler(target, 'url', (result) => {
        if (result.success) res.send('<img src="/static/' + target + '/' + result.data + '">')
        else res.sendStatus(result.data)
      })
      break
    }

    case 'url': {
      fileHandler(target, 'url', (result) => {
        if (result.success) res.send('/static/' + target + '/' + result.data)
        else res.sendStatus(result.data)
      })
      break
    }

    case 'list': {
      fileHandler(target, 'list', (result) => {
        if (result.success) res.send(result.data)
        else res.sendStatus(result.data)
      })
      break
    }

    case 'listImg': {
      let str = ''
      fileHandler(target, 'list', (result) => {
        if (result.success) {
          result.data.forEach((d) => {
            str += '<img width="100" src="/static/' + target + '/' + d + '">'
          })
          res.send(str)
        } else res.sendStatus(result.data)
      })
      break
    }

    default:
      res.sendStatus(404)
  }
}

module.exports.apis = apiHandler
