use std::fs;

type Point = (usize, usize);

#[derive(Debug)]
struct Map {
    num_rows: usize,
    num_columns: usize,
    start: Point,
    end: Point,
    map: Vec<Vec<u8>>,
}

fn parse(input: String) -> Map {
    let mut my_map = Vec::new();
    let mut start: Point = (0, 0);
    let mut end: Point = (0, 0);
    for (line_i, line) in input.lines().enumerate() {
        let mut my_map_line: Vec<u8> = Vec::new();
        for (c_i, c) in line.chars().enumerate() {
            let value = match c {
                'S' => {
                    start = (line_i, c_i);
                    0
                }
                'E' => {
                    end = (line_i, c_i);
                    25
                }
                _ => c as u8 - 'a' as u8,
            };
            my_map_line.push(value);
        }
        my_map.push(my_map_line);
    }

    let rows: usize = my_map.iter().count();
    let columns = my_map.iter().next().unwrap().iter().count();
    Map {
        num_rows: rows,
        num_columns: columns,
        start: start,
        end: end,
        map: my_map,
    }
}

fn find_neighbors(vertex: Point, my_map: &Map) -> Vec<Point> {
    let num_columns = (*my_map).num_columns as isize;
    let num_rows = (*my_map).num_rows as isize;
    let current_row = vertex.0 as isize;
    let current_column = vertex.1 as isize;
    let candidates = [
        (current_row, current_column + 1),
        (current_row + 1, current_column),
        (current_row - 1, current_column),
        (current_row, current_column - 1),
    ];

    // You cannot leave the area
    candidates
        .iter()
        .map(|(r, c)| (*r as isize, *c as isize))
        .filter(|v| match v {
            (_, c) if *c == num_columns => false,
            (r, _) if *r == num_rows => false,
            (_, -1) => false,
            (-1, _) => false,
            _ => true,
        })
        .map(|(r, c)| (r as usize, c as usize))
        // Make sure you can access the next node
        .filter(|(r, c)| {
            my_map.map[current_row as usize][current_column as usize] + 1 >= my_map.map[*r][*c]
        })
        .collect()
}

fn find_min_and_remove_former(
    nodes: &mut Vec<Point>,
    distance: &Vec<u16>,
    num_columns: usize,
) -> Point {
    let (index, _) = nodes
        .iter()
        .enumerate()
        .map(|(ind, (r, c))| (ind, distance[*r * num_columns + *c]))
        .min_by_key(|(_, d)| *d)
        .unwrap();

    nodes.remove(index)
}

fn dijkstra(my_map: Map) -> Option<u32> {
    /*
     * Using Dijkstra algorithm but we return only the distance, we don't need
     * the shortest path here.
     */

    // Init
    let rows = my_map.num_rows;
    let columns = my_map.num_columns;
    let mut distance = vec![u16::MAX; rows * columns];
    let mut nodes = Vec::with_capacity(rows * columns);

    // Init nodes
    for i in 0..rows {
        for j in 0..columns {
            nodes.push((i, j));
        }
    }

    // Distance to the starting node from the starting node is 0
    let current_node: Point = my_map.start;
    distance[current_node.0 * columns + current_node.1] = 0;

    while !nodes.is_empty() {
        let current_node: Point = find_min_and_remove_former(&mut nodes, &distance, columns);
        let neighbors: Vec<Point> = find_neighbors(current_node, &my_map);
        let current_distance: u16 = distance[current_node.0 * columns + current_node.1];

        // Testing if we accessed the last node
        if current_node == my_map.end {
            return Some(current_distance as u32);
        }

        // Updating the next potential nodes' distance
        for v in neighbors {
            if distance[v.0 * columns + v.1] > current_distance.saturating_add(1) {
                distance[v.0 * columns + v.1] = current_distance.saturating_add(1);
            }
        }
    }
    return None;
}

fn main() {
    let content: String =
        fs::read_to_string("./data/day12.txt").expect("Should have been able to read file");

    let my_map: Map = parse(content);
    let part1: u32 = dijkstra(my_map).expect("Error while solving part 1");
    println!("Part 1: {part1}");
}
