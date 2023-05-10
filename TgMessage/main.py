from telethon.sync import TelegramClient
from telethon.errors.rpcerrorlist import PeerFloodError
import sys
import time
import os
import tomllib
import pandas as pd

def connection(phone, api_id, api_hash):
    client = TelegramClient(phone, api_id, api_hash)
    
    client.connect()
    if not client.is_user_authorized():
        client.send_code_request(phone)
        client.sign_in(phone, input('Enter the code: '))

    return client

with open(os.getcwd()+"\config.toml","rb") as toml:
    config = tomllib.load(toml)

api_id = config['api_id']
api_hash = config['api_hash']
SLEEP_TIME = 5
CLIENT_PER_PHONE = 2
 
data = pd.read_csv(os.getcwd()+'\input.csv', names=['tgId'])
phones = pd.read_csv(os.getcwd()+'\phones.csv', names=['phone'])

data_size = data['tgId'].size - 1
user = 0

phone_size = phones['phone'].size
phone_id = 0

counter = 0

phone = str(phones['phone'][phone_id])
client = connection(phone, api_id, api_hash)

while user <= data_size:
    username = str(data['tgId'][user])
 
    message = '<a href="https://t.me/MariaSuhorukih_newbot">Подпишись и забери подарок</a>'
    photo = '\photo.jpg'
 
    if username == "":
        continue

    try:
        print("Get:", username)
        receiver = client.get_input_entity(username)
        print("Sending Message to:", username)
        client.send_file(receiver, os.getcwd()+photo)
        client.send_message(receiver, message, parse_mode="html")
        print("Waiting {} seconds".format(SLEEP_TIME))
        time.sleep(SLEEP_TIME)
    except PeerFloodError:
        print("Getting Flood Error from telegram. Script is stopping now. Please try again after some time.")
        client.disconnect()
        counter += CLIENT_PER_PHONE
    except Exception as e:
        print("Error:", e)
        print("Trying to continue...")

    user += 1
    counter += 1

    if (counter // CLIENT_PER_PHONE) > phone_id: 
        client.disconnect()

        phone_id = counter // CLIENT_PER_PHONE
        if phone_id == phone_size:
            print("No more phones")
            sys.exit

        phone = str(phones['phone'][phone_id])

        client = connection(phone, api_id, api_hash)
    print("Done. Message sent to all users.")

print(user)