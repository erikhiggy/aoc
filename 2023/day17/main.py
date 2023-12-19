import heapq
import numpy as np


def parse_input():
    with open("input.txt") as f:
        lines = f.read().splitlines()
    return lines

def find_shortest_paths(graph, start_point):
    # initialize graphs to track if a point is visited,
    # current calculated distance from start to point,
    # and previous point taken to get to current point
    visited = [[False for col in row] for row in graph]
    distance = [[float('inf') for col in row] for row in graph]
    distance[start_point[0]][start_point[1]] = 0
    prev_point = [[None for col in row] for row in graph]
    n, m = len(graph), len(graph[0])
    number_of_points, visited_count = n * m, 0
    directions = [(0, 1), (1, 0), (-1, 0), (0, -1)]
    min_heap = []

    # min_heap item format:
    # (pt's dist from start on this path, pt's row, pt's col)
    heapq.heappush(min_heap, (distance[start_point[0]][start_point[1]], start_point[0], start_point[1]))

    while visited_count < number_of_points:
        current_point = heapq.heappop(min_heap)
        distance_from_start, row, col = current_point
        for direction in directions:
            new_row, new_col = row + direction[0], col + direction[1]
            if -1 < new_row < n and -1 < new_col < m and not visited[new_row][new_col]:
                dist_to_new_point = distance_from_start + graph[new_row][new_col]
                if dist_to_new_point < distance[new_row][new_col]:
                    distance[new_row][new_col] = dist_to_new_point
                    prev_point[new_row][new_col] = (row, col)
                    heapq.heappush(min_heap, (dist_to_new_point, new_row, new_col))
        visited[row][col] = True
        visited_count += 1

    return distance, prev_point

def find_shortest_path(prev_point_graph, end_point):
    shortest_path = []
    current_point = end_point
    while current_point is not None:
        shortest_path.append(current_point)
        current_point = prev_point_graph[current_point[0]][current_point[1]]
    shortest_path.reverse()
    return shortest_path

def map_path(graph, path):
    for i in range(len(graph)):
        for j in range(len(graph[0])):
            if (i, j) in path and (i, j) != (0, 0):
                graph[i][j] = '#'
    return graph
            


def part1():
    input = parse_input()
    grid = []
    for line in input:
        grid.append(list(map(int, line)))
    dist, prev_point = find_shortest_paths(grid, (0,0))
    sp = find_shortest_path(prev_point, (12,12))
    grid = map_path(grid, sp)
    print(np.matrix(grid))
part1()