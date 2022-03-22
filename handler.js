'use strict';

const path = require('path')
const { spawn } = require('child_process')

module.exports.maintenance = async (event, context) => {
  console.log('calling Go binary')

  const child = spawn(path.join(__dirname, 'bin/maintenance'))
  child.stdout.pipe(process.stdout)
  child.stderr.pipe(process.stderr)

  return new Promise((resolve, reject) => {
    child.on('close', (code) => {
      const message = `child process exited with error code: ${code}`;
      if (code !== 0) {
        reject(message)
        return
      }
      resolve()
    })
  })
}
