import os
import subprocess

def buildDeploy(root, _, f):
    rootName = root[2:]
    currentDir = os.getcwd()
    fp = os.path.join(os.getcwd(), rootName)
    os.chdir(fp)
    subprocess.call(['bash', './builddeploy.sh'])
    os.chdir(currentDir)

def fileIs(fileName):
    return lambda _, __, file: file == fileName

predicateAndDo = []

predicateAndDo.append((fileIs('readme.md'), lambda root, _, file : os.system(f'markdown-to-slides ./{root}/{file} > ./{root}/presentation.html')))
predicateAndDo.append((fileIs('builddeploy.sh'), buildDeploy))

for (root, dirs, files) in os.walk("."):
    if './.' not in root:
        for fi in files:
            for predicate, do in predicateAndDo:
                if predicate(root, dirs, fi):
                    do(root, dirs, fi)