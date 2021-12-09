import os

# na = [1,2,3]
# for n in na:
#     print(n)

predicateAndDo = []

def buildDeploy(root, _, f):
    rootName = root[2:]
    currentDir = os.getcwd()
    fp = os.path.join(os.getcwd(), rootName)
    os.chdir(fp)
    print(fp)
    os.system(f)
    os.chdir(currentDir)

def fileIs(fileName):
    return lambda _, __, file: file == fileName


predicateAndDo.append((fileIs('readme.md'), lambda root, _, file : os.system(f'markdown-to-slides ./{root}/{file} > ./{root}/presentation.html')))
predicateAndDo.append((fileIs('builddeploy.sh'), buildDeploy))

for (root, dirs, files) in os.walk("."):
    if './.' not in root:
        for fi in files:
            for predicate, do in predicateAndDo:
                if predicate(root, dirs, fi):
                    do(root, dirs, fi)