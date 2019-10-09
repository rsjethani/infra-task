#!/usr/bin/env python3

# PS: No exception/error handling present in order to keep things simple

import os
import json
from ansinv import *
import subprocess as sb


# Initialize empty inventory
inventory = AnsibleInventory()

res = sb.run(["terraform", "output", "-json", "ip"], stdout=sb.PIPE).stdout.decode().strip()
ip = "[{}]".format(res)
host = AnsibleHost(json.loads(ip)[0])
inventory.add_hosts(host)

# Add groupvars if any to groups
inventory.group("all").groupvars.update(ansible_user=os.environ["ALFA_USER"], ansible_ssh_private_key_file=os.environ["ALFA_KEY"])

print(inventory)


with open("inventory", "w") as f:
    f.write(str(inventory))
