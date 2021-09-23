# %%
import re
from typing import Tuple
from sys import argv

if len(argv) < 2:
    print(argv[0], "(filename)")
    exit(-1)

# %%
doc_read_pattern = re.compile("<doc><docno>([^<]*)</docno>([^<]*)</doc>")


def read_doc(text: str) -> Tuple[str, str]:
    matcher = doc_read_pattern.search(text)
    if matcher == None:
        return None, None
    return matcher.group(1), matcher.group(2)


# %%
with open(argv[1], "r") as f:
    lines = f.readlines()


# %%
class IndexObject:
    def __init__(self) -> None:
        self.df = 0
        self.tf = dict()


# %%
index = dict()

# Building index

for i in range(len(lines)):
    line = lines[i]
    docno, doctext = read_doc(line)
    if docno == None:
        continue

    words = re.findall('\w+', doctext)
    for w in words:
        word = w.lower()

        if word in index:
            wl = index[word]
        else:
            wl = IndexObject()
            index[word] = wl

        if docno not in wl.tf:
            wl.df += 1
            wl.tf[docno] = 1
        else:
            wl.tf[docno] += 1


# %%
for word in sorted(index):
    io = index[word]
    print("{0}=df({1})".format(io.df, word))
    for doc in sorted(io.tf):
        print("\t{0} {1}".format(io.tf[doc], doc))


# %%
