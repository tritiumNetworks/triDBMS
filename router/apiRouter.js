const apiHandler = require('../handler/apiHandler')

function fnc (req, res) {
  const apiv = req.params.apiv || 'v1'
  if (Object.keys(apiHandler.apis).includes(apiv)) apiHandler.apis[apiv](req, res)
  else res.sendStatus(404)
}

module.exports = fnc
