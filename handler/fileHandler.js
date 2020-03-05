/* eslint-disable standard/no-callback-literal */

const path = require('path').resolve()
const fileHandler = {}
const { readdir, existsSync, mkdirSync, readFileSync } = require('fs')

if (!existsSync(path + '/datas/')) mkdirSync(path + '/datas')

fileHandler.list = getList
fileHandler.url = getRandUrl
fileHandler.byte = getRandByte

function getList (target, cb) {
  readdir(path + '/datas/', (err, files) => {
    if (err) console.log(err)

    if (files.includes(target)) {
      readdir(path + '/datas/' + target, (err2, files2) => {
        if (err2) console.log(err2)
        else if (files2.length < 1) cb({ success: false, data: 423 })
        else cb({ success: true, data: files2 })
      })
    } else cb({ success: false, data: 404 })
  })
}

function getRandUrl (target, cb) {
  readdir(path + '/datas/', (err, files) => {
    if (err) console.log(err)

    if (files.includes(target)) {
      readdir(path + '/datas/' + target, (err2, files2) => {
        if (err2) console.log(err2)
        else if (files2.length < 1) cb({ success: false, data: 423 })
        else cb({ success: true, data: files2[Math.floor(Math.random() * files2.length)] })
      })
    } else cb({ success: false, data: 404 })
  })
}

function getRandByte (target, cb) {
  getRandUrl(target, (res) => {
    if (!res.success) cb({ success: false, data: res.data })
    else cb({ success: true, data: readFileSync(path + '/datas/' + target + '/' + res.data) })
  })
}

function fnc (target, type, cb) {
  fileHandler[type](target, (res) => {
    cb(res)
  })
}

module.exports = fnc
