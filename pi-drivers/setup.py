import setuptools

with open('README.md') as readme_f:
    README = readme_f.read()

with open('HISTORY.md') as history_f:
    HISTORY = history_f.read()

reqs = []
with open('requirments.txt')as f:
    reqs = f.read().splitlines()

setuptools.setup(
    name = 'rpi-driver',
    version = '0.01',
    install_requires = reqs,
    long_description_content_type = "text/markdown",
    long_description  = f{README} \n\n {HISTORY},
    author = 'Kyle Cooper',
    url = 'TBD',
    description= 'rpi driver program',
    license = 'MIT',
)

