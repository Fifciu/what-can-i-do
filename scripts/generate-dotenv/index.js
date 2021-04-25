require('colors');
const inquirer = require('inquirer');
const fs = require('fs');
const path = require('path');

const defaults = {
  pwa_protocol: 'http',
  pwa_host: 'localhost',
  pwa_port: 3000,
  use_varnish: 0,
  supported_oauth_providers: 'google',
  ls_token_key: 'session_token',
  ls_expires_key: 'session_expiry',
  jwt_key: 'sdasdasd8328ryewch828',
  jwt_ttl: 120000000,
  jwt_offset: 10000,
  api_port: 8090,
  db_name: 'whatcanido',
  db_user: 'user',
  db_pass: 'pass',
  db_port: 3306
};

(async () => {
  console.log('.env file creator for WhatCanIDo'.green);

  const values = await inquirer.prompt([
    {
      type: 'input',
      name: 'pwa_protocol',
      message: `What's the protocol of your PWA (${defaults.pwa_protocol})?`
    },
    {
      type: 'input',
      name: 'pwa_host',
      message: `What's the host of your PWA (${defaults.pwa_host})?`
    },
    {
      type: 'number',
      name: 'pwa_port',
      message: `What's the port of your PWA (${defaults.pwa_port})?`
    },
    {
      type: 'checkbox',
      name: 'supported_oauth_providers',
      message: `What OAuth providers are you going to use (${defaults.supported_oauth_providers})?`,
      choices: [
        'Google'
      ]
    },
    {
      type: 'input',
      name: 'ls_token_key',
      message: `What's the key of session token in your localStorage (${defaults.ls_token_key})?`
    },
    {
      type: 'input',
      name: 'ls_expires_key',
      message: `What's the key of expiration date of the session token in your localStorage (${defaults.ls_token_key})?`
    },
    {
      type: 'input',
      name: 'jwt_key',
      message: `What's the secret hashing key for JWT tokens, make sure it is unique (${defaults.jwt_key})?`
    },
    {
      type: 'number',
      name: 'jwt_ttl',
      message: `What's the time to live of your JWT Token (${defaults.jwt_ttl})?`
    },
    {
      type: 'number',
      name: 'jwt_offset',
      message: `What's the time to refresh after your JWT token expired (${defaults.jwt_offset})?`
    },
    {
      type: 'number',
      name: 'api_port',
      message: `What's the port of your backend API (${defaults.api_port})?`
    },
    {
      type: 'input',
      name: 'db_name',
      message: `What's the name of your MySQL database (${defaults.db_name})?`
    },
    {
      type: 'input',
      name: 'db_user',
      message: `What's the user of your MySQL database (${defaults.db_user})?`
    },
    {
      type: 'password',
      name: 'db_pass',
      message: `What's the password of your user in MySQL database (${defaults.db_pass})?`
    },
    {
      type: 'number',
      name: 'db_port',
      message: `On what port your database is listening (${defaults.db_port})?`
    },
  ]);

  values.supported_oauth_providers = values.supported_oauth_providers.map(v => v.toLowerCase());
  if (!values.supported_oauth_providers.length) {
    values.supported_oauth_providers.push('google');
  }
  console.log(`Please add to generated .env - 'google_client_id' and 'google_client_secret' from your Google Cloud Platform`.blue);

  const resolveValue = field => !!values[field] ? values[field] : defaults[field];

  const dotEnvContent = `
pwa_protocol=${resolveValue('pwa_protocol')}
pwa_host=${resolveValue('pwa_host')}
pwa_port=${resolveValue('pwa_port')}
use_varnish=0
ls_token_key=${resolveValue('ls_token_key')}
ls_expires_key=${resolveValue('ls_expires_key')}

supported_oauth_providers=${resolveValue('supported_oauth_providers').join(',')}

jwt_key=${resolveValue('jwt_key')}
jwt_ttl=${resolveValue('jwt_ttl')}
jwt_offset=${resolveValue('jwt_offset')}
api_port=${resolveValue('api_port')}

db_name=${resolveValue('db_name')}
db_user=${resolveValue('db_user')}
db_pass=${resolveValue('db_pass')}
db_port=${resolveValue('db_port')}
  `;

  const dotEnvPath = path.resolve(__dirname, '../../server') + '/.env';
  if (fs.existsSync(dotEnvPath)) {
    console.log('We are going to overwrite your existing .env'.red);
    const { confirmation } = await inquirer.prompt([
      {
        type: 'confirm',
        name: 'confirmation',
        message: `Are you sure you want to do that?`
      },
    ]);
    if (!confirmation) {
      return;
    }
    console.log(confirmation, 'confirmation')
  }

  console.log('Updating .env'.blue);
  fs.writeFileSync(dotEnvPath, dotEnvContent.trim());
})();