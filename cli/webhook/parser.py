import json
from jinja2 import Environment, FileSystemLoader

# loading teplate
file_loader = FileSystemLoader('Template')
env = Environment(loader=file_loader)
template = env.get_template('template2')

# loading data
with open('event.json','r') as f:
    event = json.load(f)

print(template.render(event))