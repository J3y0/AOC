import networkx

g = networkx.Graph()

# Build graph
with open("./data/day25.txt") as f:
    for line in f.readlines():
        line = line.rstrip()
        node, links = line.split(": ")
        for link in links.split(" "):
            g.add_edge(node, link)

# Compute min cut
min_cut = networkx.minimum_edge_cut(g)
# Remove the found edges and find the 2 subgraphs
g.remove_edges_from(min_cut)
subgraphs_generator = networkx.connected_components(g)
# The generator should contain only 2 sets of nodes: multiply their length together
print("Part 1: ", len(next(subgraphs_generator))*len(next(subgraphs_generator)))