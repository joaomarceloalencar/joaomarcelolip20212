#!/usr/bin/python3
import json

f = open("movies.json")
data = json.load(f)
f.close()

print(data)
data[1]['Title']="O Rebelde"

f = open("movies.json", "w")
json.dump(data, f)
f.close()