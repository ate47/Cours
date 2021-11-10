library(igraph)

g <- read.graph("graphncol.txt", format="ncol")

print(V(g)$name)
print(E(g)$weight)

mat2col <- as.matrix(read.table(file= "mat2cl.txt"))
g2bis <-graph.edgelist(mat2col)

print(g2bis)