#!/bin/bash

# ARDUINO_PUBLIC_KEY_SECRET="ARDUINO_PUBLIC_KEY_SECRET"
# ARDUINO_PRIVATE_KEY_SECRET="ARDUINO_PRIVATE_KEY_SECRET"
# ARDUINO_CERT_KEY_SECRET="ARDUINO_CERT_KEY_SECRET"
# ARDUINO_CA_SECRET="ARDUINO_CA_SECRET"
# ARDUINO_CA_SECRET="ARDUINO_CA_SECRET"

# Utilize o AWS CLI para obter o valor do segredo
# PUBLIC_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_PUBLIC_KEY_SECRET --query 'SecretString' --output text)
# PRIVATE_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_PRIVATE_KEY_SECRET --query 'SecretString' --output text)
# CERT_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_CERT_KEY_SECRET --query 'SecretString' --output text)
# CA_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_CA_SECRET --query 'SecretString' --output text)

# Crie um arquivo local com o valor do segredo
# rm -rf ./keys
# mkdir ./keys
# echo "$PUBLIC_VALUE" > ./keys/greenhouse_01_humidity.public.key
# echo "$PRIVATE_VALUE" > ./keys/greenhouse_01_humidity.private.key
# echo "$CERT_VALUE" > ./keys/greenhouse_01_humidity.cert.pem
# echo "$CA_VALUE" > ./keys/root-CA.crt

AWS_SECRETS=$(aws secretsmanager get-secret-value --secret-id "greenhouse_api" --query 'SecretString' --output text)

export MONGO_URL=$(echo "$AWS_SECRETS" | jq -r '.ARDUINO_DATABASE_URL')
export DATABASE=$(echo "$AWS_SECRETS" | jq -r '.ARDUINO_DATABASE_VALUE')
export PORT=$(echo "$AWS_SECRETS" | jq -r '.HTTP_PORT')

# export AMBIENTE=DEV

exec /greenhouse_api
