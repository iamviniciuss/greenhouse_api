#!/bin/bash

# ARDUINO_PUBLIC_KEY_SECRET="ARDUINO_PUBLIC_KEY_SECRET"
# ARDUINO_PRIVATE_KEY_SECRET="ARDUINO_PRIVATE_KEY_SECRET"
# ARDUINO_CERT_KEY_SECRET="ARDUINO_CERT_KEY_SECRET"
# ARDUINO_CA_SECRET="ARDUINO_CA_SECRET"
# ARDUINO_CA_SECRET="ARDUINO_CA_SECRET"

# # Utilize o AWS CLI para obter o valor do segredo
# PUBLIC_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_PUBLIC_KEY_SECRET --query 'SecretString' --output text)
# PRIVATE_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_PRIVATE_KEY_SECRET --query 'SecretString' --output text)
# CERT_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_CERT_KEY_SECRET --query 'SecretString' --output text)
# CA_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_CA_SECRET --query 'SecretString' --output text)

# # Crie um arquivo local com o valor do segredo
# rm -rf ./keys
# mkdir ./keys
# echo "$PUBLIC_VALUE" > ./keys/greenhouse_01_humidity.public.key
# echo "$PRIVATE_VALUE" > ./keys/greenhouse_01_humidity.private.key
# echo "$CERT_VALUE" > ./keys/greenhouse_01_humidity.cert.pem
# echo "$CA_VALUE" > ./keys/root-CA.crt


# ARDUINO_DATABASE_URL_VALUE=$(aws secretsmanager get-secret-value --secret-id "ARDUINO_DATABASE_URL" --query 'SecretString' --output text)
# ARDUINO_BROKER_URL_VALUE=$(aws secretsmanager get-secret-value --secret-id "ARDUINO_BROKER_URL" --query 'SecretString' --output text)
# ARDUINO_DATABASE_VALUE=$(aws secretsmanager get-secret-value --secret-id "ARDUINO_BROKER_URL" --query 'SecretString' --output text)

# export MONGO_URL=$ARDUINO_DATABASE_URL_VALUE
# export BROKER_URL=$ARDUINO_BROKER_URL_VALUE
# export DATABASE=$ARDUINO_DATABASE_VALUE

if [ -z "$AMBIENTE" ]; then
  echo "A variável AMBIENTE não está definida."
elif [ "$AMBIENTE" = "PROD" ]; then
  echo "Você está executando em ambiente de produção."
  exec /greenhouse_api
elif [ "$AMBIENTE" = "DEV" ]; then
  exec /greenhouse_api
  echo "Você está executando em ambiente de desenvolvimento"
fi
