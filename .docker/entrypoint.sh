#!/bin/bash

ARDUINO_PUBLIC_KEY_SECRET="ARDUINO_PUBLIC_KEY_SECRET"
ARDUINO_PRIVATE_KEY_SECRET="ARDUINO_PRIVATE_KEY_SECRET"
ARDUINO_CERT_KEY_SECRET="ARDUINO_CERT_KEY_SECRET"
ARDUINO_CA_SECRET="ARDUINO_CA_SECRET"
ARDUINO_CA_SECRET="ARDUINO_CA_SECRET"

# Utilize o AWS CLI para obter o valor do segredo
PUBLIC_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_PUBLIC_KEY_SECRET --query 'SecretString' --output text)
PRIVATE_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_PRIVATE_KEY_SECRET --query 'SecretString' --output text)
CERT_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_CERT_KEY_SECRET --query 'SecretString' --output text)
CA_VALUE=$(aws secretsmanager get-secret-value --secret-id $ARDUINO_CA_SECRET --query 'SecretString' --output text)

# Crie um arquivo local com o valor do segredo
echo "$PUBLIC_VALUE" > /greenhouse_01_humidity.public.key
echo "$PRIVATE_VALUE" > /greenhouse_01_humidity.private.key
echo "$CERT_VALUE" > /greenhouse_01_humidity.cert.pem
echo "$CA_VALUE" > /root-CA.crt


ARDUINO_DATABASE_URL_VALUE=$(aws secretsmanager get-secret-value --secret-id "ARDUINO_DATABASE_URL" --query 'SecretString' --output text)
ARDUINO_BROKER_URL_VALUE=$(aws secretsmanager get-secret-value --secret-id "ARDUINO_BROKER_URL" --query 'SecretString' --output text)
ARDUINO_DATABASE_VALUE=$(aws secretsmanager get-secret-value --secret-id "ARDUINO_BROKER_URL" --query 'SecretString' --output text)

export MONGO_URL=$ARDUINO_DATABASE_URL_VALUE
export BROKER_URL=$ARDUINO_BROKER_URL_VALUE
export DATABASE=$ARDUINO_DATABASE_VALUE

# Inicie o seu aplicativo ou serviço (se aplicável)
exec /greenhouse_api